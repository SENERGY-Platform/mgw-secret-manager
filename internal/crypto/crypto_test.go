package crypto

import (
	"testing"

	"github.com/SENERGY-Platform/mgw-secret-manager/internal/util"
	"github.com/stretchr/testify/assert"
)

var testConfig, _ = util.NewConfig(util.Flags.ConfPath)
var _, _ = util.InitLogger(testConfig.Logger)

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
