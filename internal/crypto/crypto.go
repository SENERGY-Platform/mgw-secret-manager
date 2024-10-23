package crypto

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"github.com/SENERGY-Platform/mgw-secret-manager/internal/util"
	"io"
)

func CreatGCM(block cipher.Block) (gcm cipher.AEAD, err error) {
	// Creating GCM mode
	gcm, err = cipher.NewGCM(block)
	if err != nil {
		util.Logger.Errorf("cipher GCM err: %v", err.Error())
		return
	}
	return
}

func NewCipher(key []byte) (block cipher.Block, err error) {
	// Creating block of algorithm
	block, err = aes.NewCipher(key)
	if err != nil {
		util.Logger.Errorf("cipher err: %v", err.Error())
		return
	}
	return
}

func Decrypt(cipherText []byte, key []byte) (plain []byte, err error) {
	block, err := NewCipher(key)
	if err != nil {
		return
	}

	gcm, err := CreatGCM(block)
	if err != nil {
		return
	}

	// Deattached nonce and decrypt
	nonce := cipherText[:gcm.NonceSize()]
	cipherTextWithoutNonce := cipherText[gcm.NonceSize():]
	plain, err = gcm.Open(nil, nonce, cipherTextWithoutNonce, nil)
	if err != nil {
		util.Logger.Errorf("decrypt file err: %v", err.Error())
		return
	}

	return
}

func Encrypt(plain []byte, key []byte) (cipherBytes []byte, err error) {
	block, err := NewCipher(key)
	if err != nil {
		util.Logger.Errorf((err.Error()))
		return
	}
	gcm, err := CreatGCM(block)
	if err != nil {
		util.Logger.Errorf((err.Error()))
		return
	}

	// Generating random nonce
	nonce := make([]byte, gcm.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		util.Logger.Errorf("nonce  err: %v", err.Error())
		return
	}

	// Decrypt file
	cipherBytes = gcm.Seal(nonce, nonce, plain, nil)
	return
}
