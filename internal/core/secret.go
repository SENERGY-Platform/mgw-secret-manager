package core

import (
	"path/filepath"

	"github.com/SENERGY-Platform/mgw-secret-manager/internal/config"
	"github.com/SENERGY-Platform/mgw-secret-manager/internal/db"
	"github.com/SENERGY-Platform/mgw-secret-manager/internal/files"
	"github.com/SENERGY-Platform/mgw-secret-manager/internal/model"

	srv_base "github.com/SENERGY-Platform/go-service-base/srv-base"
	"github.com/google/uuid"
)

func CreateSecret(name string, value string, secretType string) model.Secret {
	return model.Secret{
		Name:       name,
		Value:      value,
		SecretType: secretType,
		ID:         uuid.New().String(),
	}
}

func StoreSecret(secret *model.Secret, db db.Database, key *[]byte, config config.Config) (err error) {
	srv_base.Logger.Debugf("Store Secret: %s", secret.Name)

	var storedSecret *model.EncryptedSecret

	if config.EnableEncryption {
		storedSecret, err = EncryptSecret(secret, *key)
		if err != nil {
			return
		}
	} else {
		storedSecret = &model.EncryptedSecret{Name: secret.Name, SecretType: secret.SecretType, Value: []byte(secret.Value), ID: secret.ID}
	}

	err = db.SetSecret(storedSecret)
	return
}

func GetSecret(secretName string, db db.Database, key *[]byte, config config.Config) (secret *model.Secret, err error) {
	srv_base.Logger.Debugf("Get Secret: %s from DB", secretName)

	storedSecret, err := db.GetSecret(secretName)
	if err != nil {
		return
	}

	if config.EnableEncryption {
		secret, err = DecryptSecret(storedSecret, *key)
		if err != nil {
			return
		}
		srv_base.Logger.Debugf("Decrypted Secret Value: %s", secret.Value)
	} else {
		secret = &model.Secret{Name: storedSecret.Name, SecretType: storedSecret.SecretType, ID: storedSecret.ID, Value: string(storedSecret.Value)}
	}

	return
}

func LoadSecretToFileSystem(secretName string, db db.Database, config config.Config, key *[]byte) (fileName string, err error) {
	srv_base.Logger.Debugf("Get Secret: %s from DB and load into TMPFS", secretName)

	secret, err := GetSecret(secretName, db, key, config)
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

func GetFullSecrets(db db.Database, config config.Config, key []byte) (secrets []*model.Secret, err error) {
	storedSecrets, err := db.GetSecrets()
	if err != nil {
		return
	}

	for _, storedSecret := range storedSecrets {
		var secret *model.Secret

		if config.EnableEncryption {
			secret, err = DecryptSecret(storedSecret, key)
			if err != nil {
				return
			}
		} else {
			secret = &model.Secret{Name: storedSecret.Name, SecretType: storedSecret.SecretType, ID: storedSecret.ID, Value: string(storedSecret.Value)}
		}

		secrets = append(secrets, secret)
	}
	return
}

func GetSecrets(db db.Database, config config.Config) (secrets []*model.ShortSecret, err error) {
	storedSecrets, err := db.GetSecrets()
	if err != nil {
		return
	}

	for _, storedSecret := range storedSecrets {
		shortSecret := model.ShortSecret{Name: storedSecret.Name, SecretType: storedSecret.SecretType, ID: storedSecret.ID}
		secrets = append(secrets, &shortSecret)
	}
	return
}
