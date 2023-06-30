package secretHandler

import (
	"testing"

	"github.com/SENERGY-Platform/mgw-secret-manager/internal/config"
	"github.com/SENERGY-Platform/mgw-secret-manager/internal/db"
	"github.com/SENERGY-Platform/mgw-secret-manager/internal/models"

	"github.com/SENERGY-Platform/mgw-secret-manager/test"

	srv_base "github.com/SENERGY-Platform/go-service-base/srv-base"

	"github.com/stretchr/testify/assert"
)

var testConfig, _ = config.NewConfig(nil)
var _, err = srv_base.InitLogger(testConfig.Logger)

func TestStoreSecret(t *testing.T) {
	var dbHandler, _ = db.NewDBHandler(testConfig)
	dbHandler.Cleanup()
	defer dbHandler.Cleanup()

	testCases := []struct {
		secretName          string
		secretValue         string
		secretType          string
		encryptionIsEnabled bool
	}{
		{"secret1", "value1", "typ1", true},
		{"secret2", "value2", "typ2", false},
	}
	for _, tc := range testCases {
		secret := CreateSecret(tc.secretName, tc.secretValue, tc.secretType)
		err = StoreSecret(&secret, dbHandler, &test.MasterKey, tc.encryptionIsEnabled)
		assert.Equal(t, err, nil)

		storedSecret, _ := GetSecret(tc.secretName, dbHandler, &test.MasterKey, tc.encryptionIsEnabled)
		assert.Equal(t, secret, *storedSecret)
	}
}

func TestLoadSecretToTMPFS(t *testing.T) {
	var dbHandler, _ = db.NewDBHandler(testConfig)
	dbHandler.Cleanup()
	defer dbHandler.Cleanup()

	testCases := []struct {
		secretName          string
		secretValue         string
		secretType          string
		encryptionIsEnabled bool
	}{
		{"secret1", "value1", "typ1", true},
		{"secret2", "value2", "typ2", false},
	}
	for _, tc := range testCases {
		secret := CreateSecret(tc.secretName, tc.secretValue, tc.secretType)
		err = StoreSecret(&secret, dbHandler, &test.MasterKey, testConfig.EnableEncryption)
		if err != nil {
			t.Errorf(err.Error())
		}
		fileName, err := LoadSecretToFileSystem(tc.secretName, dbHandler, *testConfig, &test.MasterKey)
		assert.Equal(t, nil, err)
		assert.Equal(t, secret.ID, fileName)
	}
}

func TestEncryptSecret(t *testing.T) {
	secret := &models.Secret{
		Name:  "Test",
		Value: "value",
	}
	encryptedSecret, err := EncryptSecret(secret, test.MasterKey)
	if err != nil {
		t.Errorf(err.Error())
	}
	assert.Equal(t, encryptedSecret.Name, secret.Name)
}

func TestDecryptSecret(t *testing.T) {
	secret := &models.Secret{
		Name:  "Test",
		Value: "value",
	}
	encryptedSecret, err := EncryptSecret(secret, test.MasterKey)
	if err != nil {
		t.Errorf(err.Error())
	}

	decryptedSecret, err := DecryptSecret(encryptedSecret, test.MasterKey)
	if err != nil {
		t.Errorf(err.Error())
	}

	assert.Equal(t, decryptedSecret.Name, secret.Name)
	assert.Equal(t, decryptedSecret.Value, secret.Value)
}
