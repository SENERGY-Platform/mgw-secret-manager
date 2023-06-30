package db

import (
	"github.com/SENERGY-Platform/mgw-secret-manager/internal/models"
)

type Database interface {
	SetSecret(secret *models.EncryptedSecret) (err error)
	GetSecret(secretName string) (secret *models.EncryptedSecret, err error)
	GetSecrets() (secrets []*models.EncryptedSecret, err error)
	Connect() (err error)
	Cleanup()
}
