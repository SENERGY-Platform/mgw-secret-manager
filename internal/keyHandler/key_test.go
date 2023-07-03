package keyHandler

import (
	"testing"

	srv_base "github.com/SENERGY-Platform/go-service-base/srv-base"
	"github.com/SENERGY-Platform/mgw-secret-manager/internal/config"
	"github.com/SENERGY-Platform/mgw-secret-manager/test"

	"github.com/stretchr/testify/assert"
)

var testConfig, _ = config.NewConfig(nil)
var _, err = srv_base.InitLogger(testConfig.Logger)

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
