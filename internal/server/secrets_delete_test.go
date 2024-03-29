package server

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/SENERGY-Platform/mgw-secret-manager/internal/config"
	"github.com/SENERGY-Platform/mgw-secret-manager/pkg/api_model"
	"github.com/stretchr/testify/assert"
)

func TestDeleteSecret(t *testing.T) {
	w := httptest.NewRecorder()
	ctx := context.Background()
	var config, _ = config.NewConfig(config.Flags.ConfPath)
	config.EnableEncryption = false
	router, dbHandler, secretHandler := InitServer(config, "")
	defer dbHandler.Cleanup()

	secret := SetupDummySecret(t, "secret", "geheim", "type", secretHandler)

	req, _ := http.NewRequest("DELETE", api_model.SecretsPath+"/"+secret.ID, nil)
	router.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)

	_, err := secretHandler.GetSecret(ctx, secret.ID)
	assert.NotNil(t, err)
}
