package db

import (
	"context"

	"github.com/SENERGY-Platform/mgw-secret-manager/internal/models"
)

type Database interface {
	SetSecret(ctx context.Context, secret *models.EncryptedSecret) (err error)
	GetSecret(ctx context.Context, secretName string) (secret *models.EncryptedSecret, err error)
	GetSecrets(ctx context.Context) (secrets []*models.EncryptedSecret, err error)
	UpdateSecret(ctx context.Context, secret *models.EncryptedSecret) (err error)
	DeleteSecret(ctx context.Context, secretID string) (err error)
	Connect() (err error)
	Cleanup()
}
