package secretHandler

import (
	"context"

	"github.com/SENERGY-Platform/mgw-secret-manager/internal/logger"
	"github.com/SENERGY-Platform/mgw-secret-manager/internal/models"
	"github.com/SENERGY-Platform/mgw-secret-manager/pkg/api_model"
)

func (secretHandler *SecretHandler) UpdateSecret(ctx context.Context, secretRequest api_model.SecretRequest, secretID string) (err error) {
	logger.Logger.Debugf("Update secret %s", secretID)

	secret := models.EncryptedSecret{
		Name:       secretRequest.Name,
		Value:      []byte(secretRequest.Value),
		SecretType: secretRequest.SecretType,
		ID:         secretID,
	}
	err = secretHandler.db.UpdateSecret(ctx, &secret)

	// If secret exists in TMPFS, reload it
	err = secretHandler.UpdateExistingSecretInTMPFS(ctx, api_model.SecretPostRequest{
		ID: secretID,
	}, true)
	return
}
