package server

import (
	"context"
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

func CreateDummySecretAndVariantRequest() {

}

func TestLoadSecret(t *testing.T) {
	type TestCase struct {
		ExpectedValue     string
		Secret            api_model.SecretCreateRequest
		SecretPostRequest api_model.SecretVariantRequest
		CaseName          string
	}

	var config, _ = config.NewConfig(config.Flags.ConfPath)
	config.EnableEncryption = false
	username := "username"
	password := "password"

	testCases := []TestCase{
		// Simple single value secret
		{
			ExpectedValue:     "value1",
			Secret:            api_model.SecretCreateRequest{Name: "name1", Value: "value1", SecretType: "Type1"},
			SecretPostRequest: api_model.SecretVariantRequest{ID: "", Reference: "ref1", Item: nil},
			CaseName:          "Single Value",
		},

		// Credential secret where username is expected as value
		{
			ExpectedValue:     "user",
			Secret:            api_model.SecretCreateRequest{Name: "name1", Value: "{\"username\": \"user\", \"password\": \"password\"}", SecretType: "Type1"},
			SecretPostRequest: api_model.SecretVariantRequest{ID: "", Reference: "ref2", Item: &username},
			CaseName:          "Username",
		},

		// Credential secret where password is expected as value
		{
			ExpectedValue:     "password",
			Secret:            api_model.SecretCreateRequest{Name: "name1", Value: "{\"username\": \"user\", \"password\": \"password\"}", SecretType: "Type1"},
			SecretPostRequest: api_model.SecretVariantRequest{ID: "", Reference: "ref3", Item: &password},
			CaseName:          "Password",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.CaseName, func(t *testing.T) {
			w := httptest.NewRecorder()
			router, dbHandler, secretHandler := InitServer(config)
			defer dbHandler.Cleanup()

			secret := SetupDummySecret(t, tc.Secret.Name, tc.Secret.Value, tc.Secret.SecretType, secretHandler)

			tc.SecretPostRequest.ID = secret.ID
			body, err := json.Marshal(tc.SecretPostRequest)
			if err != nil {
				t.Errorf(err.Error())
				return
			}

			req, _ := http.NewRequest("POST", api_model.LoadPathVariantPath, strings.NewReader(string(body)))
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

func TestDoubleLoad(t *testing.T) {
	var config, _ = config.NewConfig(config.Flags.ConfPath)
	config.EnableEncryption = false
	router, dbHandler, secretHandler := InitServer(config)
	defer dbHandler.Cleanup()

	secretCreateRequest := api_model.SecretCreateRequest{Name: "name1", Value: "value1", SecretType: "Type1"}
	secret := SetupDummySecret(t, secretCreateRequest.Name, secretCreateRequest.Value, secretCreateRequest.SecretType, secretHandler)
	secretPostRequest := api_model.SecretVariantRequest{ID: secret.ID, Reference: "ref1", Item: nil}

	body, err := json.Marshal(secretPostRequest)
	if err != nil {
		t.Errorf(err.Error())
		return
	}

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", api_model.LoadPathVariantPath, strings.NewReader(string(body)))
	router.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)

	// Load request after the file was already created
	w = httptest.NewRecorder()
	req, _ = http.NewRequest("POST", api_model.LoadPathVariantPath, strings.NewReader(string(body)))
	router.ServeHTTP(w, req)
	assert.Equal(t, 500, w.Code)
}

func TestInitPathVariant(t *testing.T) {
	var config, _ = config.NewConfig(config.Flags.ConfPath)
	config.EnableEncryption = false
	router, dbHandler, secretHandler := InitServer(config)
	defer dbHandler.Cleanup()

	secretCreateRequest := api_model.SecretCreateRequest{Name: "name1", Value: "value1", SecretType: "Type1"}
	secret := SetupDummySecret(t, secretCreateRequest.Name, secretCreateRequest.Value, secretCreateRequest.SecretType, secretHandler)
	secretPostRequest := api_model.SecretVariantRequest{ID: secret.ID, Reference: "ref1", Item: nil}

	body, err := json.Marshal(secretPostRequest)
	if err != nil {
		t.Errorf(err.Error())
		return
	}

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", api_model.InitPathVariantPath, strings.NewReader(string(body)))
	router.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)

	var secretPathVariant []api_model.SecretPathVariant
	json.NewDecoder(w.Body).Decode(&secretPathVariant)
	// TODO test variant returned correctly

	pathToSecretInTMPFS := secretHandler.BuildTMPFSOutputPath(secretPostRequest)
	fullSecretPath := filepath.Join(config.TMPFSPath, pathToSecretInTMPFS)
	fileContent, err := ioutil.ReadFile(fullSecretPath)
	if err != nil {
		t.Errorf(err.Error())
		return
	}

	assert.Equal(t, "", string(fileContent))
}

func TestInitLoad(t *testing.T) {
	var config, _ = config.NewConfig(config.Flags.ConfPath)
	config.EnableEncryption = false
	router, dbHandler, secretHandler := InitServer(config)
	defer dbHandler.Cleanup()

	secretCreateRequest := api_model.SecretCreateRequest{Name: "name1", Value: "value1", SecretType: "Type1"}
	expectedValue := "value1"
	secret := SetupDummySecret(t, secretCreateRequest.Name, secretCreateRequest.Value, secretCreateRequest.SecretType, secretHandler)
	secretPostRequest := api_model.SecretVariantRequest{ID: secret.ID, Reference: "ref1", Item: nil}

	body, err := json.Marshal(secretPostRequest)
	if err != nil {
		t.Errorf(err.Error())
		return
	}

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", api_model.InitPathVariantPath, strings.NewReader(string(body)))
	router.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)

	w = httptest.NewRecorder()
	req, _ = http.NewRequest("POST", api_model.LoadPathVariantPath, strings.NewReader(string(body)))
	router.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)

	pathToSecretInTMPFS := secretHandler.BuildTMPFSOutputPath(secretPostRequest)
	fullSecretPath := filepath.Join(config.TMPFSPath, pathToSecretInTMPFS)
	fileContent, err := ioutil.ReadFile(fullSecretPath)
	if err != nil {
		t.Errorf(err.Error())
		return
	}

	assert.Equal(t, expectedValue, string(fileContent))
}

func TestUnloadSecret(t *testing.T) {
	type TestCase struct {
		Secret            api_model.SecretCreateRequest
		SecretPostRequest api_model.SecretVariantRequest
		CaseName          string
	}

	var config, _ = config.NewConfig(config.Flags.ConfPath)
	config.EnableEncryption = false
	username := "username"

	testCases := []TestCase{
		// Simple single value secret
		{
			Secret:            api_model.SecretCreateRequest{Name: "name1", Value: "value1", SecretType: "Type1"},
			SecretPostRequest: api_model.SecretVariantRequest{ID: "", Reference: "ref1", Item: nil},
			CaseName:          "Single Value",
		},
		// Credential secret where filename is Item value
		{
			Secret:            api_model.SecretCreateRequest{Name: "name1", Value: "{\"username\": \"user\", \"password\": \"password\"}", SecretType: "Type1"},
			SecretPostRequest: api_model.SecretVariantRequest{ID: "", Reference: "ref2", Item: &username},
			CaseName:          "JSON Value",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.CaseName, func(t *testing.T) {
			w := httptest.NewRecorder()
			ctx := context.Background()
			router, dbHandler, secretHandler := InitServer(config)
			defer dbHandler.Cleanup()

			secret := SetupDummySecret(t, tc.Secret.Name, tc.Secret.Value, tc.Secret.SecretType, secretHandler)

			tc.SecretPostRequest.ID = secret.ID
			secretHandler.LoadSecretToFileSystem(ctx, tc.SecretPostRequest)

			tc.SecretPostRequest.ID = secret.ID
			body, err := json.Marshal(tc.SecretPostRequest)
			if err != nil {
				t.Errorf(err.Error())
				return
			}

			req, _ := http.NewRequest("POST", api_model.UnLoadPathVariantPath, strings.NewReader(string(body)))
			router.ServeHTTP(w, req)

			assert.Equal(t, 200, w.Code)

			pathToSecretInTMPFS := secretHandler.BuildTMPFSOutputPath(tc.SecretPostRequest)
			fullSecretPath := filepath.Join(config.TMPFSPath, pathToSecretInTMPFS)
			_, err = os.Stat(fullSecretPath)
			assert.NotNil(t, err)
		})
	}
}

func TestCleanReference(t *testing.T) {
	var config, _ = config.NewConfig(config.Flags.ConfPath)
	w := httptest.NewRecorder()
	ctx := context.Background()
	router, dbHandler, secretHandler := InitServer(config)
	defer dbHandler.Cleanup()

	referenceToDelete := "ref1"
	referenceToKeep := "ref2"

	// Load dummy secrets
	secret1 := SetupDummySecret(t, "secret", "value", "type", secretHandler)
	secret2 := SetupDummySecret(t, "secret", "value", "type", secretHandler)

	type Case struct {
		SecretID  string
		Reference string
	}

	cases := []Case{
		Case{
			SecretID:  secret1.ID,
			Reference: referenceToDelete,
		},
		Case{
			SecretID:  secret2.ID,
			Reference: referenceToKeep,
		},
	}
	for _, secretCase := range cases {
		secretPostRequest := api_model.SecretVariantRequest{
			ID:        secretCase.SecretID,
			Reference: secretCase.Reference,
		}
		secretHandler.LoadSecretToFileSystem(ctx, secretPostRequest)
	}

	req, _ := http.NewRequest("POST", api_model.CleanPathVariantsPath+"?reference="+referenceToDelete, nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)

	refToBeDeletedPath := filepath.Join(config.TMPFSPath, referenceToDelete)
	_, err := os.Stat(refToBeDeletedPath)
	assert.NotNil(t, err)

	refToKeepPath := filepath.Join(config.TMPFSPath, referenceToKeep)
	_, err = os.Stat(refToKeepPath)
	assert.Nil(t, err)
}
