package db

import (
	"testing"

	"github.com/SENERGY-Platform/mgw-secret-manager/internal/model"
	"github.com/SENERGY-Platform/mgw-secret-manager/test"

	srv_base "github.com/SENERGY-Platform/go-service-base/srv-base"

	"github.com/stretchr/testify/assert"
)

var _, err = srv_base.InitLogger(test.TestConfig.Logger)

func TestSetSecret(t *testing.T) {
	testHandler, err := NewDBHandler(test.TestConfig)
	defer testHandler.Cleanup()
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
	defer testHandler.Cleanup()
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
