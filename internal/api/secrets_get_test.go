package api

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
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
	var expectedSecrets []api_model.Secret
	secret1 := SetupDummySecret(t, "secret", "geheim", "type", secretHandler)
	expectedSecrets = append(expectedSecrets, api_model.Secret{
		Name:       secret1.Name,
		SecretType: secret1.SecretType,
		ID:         secret1.ID,
	})
	secret2 := SetupDummySecret(t, "secret2", "geheim2", "type2", secretHandler)
	expectedSecrets = append(expectedSecrets, api_model.Secret{
		Name:       secret2.Name,
		SecretType: secret2.SecretType,
		ID:         secret2.ID,
	})

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/secrets", w.Body)
	router.ServeHTTP(w, req)

	var secretResult []api_model.Secret
	json.NewDecoder(w.Body).Decode(&secretResult)
	assert.Equal(t, 200, w.Code)
	assert.ElementsMatch(t, expectedSecrets, secretResult)
}
