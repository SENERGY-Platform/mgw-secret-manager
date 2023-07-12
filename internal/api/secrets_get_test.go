package api

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/SENERGY-Platform/mgw-secret-manager/internal/config"
	"github.com/SENERGY-Platform/mgw-secret-manager/pkg/api_model"
	"github.com/stretchr/testify/assert"
)

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

func TestGetFullSecret(t *testing.T) {
	type TestCase struct {
		Secret            api_model.SecretRequest
		SecretPostRequest api_model.SecretPostRequest
		CaseName          string
		ExpectedValue     string
	}

	username := "username"

	testCases := []TestCase{
		{
			Secret:            api_model.SecretRequest{Name: "name1", Value: "value1", SecretType: "Type1"},
			SecretPostRequest: api_model.SecretPostRequest{ID: "", Reference: "ref1", Item: nil},
			CaseName:          "Without Item",
			ExpectedValue:     "value1",
		},
		{
			Secret:            api_model.SecretRequest{Name: "name2", Value: "{\"username\": \"user\", \"password\": \"password\"}", SecretType: "Type2"},
			SecretPostRequest: api_model.SecretPostRequest{ID: "", Reference: "ref2", Item: &username},
			CaseName:          "With Item",
			ExpectedValue:     "user",
		},
	}
	for _, tc := range testCases {
		t.Run(tc.CaseName, func(t *testing.T) {
			var config, _ = config.NewConfig(config.Flags.ConfPath)
			config.EnableEncryption = false
			config.ExposeConfidentialEndpoints = true
			router, dbHandler, secretHandler := InitServer(config)
			defer dbHandler.Cleanup()

			_, shortSecret := SetupDummySecret(t, tc.Secret.Name, tc.Secret.Value, tc.Secret.SecretType, secretHandler)

			tc.SecretPostRequest.ID = shortSecret.ID
			body, err := json.Marshal(tc.SecretPostRequest)
			if err != nil {
				t.Errorf(err.Error())
				return
			}

			req, _ := http.NewRequest("POST", "/confidential/secret", strings.NewReader(string(body)))
			w := httptest.NewRecorder()
			router.ServeHTTP(w, req)

			var secretResult api_model.Secret
			json.NewDecoder(w.Body).Decode(&secretResult)
			assert.Equal(t, 200, w.Code)

			expectedSecret := api_model.Secret{
				Name:       shortSecret.Name,
				ID:         shortSecret.ID,
				Value:      tc.ExpectedValue,
				SecretType: shortSecret.SecretType,
				Path:       secretHandler.BuildTMPFSOutputPath(tc.SecretPostRequest),
				Item:       tc.SecretPostRequest.Item,
			}
			assert.Equal(t, expectedSecret, secretResult)
		})
	}
}
