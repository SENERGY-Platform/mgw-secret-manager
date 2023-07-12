package api

import (
	"encoding/json"
	"io/ioutil"
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

type TestCase struct {
	ExpectedValue     string
	Secret            api_model.SecretRequest
	SecretPostRequest api_model.SecretPostRequest
	CaseName          string
}

func TestLoadSecret(t *testing.T) {
	var config, _ = config.NewConfig(config.Flags.ConfPath)
	config.EnableEncryption = false
	username := "username"
	password := "password"

	testCases := []TestCase{
		// Simple single value secret
		{
			ExpectedValue:     "value1",
			Secret:            api_model.SecretRequest{Name: "name1", Value: "value1", SecretType: "Type1"},
			SecretPostRequest: api_model.SecretPostRequest{ID: "", Reference: "ref1", Item: nil},
			CaseName:          "Single Value",
		},

		// Credential secret where username is expected as value
		{
			ExpectedValue:     "user",
			Secret:            api_model.SecretRequest{Name: "name1", Value: "{\"username\": \"user\", \"password\": \"password\"}", SecretType: "Type1"},
			SecretPostRequest: api_model.SecretPostRequest{ID: "", Reference: "ref2", Item: &username},
			CaseName:          "Username",
		},

		// Credential secret where password is expected as value
		{
			ExpectedValue:     "password",
			Secret:            api_model.SecretRequest{Name: "name1", Value: "{\"username\": \"user\", \"password\": \"password\"}", SecretType: "Type1"},
			SecretPostRequest: api_model.SecretPostRequest{ID: "", Reference: "ref3", Item: &password},
			CaseName:          "Password",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.CaseName, func(t *testing.T) {
			w := httptest.NewRecorder()
			router, dbHandler, secretHandler := InitServer(config)
			defer dbHandler.Cleanup()

			secret, _ := SetupDummySecret(t, tc.Secret.Name, tc.Secret.Value, tc.Secret.SecretType, secretHandler)

			tc.SecretPostRequest.ID = secret.ID
			body, err := json.Marshal(tc.SecretPostRequest)
			if err != nil {
				t.Errorf(err.Error())
				return
			}

			req, _ := http.NewRequest("POST", "/load", strings.NewReader(string(body)))
			router.ServeHTTP(w, req)

			assert.Equal(t, 200, w.Code)

			pathToSecretInTMPFS := secretHandler.BuildTMPFSOutputPath(tc.SecretPostRequest)
			fullSecretPath := filepath.Join(config.TMPFSPath, pathToSecretInTMPFS)
			_, err = os.Stat(fullSecretPath)
			assert.Equal(t, nil, err)

			fileContent, err := ioutil.ReadFile(fullSecretPath)
			if err != nil {
				t.Errorf(err.Error())
				return
			}

			assert.Equal(t, tc.ExpectedValue, string(fileContent))
		})
	}
}
