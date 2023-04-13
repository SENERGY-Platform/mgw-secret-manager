package db

import (
	"testing"

	"github.com/SENERGY-Platform/mgw-secret-manager/internal/config"
	"github.com/SENERGY-Platform/mgw-secret-manager/internal/model"

	srv_base "github.com/SENERGY-Platform/go-service-base/srv-base"

	"github.com/stretchr/testify/assert"
)

var testConfig, _ = config.NewConfig(nil)
var _, err = srv_base.InitLogger(testConfig.Logger)

func TestSetSecret(t *testing.T) {
	testHandler, err := NewDBHandler(*testConfig)
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
	testHandler, err := NewDBHandler(*testConfig)
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
