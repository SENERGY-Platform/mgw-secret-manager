package crypto

import (
	"testing"

	srv_base "github.com/SENERGY-Platform/go-service-base/srv-base"

	"github.com/SENERGY-Platform/mgw-secret-manager/internal/config"
	"github.com/stretchr/testify/assert"
)

var testConfig, _ = config.NewConfig(nil)
var _, _ = srv_base.InitLogger(testConfig.Logger)

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
