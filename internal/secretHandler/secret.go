package secretHandler

import (
	"context"
	"sync"

	"github.com/SENERGY-Platform/mgw-secret-manager/internal/db"
	"github.com/SENERGY-Platform/mgw-secret-manager/internal/logger"
)

type SecretHandler struct {
	encryptionIsEnabled bool
	db                  db.Database
	TMPFSPath           string
	Key                 []byte
	KeyMutex            sync.RWMutex // need a mutex for the master key
}

func NewSecretHandler(encryptionIsEnabled bool, db db.Database, TMPFSPath string) SecretHandler {
	return SecretHandler{
		encryptionIsEnabled: encryptionIsEnabled,
		db:                  db,
		TMPFSPath:           TMPFSPath,
	}
}

func (secretHandler *SecretHandler) SetKey(ctx context.Context, key []byte) {
	logger.Logger.Debugf("Save encryption key in secret handler")

	secretHandler.KeyMutex.Lock()
	secretHandler.Key = key
	secretHandler.KeyMutex.Unlock()
}
