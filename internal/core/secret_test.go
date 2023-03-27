package core

import (
	"testing"

	"github.com/SENERGY-Platform/mgw-secret-manager/internal/db"
	"github.com/SENERGY-Platform/mgw-secret-manager/test"

	srv_base "github.com/SENERGY-Platform/go-service-base/srv-base"

	"github.com/stretchr/testify/assert"
)

var _, err = srv_base.InitLogger(test.TestConfig.Logger)

func TestStoreSecret(t *testing.T) {
	var dbHandler, _ = db.NewDBHandler(test.TestConfig)
	defer dbHandler.Cleanup()

	secretName := "test"
	secret := CreateSecret(secretName, "secret", "type")
	err = StoreSecret(&secret, dbHandler, &test.MasterKey, test.TestConfig)
	assert.Equal(t, err, nil)

	storedSecret, _ := GetSecret(secretName, dbHandler, &test.MasterKey, test.TestConfig)
	assert.Equal(t, *storedSecret, secret)
}

func TestLoadSecretToTMPFS(t *testing.T) {
	var dbHandler, _ = db.NewDBHandler(test.TestConfig)
	defer dbHandler.Cleanup()

	config := test.TestConfig
	secretName := "test"
	secret := CreateSecret(secretName, "secret", "type")
	_ = StoreSecret(&secret, dbHandler, &test.MasterKey, config)
	fileName, err := LoadSecretToFileSystem(secretName, dbHandler, config, &test.MasterKey)
	assert.Equal(t, nil, err)
	assert.Equal(t, secret.ID, fileName)
	// expect file exists
}
