package keyHandler

import (
	"testing"

	"github.com/SENERGY-Platform/mgw-secret-manager/internal/model"
	"github.com/SENERGY-Platform/mgw-secret-manager/test"

	"github.com/stretchr/testify/assert"
)

func TestGetMasterKey(t *testing.T) {
	generatedKey, err := CreateAndStoreMasterKey(*testConfig, test.EncryptionKey)
	if err != nil {
		t.Errorf(err.Error())
	}

	key, err := GetMasterKey(*testConfig, test.EncryptionKey)
	assert.Equal(t, err, nil)
	assert.Equal(t, generatedKey, key)
}

func TestGenerateMasterKey(t *testing.T) {
	key, err := GenerateMasterKey()
	assert.Equal(t, err, nil)
	assert.Equal(t, len(key), 32)
}

func TestCreateAndStoreMasterKey(t *testing.T) {
	key, err := CreateAndStoreMasterKey(*testConfig, test.EncryptionKey)
	assert.Equal(t, err, nil)
	assert.Equal(t, len(key), 32)
}
