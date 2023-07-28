package server

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

func TestGetValueVariant(t *testing.T) {
	type TestCase struct {
		Secret               api_model.SecretCreateRequest
		SecretVariantRequest api_model.SecretVariantRequest
		CaseName             string
		ExpectedValue        string
	}

	username := "username"

	testCases := []TestCase{
		{
			Secret:               api_model.SecretCreateRequest{Name: "name1", Value: "value1", SecretType: "Type1"},
			SecretVariantRequest: api_model.SecretVariantRequest{ID: "", Reference: "ref1", Item: nil},
			CaseName:             "Without Item",
			ExpectedValue:        "value1",
		},
		{
			Secret:               api_model.SecretCreateRequest{Name: "name2", Value: "{\"username\": \"user\", \"password\": \"password\"}", SecretType: "Type2"},
			SecretVariantRequest: api_model.SecretVariantRequest{ID: "", Reference: "ref2", Item: &username},
			CaseName:             "With Item",
			ExpectedValue:        "user",
		},
	}
	for _, tc := range testCases {
		t.Run(tc.CaseName, func(t *testing.T) {
			var config, _ = config.NewConfig(config.Flags.ConfPath)
			config.EnableEncryption = false
			config.ExposeConfidentialEndpoints = true
			router, dbHandler, secretHandler := InitServer(config)
			defer dbHandler.Cleanup()

			secret := SetupDummySecret(t, tc.Secret.Name, tc.Secret.Value, tc.Secret.SecretType, secretHandler)

			tc.SecretVariantRequest.ID = secret.ID
			body, err := json.Marshal(tc.SecretVariantRequest)
			if err != nil {
				t.Errorf(err.Error())
				return
			}

			req, _ := http.NewRequest("POST", api_model.ValueVariantPath, strings.NewReader(string(body)))
			w := httptest.NewRecorder()
			router.ServeHTTP(w, req)

			var secretResult api_model.SecretValueVariant
			json.NewDecoder(w.Body).Decode(&secretResult)
			assert.Equal(t, 200, w.Code)

			expectedSecret := api_model.SecretValueVariant{
				SecretVariant: api_model.SecretVariant{
					Secret: api_model.Secret{
						SecretType: secret.SecretType,
						Name:       secret.Name,
						ID:         secret.ID,
					},
					Item: tc.SecretVariantRequest.Item,
				},
				Value: tc.ExpectedValue,
			}
			assert.Equal(t, expectedSecret, secretResult)
		})
	}
}

func TestGetValueVariantBadPayload(t *testing.T) {
	var config, _ = config.NewConfig(config.Flags.ConfPath)
	config.EnableEncryption = false
	config.ExposeConfidentialEndpoints = true
	router, dbHandler, _ := InitServer(config)
	defer dbHandler.Cleanup()

	badPayload := "bad"
	body, err := json.Marshal(badPayload)
	if err != nil {
		t.Errorf(err.Error())
		return
	}

	req, _ := http.NewRequest("POST", api_model.ValueVariantPath, strings.NewReader(string(body)))
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, 500, w.Code)
}

func TestGetValueVariantNotFound(t *testing.T) {
	var config, _ = config.NewConfig(config.Flags.ConfPath)
	config.EnableEncryption = false
	config.ExposeConfidentialEndpoints = true
	router, dbHandler, _ := InitServer(config)
	defer dbHandler.Cleanup()

	badPayload := api_model.SecretVariantRequest{ID: "not_exist_id", Reference: "ref2"}
	body, err := json.Marshal(badPayload)
	if err != nil {
		t.Errorf(err.Error())
		return
	}

	req, _ := http.NewRequest("POST", api_model.ValueVariantPath, strings.NewReader(string(body)))
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, 500, w.Code)
}
