package client

import "github.com/SENERGY-Platform/mgw-secret-manager/pkg/api_model"

type Client interface {
	StoreSecret(name string, value string, secretType string) (err error, errCode int)
	GetSecrets() (secrets []api_model.ShortSecret, err error, errCode int)
	LoadSecretToTMPFS(secretName string) (fullTMPFSPath string, err error, errCode int)
	SetEncryptionKey(encryptionKey []byte) (err error, errCode int)
	UpdateSecret(name string, value string, secretType string, id string) (err error, errCode int)
	DeleteSecret(secretID string) (err error, errCode int)
}
