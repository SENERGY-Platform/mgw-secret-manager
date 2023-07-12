package api

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/SENERGY-Platform/mgw-secret-manager/internal/config"
	"github.com/SENERGY-Platform/mgw-secret-manager/pkg/api_model"

	"github.com/stretchr/testify/assert"
)

func TestStoreSecret(t *testing.T) {
	var config, _ = config.NewConfig(config.Flags.ConfPath)
	config.EnableEncryption = false
	ctx := context.Background()

	testCases := []api_model.SecretRequest{
		{Name: "name1", Value: "value1", SecretType: "Type1"},
		{Name: "name1", Value: "value1", SecretType: "Type1"},
		{Name: "name2", Value: "value2", SecretType: "Type2"},
	}
	for _, tc := range testCases {
		w := httptest.NewRecorder()
		router, dbHandler, secretHandler := InitServer(config)
		defer dbHandler.Cleanup()

		body, err := json.Marshal(tc)
		if err != nil {
			t.Errorf(err.Error())
			return
		}

		req, _ := http.NewRequest("POST", "/secrets", strings.NewReader(string(body)))
		router.ServeHTTP(w, req)

		secretID := w.Body.String()

		assert.Equal(t, 200, w.Code)

		secretFromDB, err := secretHandler.GetSecret(ctx, api_model.SecretPostRequest{ID: secretID})
		if err != nil {
			t.Errorf(err.Error())
			return
		}
		assert.Equal(t, tc.Name, secretFromDB.Name)
		assert.Equal(t, tc.SecretType, secretFromDB.SecretType)
		assert.Equal(t, secretID, secretFromDB.ID)
	}
}

func TestGetSecrets(t *testing.T) {
	var config, _ = config.NewConfig(config.Flags.ConfPath)
	config.EnableEncryption = false
	router, dbHandler, secretHandler := InitServer(config)
	defer dbHandler.Cleanup()

	// Setup dummy secrets
	var expectedSecrets []api_model.ShortSecret
	_, shortSecret1 := SetupDummySecret(t, "secret", "geheim", "type", secretHandler)
	expectedSecrets = append(expectedSecrets, shortSecret1)
	_, shortSecret2 := SetupDummySecret(t, "secret2", "geheim2", "type2", secretHandler)
	expectedSecrets = append(expectedSecrets, shortSecret2)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/secrets", w.Body)
	router.ServeHTTP(w, req)

	var secretResult []api_model.ShortSecret
	json.NewDecoder(w.Body).Decode(&secretResult)
	assert.Equal(t, 200, w.Code)
	assert.ElementsMatch(t, expectedSecrets, secretResult)

}

func TestDeleteSecret(t *testing.T) {
	w := httptest.NewRecorder()
	ctx := context.Background()
	var config, _ = config.NewConfig(config.Flags.ConfPath)
	config.EnableEncryption = false
	router, dbHandler, secretHandler := InitServer(config)
	defer dbHandler.Cleanup()

	secret, _ := SetupDummySecret(t, "secret", "geheim", "type", secretHandler)

	req, _ := http.NewRequest("DELETE", "/secrets/"+secret.ID, nil)
	router.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)

	_, err := secretHandler.GetSecret(ctx, api_model.SecretPostRequest{ID: secret.ID})
	assert.NotNil(t, err)
}
