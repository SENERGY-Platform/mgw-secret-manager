package api

import (
	"testing"

	srv_base "github.com/SENERGY-Platform/go-service-base/srv-base"
	"github.com/SENERGY-Platform/mgw-secret-manager/internal/config"
	"github.com/SENERGY-Platform/mgw-secret-manager/internal/models"
	"github.com/SENERGY-Platform/mgw-secret-manager/internal/secretHandler"

	"github.com/SENERGY-Platform/mgw-secret-manager/pkg/api_model"
)

var testConfig, _ = config.NewConfig(nil)
var _, _ = srv_base.InitLogger(testConfig.Logger)

func SetupDummySecret(t *testing.T, name string, value string, secretType string, secretHandler secretHandler.SecretHandler) (models.Secret, api_model.ShortSecret) {
	secret := secretHandler.CreateSecret(name, value, secretType)
	err := secretHandler.StoreSecret(&secret)
	if err != nil {
		t.Errorf(err.Error())
	}

	return secret, api_model.ShortSecret{Name: name, SecretType: secretType, ID: secret.ID}
}
