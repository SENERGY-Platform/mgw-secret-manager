package keyHandler

import (
	"testing"

	"github.com/SENERGY-Platform/mgw-secret-manager/internal/config"
	"github.com/SENERGY-Platform/mgw-secret-manager/internal/logger"
	"github.com/SENERGY-Platform/mgw-secret-manager/test"

	"github.com/stretchr/testify/assert"
)

var testConfig, _ = config.NewConfig(config.Flags.ConfPath)
var _, _ = logger.InitLogger(testConfig.Logger)

func TestGetMasterKey(t *testing.T) {
	keyHandler := NewKeyHandler("./key", &test.EncryptionKey)
	generatedKey, err := keyHandler.CreateAndStoreMasterKey()
	if err != nil {
		t.Errorf(err.Error())
	}

	key, err := keyHandler.GetMasterKey()
	assert.Equal(t, err, nil)
	assert.Equal(t, generatedKey, key)
}

func TestGenerateMasterKey(t *testing.T) {
	keyHandler := NewKeyHandler("./key", &test.EncryptionKey)
	key, err := keyHandler.GenerateMasterKey()
	assert.Equal(t, err, nil)
	assert.Equal(t, len(key), 32)
}

func TestCreateAndStoreMasterKey(t *testing.T) {
	keyHandler := NewKeyHandler("./key", &test.EncryptionKey)
	key, err := keyHandler.CreateAndStoreMasterKey()
	assert.Equal(t, err, nil)
	assert.Equal(t, len(key), 32)
}
