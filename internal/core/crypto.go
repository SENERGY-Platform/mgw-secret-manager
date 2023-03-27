package core

import (
	"crypto/rand"
	"os"

	"github.com/SENERGY-Platform/mgw-secret-manager/internal/config"
	"github.com/SENERGY-Platform/mgw-secret-manager/internal/crypto"
	"github.com/SENERGY-Platform/mgw-secret-manager/internal/files"
	"github.com/SENERGY-Platform/mgw-secret-manager/internal/model"

	srv_base "github.com/SENERGY-Platform/go-service-base/srv-base"
)

func GetMasterKey(config config.Config, encryptionKey []byte) (decryptedMasterKey []byte, err error) {
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

func CreateAndStoreMasterKey(config config.Config, encryptionKey []byte) (masterKey []byte, err error) {
	masterKey, err = GenerateMasterKey()
	if err != nil {
		return
	}

	encryptedMasterKey, err := crypto.Encrypt(masterKey, encryptionKey)
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
		Name:       secret.Name,
		Value:      encryptedValue,
		SecretType: secret.SecretType,
		ID:         secret.ID,
	}
	return
}

func DecryptSecret(secret *model.EncryptedSecret, key []byte) (decryptedSecret *model.Secret, err error) {
	decryptedValue, err := crypto.Decrypt(secret.Value, key)
	if err != nil {
		return
	}
	decryptedSecret = &model.Secret{
		Name:       secret.Name,
		Value:      string(decryptedValue),
		SecretType: secret.SecretType,
		ID:         secret.ID,
	}
	return
}

func SetEncryptionKey(encryptionKey []byte, config config.Config) (masterKey []byte, err error) {
	if _, err = os.Stat(config.MasterKeyPath); err == nil {
		srv_base.Logger.Debug(("Master Encryption Key found -> Decrypt and Load"))
		masterKey, err = GetMasterKey(config, encryptionKey)
		if err != nil {
			srv_base.Logger.Error(err)
			return
		}
	} else {
		srv_base.Logger.Debug(("Master Encryption Key not found -> Create, Encrypt and Store"))
		masterKey, err = CreateAndStoreMasterKey(config, encryptionKey)
		if err != nil {
			srv_base.Logger.Error(err)
			return
		}
	}
	return
}
