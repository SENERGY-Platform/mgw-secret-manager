package core

import (
	"testing"

	"github.com/SENERGY-Platform/mgw-secret-manager/internal/config"
	"github.com/SENERGY-Platform/mgw-secret-manager/internal/db"
	"github.com/SENERGY-Platform/mgw-secret-manager/test"

	srv_base "github.com/SENERGY-Platform/go-service-base/srv-base"

	"github.com/stretchr/testify/assert"
)

var testConfig, _ = config.NewConfig(nil)
var _, err = srv_base.InitLogger(testConfig.Logger)

func TestStoreSecret(t *testing.T) {
	var dbHandler, _ = db.GetTestDB(testConfig)
	defer dbHandler.Cleanup()

	secretName := "test"
	secret := CreateSecret(secretName, "secret", "type")
	err = StoreSecret(&secret, &dbHandler, &test.MasterKey, *testConfig)
	assert.Equal(t, err, nil)

	storedSecret, _ := GetSecret(secretName, &dbHandler, &test.MasterKey, *testConfig)
	assert.Equal(t, *storedSecret, secret)
}

func TestLoadSecretToTMPFS(t *testing.T) {
	var dbHandler, _ = db.GetTestDB(testConfig)
	defer dbHandler.Cleanup()

	secretName := "test"
	secret := CreateSecret(secretName, "secret", "type")
	_ = StoreSecret(&secret, &dbHandler, &test.MasterKey, *testConfig)
	fileName, err := LoadSecretToFileSystem(secretName, &dbHandler, *testConfig, &test.MasterKey)
	assert.Equal(t, nil, err)
	assert.Equal(t, secret.ID, fileName)
	// expect file exists
}
