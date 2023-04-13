package db

import (
	"errors"

	"github.com/SENERGY-Platform/mgw-secret-manager/internal/model"
)

type MockDBHandler struct {
	Secrets []*model.EncryptedSecret
}

func NewMockDB() (db *MockDBHandler) {
	return &MockDBHandler{}
}

func (handler *MockDBHandler) SetSecret(secret *model.EncryptedSecret) (err error) {
	handler.Secrets = append(handler.Secrets, secret)
	return nil
}

func (handler *MockDBHandler) GetSecret(secretName string) (secret *model.EncryptedSecret, err error) {
	for _, secret := range handler.Secrets {
		if secret.Name == secretName {
			return secret, nil
		}
	}
	return nil, errors.New("Secret not found")
}

func (handler *MockDBHandler) GetSecrets() (secrets []*model.EncryptedSecret, err error) {
	secrets = handler.Secrets
	return
}

func (handler *MockDBHandler) Connect() (err error) {
	return nil
}

func (handler *MockDBHandler) Cleanup() {
	handler.Secrets = []*model.EncryptedSecret{}
}
