package secretHandler

import (
	"context"

	"github.com/SENERGY-Platform/mgw-secret-manager/internal/logger"
	"github.com/SENERGY-Platform/mgw-secret-manager/internal/models"
	"github.com/SENERGY-Platform/mgw-secret-manager/pkg/api_model"
)

func (secretHandler *SecretHandler) GetSecrets(ctx context.Context) (secrets []*api_model.Secret, err error) {
	logger.Logger.Debugf("Load all short secrets")

	storedSecrets, err := secretHandler.db.GetSecrets(ctx)
	if err != nil {
		return
	}

	for _, storedSecret := range storedSecrets {
		shortSecret := api_model.Secret{
			Name:       storedSecret.Name,
			SecretType: storedSecret.SecretType,
			ID:         storedSecret.ID,
		}
		secrets = append(secrets, &shortSecret)
	}
	return
}

func (secretHandler *SecretHandler) GetSecret(ctx context.Context, id string) (secret models.Secret, err error) {
	logger.Logger.Debugf("Get clear text secret")
	storedSecret, err := secretHandler.db.GetSecret(ctx, id)
	if err != nil {
		return
	}

	if secretHandler.encryptionIsEnabled {
		logger.Logger.Debugf("Decrypt Secret Value")
		secret, err = secretHandler.DecryptSecret(storedSecret)
		if err != nil {
			return models.Secret{}, err
		}
	} else {
		secret = models.Secret{
			Name:       storedSecret.Name,
			SecretType: storedSecret.SecretType,
			ID:         storedSecret.ID,
			Value:      string(storedSecret.Value),
		}
	}
	return
}
