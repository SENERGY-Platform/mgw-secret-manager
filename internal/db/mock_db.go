package db

import (
	"errors"
	"secret-manager/internal/model"
)

type MockDBHandler struct {
	Secrets []model.EncryptedSecret
}

func NewMockDB() (db Database) {
	return &MockDBHandler{}
}

func (handler *MockDBHandler) SetSecret(secret *model.EncryptedSecret) (err error) {
	handler.Secrets = append(handler.Secrets, *secret)
	return nil
}

func (handler *MockDBHandler) GetSecret(secretName string) (secret *model.EncryptedSecret, err error) {
	for _, secret := range handler.Secrets {
		if secret.Name == secretName {
			return &secret, nil
		}
	}
	return nil, errors.New("Secret not found")
}

func (handler *MockDBHandler) Connect() (err error) {
	return nil
}

func (handler *MockDBHandler) Cleanup() {
	handler.Secrets = []model.EncryptedSecret{}
}
