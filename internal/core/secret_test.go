package core

import (
	"secret-manager/internal/db"
	"secret-manager/test"
	"testing"

	srv_base "github.com/SENERGY-Platform/go-service-base/srv-base"

	"github.com/stretchr/testify/assert"
)

var _, err = srv_base.InitLogger(test.TestConfig.Logger)
var dbHandler, _ = db.NewDBHandler(test.TestConfig)

func TestStoreSecret(t *testing.T) {
	defer dbHandler.Cleanup()

	secretName := "test"
	secret := CreateSecret(secretName, "secret")
	err = StoreSecret(&secret, dbHandler, test.MasterKey)
	assert.Equal(t, err, nil)

	storedSecret, _ := GetSecret(secretName, dbHandler, test.MasterKey)
	assert.Equal(t, *storedSecret, secret)
}

func TestLoadSecretToTMPFS(t *testing.T) {
	defer dbHandler.Cleanup()

	config := test.TestConfig
	secretName := "test"
	secret := CreateSecret(secretName, "secret")
	_ = StoreSecret(&secret, dbHandler, test.MasterKey)
	fileName, err := LoadSecretToFileSystem(secretName, dbHandler, config, test.MasterKey)
	assert.Equal(t, nil, err)
	assert.Equal(t, secret.ID, fileName)
	// expect file exists
}
