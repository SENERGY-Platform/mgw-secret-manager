package client

import (
	"context"

	"github.com/SENERGY-Platform/mgw-secret-manager/pkg/api_model"
)

type Client interface {
	StoreSecret(ctx context.Context, name string, value string, secretType string) (err error, errCode int)
	GetSecrets(ctx context.Context) (secrets []api_model.ShortSecret, err error, errCode int)
	GetSecret(ctx context.Context, secretRequest api_model.SecretPostRequest) (secrets *api_model.ShortSecret, err error, errCode int)
	GetFullSecret(ctx context.Context, secretRequest api_model.SecretPostRequest) (secrets *api_model.Secret, err error, errCode int)
	LoadSecretToTMPFS(ctx context.Context, secretRequest api_model.SecretPostRequest) (fullTMPFSPath string, err error, errCode int)
	SetEncryptionKey(ctx context.Context, encryptionKey []byte) (err error, errCode int)
	UpdateSecret(ctx context.Context, name string, value string, secretType string, id string) (err error, errCode int)
	DeleteSecret(ctx context.Context, secretID string) (err error, errCode int)
}
