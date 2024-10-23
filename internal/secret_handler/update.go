package secret_handler

import (
	"context"
	"github.com/SENERGY-Platform/mgw-secret-manager/internal/models"
	"github.com/SENERGY-Platform/mgw-secret-manager/internal/util"
	"github.com/SENERGY-Platform/mgw-secret-manager/pkg/api_model"
)

func (secretHandler *SecretHandler) UpdateSecret(ctx context.Context, secretRequest api_model.SecretCreateRequest, secretID string) (err error) {
	util.Logger.Debugf("Update secret %s", secretID)

	secret := models.EncryptedSecret{
		Name:       secretRequest.Name,
		Value:      []byte(secretRequest.Value),
		SecretType: secretRequest.SecretType,
		ID:         secretID,
	}
	err = secretHandler.db.UpdateSecret(ctx, &secret)

	// If secret exists in TMPFS, reload it
	err = secretHandler.UpdateExistingSecretInTMPFS(ctx, secretID)
	return
}
