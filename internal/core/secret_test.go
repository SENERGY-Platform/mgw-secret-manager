package core

import (
	"secret-manager/internal/db"
	"secret-manager/internal/model"
	"secret-manager/test"
	"testing"

	srv_base "github.com/SENERGY-Platform/go-service-base/srv-base"

	"github.com/stretchr/testify/assert"
)

var dbHandler, _ = db.NewDBHandler(test.TestConfig)
var _, err = srv_base.InitLogger(test.TestConfig.Logger)

func TestStoreSecret(t *testing.T) {
	defer dbHandler.Cleanup()

	secretName := "test"
	secret := model.Secret{
		Name:  secretName,
		Value: "value",
	}
	masterKey, _ := GenerateMasterKey()
	err = StoreSecret(&secret, dbHandler, masterKey)
	assert.Equal(t, err, nil)

	storedSecret, _ := GetSecret(secretName, dbHandler, masterKey)
	assert.Equal(t, *storedSecret, secret)
}

func TestLoadSecretToTMPFS(t *testing.T) {
	defer dbHandler.Cleanup()

	config := test.TestConfig
	config.TMPFSPath = "/tmp"
	secretName := "test"
	secret := model.Secret{
		Name:  secretName,
		Value: "value",
	}
	masterKey, _ := GenerateMasterKey()
	_ = StoreSecret(&secret, dbHandler, masterKey)
	err := LoadSecretToFileSystem(secretName, dbHandler, config, masterKey)
	assert.Equal(t, err, nil)
}
