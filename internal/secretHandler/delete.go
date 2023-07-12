package secretHandler

import (
	"context"

	"github.com/SENERGY-Platform/mgw-secret-manager/internal/logger"
)

func (secretHandler *SecretHandler) DeleteSecret(ctx context.Context, secretID string) (err error) {
	logger.Logger.Debugf("Delete secret %s", secretID)
	err = secretHandler.db.DeleteSecret(ctx, secretID)
	return
}
