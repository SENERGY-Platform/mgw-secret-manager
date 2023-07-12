package api

import (
	"encoding/json"
	"fmt"
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
	ExpectedValue string
	Secret        api_model.SecretRequest
	Reference     string
	Item          string
}

func TestLoadSecret(t *testing.T) {
	var config, _ = config.NewConfig(config.Flags.ConfPath)
	config.EnableEncryption = false
	w := httptest.NewRecorder()

	testCases := []TestCase{
		// Simple single value secret
		{
			ExpectedValue: "value1",
			Secret:        api_model.SecretRequest{Name: "name1", Value: "value1", SecretType: "Type1"},
		},

		// Credential secret where username is expected as value
		{
			ExpectedValue: "user",
			Secret:        api_model.SecretRequest{Name: "name1", Value: "{\"username\": \"user\", \"password\": \"password\"}", SecretType: "Type1"},
			Reference:     "ref",
			Item:          "username",
		},

		// Credential secret where password is expected as value
		{
			ExpectedValue: "password",
			Secret:        api_model.SecretRequest{Name: "name1", Value: "{\"username\": \"useer\", \"password\": \"password\"}", SecretType: "Type1"},
			Reference:     "ref",
			Item:          "password",
		},
	}

	for _, tc := range testCases {
		router, dbHandler, secretHandler := InitServer(config)
		defer dbHandler.Cleanup()

		secret, _ := SetupDummySecret(t, tc.Secret.Name, tc.Secret.Value, tc.Secret.SecretType, secretHandler)

		body, err := json.Marshal(api_model.SecretPostRequest{ID: secret.ID, Reference: tc.Reference, Item: tc.Item})
		if err != nil {
			t.Errorf(err.Error())
		}

		req, _ := http.NewRequest("POST", "/load", strings.NewReader(string(body)))
		router.ServeHTTP(w, req)

		assert.Equal(t, 200, w.Code)

		pathToSecretInTMPFS := filepath.Join(secretHandler.TMPFSPath, secret.ID, fmt.Sprintf("%s_%s_%s", secret.ID, tc.Reference, tc.Item))
		_, err = os.Stat(pathToSecretInTMPFS)
		assert.Equal(t, nil, err)

		fileContent, err := ioutil.ReadFile(pathToSecretInTMPFS)
		if err != nil {
			t.Errorf(err.Error())
		}

		assert.Equal(t, tc.ExpectedValue, string(fileContent))
	}
}
