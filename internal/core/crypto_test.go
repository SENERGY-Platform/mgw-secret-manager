package core

import (
	"secret-manager/internal/model"
	"secret-manager/test"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetMasterKey(t *testing.T) {
	config := test.TestConfig
	config.MasterKeyPath = "/tmp/key"
	config.EncryptionKey = string(test.MasterKey)
	generatedKey, err := CreateAndStoreMasterKey(config)
	if err != nil {
		t.Errorf(err.Error())
	}

	key, err := GetMasterKey(config)
	assert.Equal(t, err, nil)
	assert.Equal(t, generatedKey, key)
}

func TestGenerateMasterKey(t *testing.T) {
	key, err := GenerateMasterKey()
	assert.Equal(t, err, nil)
	assert.Equal(t, len(key), 32)
}

func TestCreateAndStoreMasterKey(t *testing.T) {
	config := test.TestConfig
	config.MasterKeyPath = "/tmp/key"
	config.EncryptionKey = string(test.MasterKey)

	key, err := CreateAndStoreMasterKey(config)
	assert.Equal(t, err, nil)
	assert.Equal(t, len(key), 32)
}

func TestEncryptSecret(t *testing.T) {
	secret := &model.Secret{
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
	secret := &model.Secret{
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
