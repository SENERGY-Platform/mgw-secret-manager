package crypto

import (
	"testing"

	"github.com/SENERGY-Platform/mgw-secret-manager/internal/config"
	"github.com/SENERGY-Platform/mgw-secret-manager/internal/logger"
	"github.com/stretchr/testify/assert"
)

var testConfig, _ = config.NewConfig(config.Flags.ConfPath)
var _, _ = logger.InitLogger(testConfig.Logger)

func TestEncrytpionAndDecryption(t *testing.T) {
	plainText := "Test text"
	key := make([]byte, 32)
	encrypted, err := Encrypt([]byte(plainText), key)
	if err != nil {
		t.Errorf(err.Error())
	}
	decryptedBytes, err := Decrypt(encrypted, key)
	decryptedText := string(decryptedBytes)
	assert.Equal(t, plainText, decryptedText)
}

func TestGoodKey(t *testing.T) {
	key := make([]byte, 32)
	_, err := NewCipher(key)
	assert.Equal(t, err, nil)
}

func TestBadKey(t *testing.T) {
	key := make([]byte, 12)

	_, err := NewCipher(key)
	assert.NotEqual(t, err, nil)

}
