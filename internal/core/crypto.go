package core

import (
	"crypto/rand"
	"os"
	"secret-manager/internal/config"
	"secret-manager/internal/crypto"
	"secret-manager/internal/files"
	"secret-manager/internal/model"
)

func GetMasterKey(config config.Config) (decryptedMasterKey []byte, err error) {
	encryptionKey := []byte(config.EncryptionKey)
	encryptedMasterKey, err := os.ReadFile(config.MasterKeyPath)
	if err != nil {
		return
	}

	decryptedMasterKey, err = crypto.Decrypt(encryptedMasterKey, encryptionKey)
	return
}

func GenerateMasterKey() (key []byte, err error) {
	key = make([]byte, 32) //generate a random 32 byte key for AES-256
	if _, err = rand.Read(key); err != nil {
		return
	}

	return
}

func CreateAndStoreMasterKey(config config.Config) (masterKey []byte, err error) {
	masterKey, err = GenerateMasterKey()
	if err != nil {
		return
	}

	encryptedMasterKey, err := crypto.Encrypt(masterKey, []byte(config.EncryptionKey))
	if err != nil {
		return
	}
	err = files.WriteBytesToFile(encryptedMasterKey, config.MasterKeyPath)

	return
}

func EncryptSecret(secret *model.Secret, key []byte) (encryptedSecret *model.EncryptedSecret, err error) {
	encryptedValue, err := crypto.Encrypt([]byte(secret.Value), key)
	if err != nil {
		return
	}
	encryptedSecret = &model.EncryptedSecret{
		Name:  secret.Name,
		Value: encryptedValue,
	}
	return
}

func DecryptSecret(secret *model.EncryptedSecret, key []byte) (decryptedSecret *model.Secret, err error) {
	decryptedValue, err := crypto.Decrypt(secret.Value, key)
	if err != nil {
		return
	}
	decryptedSecret = &model.Secret{
		Name:  secret.Name,
		Value: string(decryptedValue),
		ID:    secret.ID,
	}
	return
}
