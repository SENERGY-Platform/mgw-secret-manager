package keyHandler

import (
	"crypto/rand"
	"os"

	"github.com/SENERGY-Platform/mgw-secret-manager/internal/crypto"
	"github.com/SENERGY-Platform/mgw-secret-manager/internal/files"
	"github.com/SENERGY-Platform/mgw-secret-manager/internal/logger"
	"github.com/SENERGY-Platform/mgw-secret-manager/internal/secretHandler"
)

type KeyHandler struct {
	MasterKeyPath string
	EncryptionKey *[]byte
}

func NewKeyHandler(masterKeyPath string, encryptionKey *[]byte) (keyHandler KeyHandler) {
	return KeyHandler{
		MasterKeyPath: masterKeyPath,
		EncryptionKey: encryptionKey,
	}
}

func (keyHandler *KeyHandler) GetMasterKey() (decryptedMasterKey []byte, err error) {
	encryptedMasterKey, err := os.ReadFile(keyHandler.MasterKeyPath)
	if err != nil {
		return
	}

	decryptedMasterKey, err = crypto.Decrypt(encryptedMasterKey, *keyHandler.EncryptionKey)
	return
}

func (keyHandler *KeyHandler) GenerateMasterKey() (key []byte, err error) {
	key = make([]byte, 32) //generate a random 32 byte key for AES-256
	if _, err = rand.Read(key); err != nil {
		return
	}

	return
}

func (keyHandler *KeyHandler) CreateAndStoreMasterKey() (masterKey []byte, err error) {
	masterKey, err = keyHandler.GenerateMasterKey()
	if err != nil {
		return
	}

	encryptedMasterKey, err := crypto.Encrypt(masterKey, *keyHandler.EncryptionKey)
	if err != nil {
		return
	}
	err = files.WriteBytesToFile(encryptedMasterKey, keyHandler.MasterKeyPath)

	return
}

func (keyHandler *KeyHandler) SetEncryptionKey(encryptionKey []byte, secretHandler *secretHandler.SecretHandler) error {
	keyHandler.EncryptionKey = &encryptionKey
	var masterKey []byte

	if _, err := os.Stat(keyHandler.MasterKeyPath); err == nil {
		logger.Logger.Debug(("Master Encryption Key found -> Decrypt and Load"))
		masterKey, err = keyHandler.GetMasterKey()
		if err != nil {
			logger.Logger.Error(err)
			return err
		}
	} else {
		logger.Logger.Debug(("Master Encryption Key not found -> Create, Encrypt and Store"))
		masterKey, err = keyHandler.CreateAndStoreMasterKey()
		if err != nil {
			logger.Logger.Error(err)
			return err
		}
	}

	secretHandler.SetKey(masterKey)
	return nil
}
