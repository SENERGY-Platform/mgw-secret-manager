package core

import (
	"path/filepath"
	"secret-manager/internal/config"
	"secret-manager/internal/db"
	"secret-manager/internal/files"
	"secret-manager/internal/model"

	srv_base "github.com/SENERGY-Platform/go-service-base/srv-base"
	"github.com/google/uuid"
)

func CreateSecret(name string, value string) model.Secret {
	return model.Secret{
		Name:  name,
		Value: value,
		ID:    uuid.New().String(),
	}
}

func StoreSecret(secret *model.Secret, db db.Database, key []byte) (err error) {
	srv_base.Logger.Debugf("Store Secret: %s", secret.Name)
	encryptedSecret, err := EncryptSecret(secret, key)
	if err != nil {
		return
	}

	err = db.SetSecret(encryptedSecret)
	return
}

func GetSecret(secretName string, db db.Database, key []byte) (decryptedSecret *model.Secret, err error) {
	srv_base.Logger.Debugf("Get Secret: %s from DB", secretName)

	encryptedSecret, err := db.GetSecret(secretName)
	if err != nil {
		return
	}
	decryptedSecret, err = DecryptSecret(encryptedSecret, key)
	srv_base.Logger.Debugf("Decrypted Secret Value: %s", decryptedSecret.Value)

	return
}

func LoadSecretToFileSystem(secretName string, db db.Database, config config.Config, key []byte) (fileName string, err error) {
	srv_base.Logger.Debugf("Get Secret: %s from DB and load into TMPFS", secretName)

	secret, err := GetSecret(secretName, db, key)
	if err != nil {
		return
	}

	fileName = secret.ID
	fullOutputPath := filepath.Join(config.TMPFSPath, fileName)
	srv_base.Logger.Debugf("Load Secret: %s to %s", secretName, fullOutputPath)

	err = files.WriteToFile(secret.Value, fullOutputPath)
	if err != nil {
		srv_base.Logger.Errorf("Write to TMPFS failed: %s", err.Error())
	}
	return
}
