package db

import (
	"github.com/SENERGY-Platform/mgw-secret-manager/internal/model"
)

type Database interface {
	SetSecret(secret *model.EncryptedSecret) (err error)
	GetSecret(secretName string) (secret *model.EncryptedSecret, err error)
	Connect() (err error)
	Cleanup()
}
