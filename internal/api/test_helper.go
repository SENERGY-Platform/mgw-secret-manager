package api

import (
	"context"
	"testing"

	"github.com/SENERGY-Platform/mgw-secret-manager/internal/secretHandler"

	"github.com/SENERGY-Platform/mgw-secret-manager/pkg/api_model"
)

func SetupDummySecret(t *testing.T, name string, value string, secretType string, fileName string, secretHandler secretHandler.SecretHandler) (api_model.Secret, api_model.ShortSecret) {
	secret := secretHandler.CreateSecret(name, value, secretType, fileName)
	ctx := context.Background()
	err := secretHandler.StoreSecret(ctx, &secret)
	if err != nil {
		t.Errorf(err.Error())
	}

	return secret, api_model.ShortSecret{
		Name:       name,
		SecretType: secretType,
		ID:         secret.ID,
		FileName:   secret.FileName,
	}
}
