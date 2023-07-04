package api

import (
	"testing"

	"github.com/SENERGY-Platform/mgw-secret-manager/internal/models"
	"github.com/SENERGY-Platform/mgw-secret-manager/internal/secretHandler"

	"github.com/SENERGY-Platform/mgw-secret-manager/pkg/api_model"
)

func SetupDummySecret(t *testing.T, name string, value string, secretType string, secretHandler secretHandler.SecretHandler) (models.Secret, api_model.ShortSecret) {
	secret := secretHandler.CreateSecret(name, value, secretType)
	err := secretHandler.StoreSecret(&secret)
	if err != nil {
		t.Errorf(err.Error())
	}

	return secret, api_model.ShortSecret{Name: name, SecretType: secretType, ID: secret.ID}
}
