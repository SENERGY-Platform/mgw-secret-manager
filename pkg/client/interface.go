package client

import (
	"context"
	"github.com/SENERGY-Platform/mgw-secret-manager/pkg/api_model"
)

type Client interface {
	StoreSecret(ctx context.Context, name string, value string, secretType string) (err error, errCode int)
	GetSecrets(ctx context.Context) (secrets []api_model.Secret, err error, errCode int)
	UpdateSecret(ctx context.Context, name string, value string, secretType string, id string) (err error, errCode int)
	DeleteSecret(ctx context.Context, secretID string) (err error, errCode int)
	SetEncryptionKey(ctx context.Context, encryptionKey []byte) (err error, errCode int)
	InitPathVariant(ctx context.Context, secretRequest api_model.SecretVariantRequest) (variant api_model.SecretPathVariant, err error, errCode int)
	LoadPathVariant(ctx context.Context, secretRequest api_model.SecretVariantRequest) (err error, errCode int)
	UnloadPathVariant(ctx context.Context, secretRequest api_model.SecretVariantRequest) (err error, errCode int)
	CleanPathVariants(ctx context.Context, ref string) (err error, errCode int)
	GetValueVariant(ctx context.Context, secretRequest api_model.SecretVariantRequest) (variant api_model.SecretValueVariant, err error, errCode int)
}
