package db

import (
	"secret-manager/internal/model"
	"secret-manager/test"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSetSecret(t *testing.T) {
	testHandler, err := NewDBHandler(test.TestConfig)
	if err != nil {
		t.Errorf(err.Error())
	}
	secret := &model.EncryptedSecret{
		Name:  "test",
		Value: make([]byte, 2),
	}
	err = testHandler.SetSecret(secret)
	assert.Equal(t, err, nil)
}

func TestGetSecret(t *testing.T) {
	secretName := "test"
	testHandler, err := NewDBHandler(test.TestConfig)
	if err != nil {
		t.Errorf(err.Error())
	}
	secret := &model.EncryptedSecret{
		Name:  secretName,
		Value: make([]byte, 2),
	}
	err = testHandler.SetSecret(secret)

	storedSecret, err := testHandler.GetSecret(secretName)
	assert.Equal(t, err, nil)
	assert.Equal(t, storedSecret, secret)
}
