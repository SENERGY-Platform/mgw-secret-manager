package api

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/SENERGY-Platform/mgw-secret-manager/internal/config"
	"github.com/SENERGY-Platform/mgw-secret-manager/pkg/api_model"

	"github.com/stretchr/testify/assert"
)

func TestLoadSecret(t *testing.T) {
	var config, _ = config.NewConfig(config.Flags.ConfPath)
	config.EnableEncryption = false

	testCases := []api_model.SecretRequest{
		{Name: "name1", Value: "value1", SecretType: "Type1", FileName: "file.txt"},
		{Name: "name1", Value: "value1", SecretType: "Type1", FileName: "file1.txt"},
		{Name: "name2", Value: "value2", SecretType: "Type2", FileName: "file2.txt"},
	}
	for _, tc := range testCases {
		router, dbHandler, secretHandler := InitServer(config)
		defer dbHandler.Cleanup()

		secret, _ := SetupDummySecret(t, tc.Name, tc.Value, tc.SecretType, tc.FileName, secretHandler)

		w := httptest.NewRecorder()
		req, _ := http.NewRequest("POST", fmt.Sprintf("/load?secret=%s", secret.ID), nil)
		router.ServeHTTP(w, req)

		assert.Equal(t, 200, w.Code)

		pathToSecretInTMPFS := filepath.Join(secretHandler.TMPFSPath, secret.ID, secret.FileName)
		_, err := os.Stat(pathToSecretInTMPFS)
		// TODO assert file value == secret value
		assert.Equal(t, nil, err)
	}
}

func TestLoadSecretMissingQuery(t *testing.T) {
	var config, _ = config.NewConfig(config.Flags.ConfPath)
	config.EnableEncryption = false
	router, dbHandler, _ := InitServer(config)
	defer dbHandler.Cleanup()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/load", nil)
	router.ServeHTTP(w, req)
	assert.Equal(t, 500, w.Code)
}

func TestStoreSecret(t *testing.T) {
	var config, _ = config.NewConfig(config.Flags.ConfPath)
	config.EnableEncryption = false
	w := httptest.NewRecorder()
	ctx := context.Background()

	testCases := []api_model.SecretRequest{
		{Name: "name1", Value: "value1", SecretType: "Type1"},
		{Name: "name1", Value: "value1", SecretType: "Type1"},
		{Name: "name2", Value: "value2", SecretType: "Type2"},
	}
	for _, tc := range testCases {
		router, dbHandler, secretHandler := InitServer(config)
		defer dbHandler.Cleanup()

		body, err := json.Marshal(tc)
		if err != nil {
			t.Errorf(err.Error())
		}

		req, _ := http.NewRequest("POST", "/secrets", strings.NewReader(string(body)))
		router.ServeHTTP(w, req)

		var secretID string

		json.NewDecoder(w.Body).Decode(&secretID)

		assert.Equal(t, 200, w.Code)

		secretFromDB, err := secretHandler.GetSecret(ctx, secretID)
		if err != nil {
			t.Errorf(err.Error())
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
	_, shortSecret1 := SetupDummySecret(t, "secret", "geheim", "type", "file.txt", secretHandler)
	expectedSecrets = append(expectedSecrets, shortSecret1)
	_, shortSecret2 := SetupDummySecret(t, "secret2", "geheim2", "type2", "file2.txt", secretHandler)
	expectedSecrets = append(expectedSecrets, shortSecret2)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/secrets", w.Body)
	router.ServeHTTP(w, req)

	var secretResult []api_model.ShortSecret
	json.NewDecoder(w.Body).Decode(&secretResult)
	assert.Equal(t, 200, w.Code)
	assert.ElementsMatch(t, expectedSecrets, secretResult)

}

type a struct {
	ExistingSecret api_model.SecretRequest
	ChangedSecret  api_model.SecretRequest
}

func TestUpdateSecret(t *testing.T) {
	var config, _ = config.NewConfig(config.Flags.ConfPath)
	config.EnableEncryption = false
	w := httptest.NewRecorder()
	ctx := context.Background()

	testCases := []a{
		// Change Name
		{
			ExistingSecret: api_model.SecretRequest{
				Name:       "name",
				FileName:   "file.txt",
				Value:      "value1",
				SecretType: "type1",
			},
			ChangedSecret: api_model.SecretRequest{
				Name:       "name2",
				FileName:   "file2.txt",
				Value:      "value2",
				SecretType: "type2",
			},
		},
	}
	for _, tc := range testCases {
		router, dbHandler, secretHandler := InitServer(config)
		defer dbHandler.Cleanup()

		_, shortSecret := SetupDummySecret(t, tc.ExistingSecret.Name, tc.ExistingSecret.Value, tc.ExistingSecret.SecretType, tc.ExistingSecret.FileName, secretHandler)
		secretID := shortSecret.ID

		body, err := json.Marshal(tc.ChangedSecret)
		if err != nil {
			t.Errorf(err.Error())
		}

		req, _ := http.NewRequest("PUT", "/secrets/"+secretID, strings.NewReader(string(body)))
		router.ServeHTTP(w, req)

		var response string

		json.NewDecoder(w.Body).Decode(&response)

		assert.Equal(t, 200, w.Code)

		secretFromDB, err := secretHandler.GetSecret(ctx, secretID)
		if err != nil {
			t.Errorf(err.Error())
		}
		assert.Equal(t, tc.ChangedSecret.Name, secretFromDB.Name)
		assert.Equal(t, tc.ChangedSecret.SecretType, secretFromDB.SecretType)
		assert.Equal(t, secretID, secretFromDB.ID)
	}
}

func TestDeleteSecret(t *testing.T) {
	w := httptest.NewRecorder()
	ctx := context.Background()
	var config, _ = config.NewConfig(config.Flags.ConfPath)
	config.EnableEncryption = false
	router, dbHandler, secretHandler := InitServer(config)
	defer dbHandler.Cleanup()

	secret, _ := SetupDummySecret(t, "secret", "geheim", "type", "file.txt", secretHandler)

	req, _ := http.NewRequest("DELETE", "/secrets/"+secret.ID, nil)
	router.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)

	_, err := secretHandler.GetSecret(ctx, secret.ID)
	assert.NotNil(t, err)
}
