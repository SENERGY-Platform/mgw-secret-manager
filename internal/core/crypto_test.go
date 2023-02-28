package core

import (
	"testing"

	"github.com/SENERGY-Platform/mgw-secret-manager/internal/model"
	"github.com/SENERGY-Platform/mgw-secret-manager/test"

	"github.com/stretchr/testify/assert"
)

func TestGetMasterKey(t *testing.T) {
	config := test.TestConfig
	generatedKey, err := CreateAndStoreMasterKey(config, test.EncryptionKey)
	if err != nil {
		t.Errorf(err.Error())
	}

	key, err := GetMasterKey(config, test.EncryptionKey)
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

	key, err := CreateAndStoreMasterKey(config, test.EncryptionKey)
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
