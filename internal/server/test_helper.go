package server

import (
	"context"
	"testing"

	"github.com/SENERGY-Platform/mgw-secret-manager/internal/models"
	"github.com/SENERGY-Platform/mgw-secret-manager/internal/secretHandler"
)

func SetupDummySecret(t *testing.T, name string, value string, secretType string, secretHandler secretHandler.SecretHandler) models.Secret {
	secret := secretHandler.CreateSecret(name, value, secretType)
	ctx := context.Background()
	err := secretHandler.StoreSecret(ctx, &secret)
	if err != nil {
		t.Errorf(err.Error())
	}

	return models.Secret{
		Name:       name,
		SecretType: secretType,
		ID:         secret.ID,
		Value:      value,
	}
}
