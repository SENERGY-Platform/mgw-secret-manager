package secret_handler

import (
	"context"
	"github.com/SENERGY-Platform/mgw-secret-manager/internal/util"
)

func (secretHandler *SecretHandler) DeleteSecret(ctx context.Context, secretID string) (err error) {
	util.Logger.Debugf("Delete secret %s", secretID)
	err = secretHandler.db.DeleteSecret(ctx, secretID)
	return
}
