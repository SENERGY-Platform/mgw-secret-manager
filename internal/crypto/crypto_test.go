package crypto

import (
	"testing"

	srv_base "github.com/SENERGY-Platform/go-service-base/srv-base"

	"github.com/SENERGY-Platform/mgw-secret-manager/test"
	"github.com/stretchr/testify/assert"
)

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
	var _, _ = srv_base.InitLogger(test.TestConfig.Logger)
	key := make([]byte, 12)

	_, err := NewCipher(key)
	assert.NotEqual(t, err, nil)

}
