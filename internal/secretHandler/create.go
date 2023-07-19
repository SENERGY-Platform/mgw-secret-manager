package secretHandler

import (
	"context"

	"github.com/SENERGY-Platform/mgw-secret-manager/internal/logger"
	"github.com/SENERGY-Platform/mgw-secret-manager/internal/models"
	"github.com/google/uuid"
)

func (secretHandler *SecretHandler) CreateSecret(name string, value string, secretType string) models.Secret {
	return models.Secret{
		Name:       name,
		SecretType: secretType,
		ID:         uuid.New().String(),
		Value:      value,
	}
}

func (secretHandler *SecretHandler) StoreSecret(ctx context.Context, secret *models.Secret) (err error) {
	logger.Logger.Debugf("Store Secret: %s", secret.ID)

	var storedSecret *models.EncryptedSecret

	if secretHandler.encryptionIsEnabled {
		storedSecret, err = secretHandler.EncryptSecret(secret)
		if err != nil {
			return
		}
	} else {
		storedSecret = &models.EncryptedSecret{
			Name:       secret.Name,
			SecretType: secret.SecretType,
			Value:      []byte(secret.Value),
			ID:         secret.ID,
		}
	}

	err = secretHandler.db.SetSecret(ctx, storedSecret)
	return
}
