package key_handler

import (
	"github.com/SENERGY-Platform/mgw-secret-manager/internal/util"
	"github.com/stretchr/testify/assert"
	"testing"
)

var EncryptionKey = []byte("eShVmYq3t6w9z$C&E)H@McQfTjWnZr4u")

var testConfig, _ = util.NewConfig(util.Flags.ConfPath)
var _, _ = util.InitLogger(testConfig.Logger)

func TestGetMasterKey(t *testing.T) {
	keyHandler := NewKeyHandler("./key", &EncryptionKey)
	generatedKey, err := keyHandler.CreateAndStoreMasterKey()
	if err != nil {
		t.Errorf(err.Error())
	}

	key, err := keyHandler.GetMasterKey()
	assert.Equal(t, err, nil)
	assert.Equal(t, generatedKey, key)
}

func TestGenerateMasterKey(t *testing.T) {
	keyHandler := NewKeyHandler("./key", &EncryptionKey)
	key, err := keyHandler.GenerateMasterKey()
	assert.Equal(t, err, nil)
	assert.Equal(t, len(key), 32)
}

func TestCreateAndStoreMasterKey(t *testing.T) {
	keyHandler := NewKeyHandler("./key", &EncryptionKey)
	key, err := keyHandler.CreateAndStoreMasterKey()
	assert.Equal(t, err, nil)
	assert.Equal(t, len(key), 32)
}
