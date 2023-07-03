package secretHandler

import (
	"path/filepath"
	"sync"

	"github.com/SENERGY-Platform/mgw-secret-manager/internal/crypto"
	"github.com/SENERGY-Platform/mgw-secret-manager/internal/db"
	"github.com/SENERGY-Platform/mgw-secret-manager/internal/files"
	"github.com/SENERGY-Platform/mgw-secret-manager/internal/models"
	"github.com/SENERGY-Platform/mgw-secret-manager/pkg/api_model"

	srv_base "github.com/SENERGY-Platform/go-service-base/srv-base"
	"github.com/google/uuid"
)

type SecretHandler struct {
	encryptionIsEnabled bool
	db                  db.Database
	TMPFSPath           string
	Key                 []byte
	KeyMutex            sync.RWMutex // need a mutex for the master key
}

func NewSecretHandler(encryptionIsEnabled bool, db db.Database, TMPFSPath string) SecretHandler {
	return SecretHandler{
		encryptionIsEnabled: encryptionIsEnabled,
		db:                  db,
		TMPFSPath:           TMPFSPath,
	}
}

func (secretHandler *SecretHandler) CreateSecret(name string, value string, secretType string) models.Secret {
	return models.Secret{
		Name:       name,
		Value:      value,
		SecretType: secretType,
		ID:         uuid.New().String(),
	}
}

func (secretHandler *SecretHandler) SetKey(key []byte) {
	srv_base.Logger.Debugf("Save encryption key in secret handler")

	secretHandler.KeyMutex.Lock()
	secretHandler.Key = key
	secretHandler.KeyMutex.Unlock()
}

func (secretHandler *SecretHandler) StoreSecret(secret *models.Secret) (err error) {
	srv_base.Logger.Debugf("Store Secret: %s", secret.Name)

	var storedSecret *models.EncryptedSecret

	if secretHandler.encryptionIsEnabled {
		storedSecret, err = secretHandler.EncryptSecret(secret)
		if err != nil {
			return
		}
	} else {
		storedSecret = &models.EncryptedSecret{Name: secret.Name, SecretType: secret.SecretType, Value: []byte(secret.Value), ID: secret.ID}
	}

	err = secretHandler.db.SetSecret(storedSecret)
	return
}

func (secretHandler *SecretHandler) GetSecret(secretID string) (secret *models.Secret, err error) {
	srv_base.Logger.Debugf("Get Secret: %s from DB", secretID)

	storedSecret, err := secretHandler.db.GetSecret(secretID)
	if err != nil {
		return
	}

	if secretHandler.encryptionIsEnabled {
		secret, err = secretHandler.DecryptSecret(storedSecret)
		if err != nil {
			return
		}
		srv_base.Logger.Debugf("Decrypted Secret Value: %s", secret.Value)
	} else {
		secret = &models.Secret{Name: storedSecret.Name, SecretType: storedSecret.SecretType, ID: storedSecret.ID, Value: string(storedSecret.Value)}
	}

	return
}

func (secretHandler *SecretHandler) LoadSecretToFileSystem(secretID string) (fileName string, err error) {
	srv_base.Logger.Debugf("Get Secret and load into TMPFS")

	secret, err := secretHandler.GetSecret(secretID)
	if err != nil {
		return
	}

	fileName = secret.ID
	fullOutputPath := filepath.Join(secretHandler.TMPFSPath, fileName)
	srv_base.Logger.Debugf("Load Secret: %s to %s", secret.ID, fullOutputPath)

	err = files.WriteToFile(secret.Value, fullOutputPath)
	if err != nil {
		srv_base.Logger.Errorf("Write to TMPFS failed: %s", err.Error())
	}
	return
}

func (secretHandler *SecretHandler) GetFullSecrets(db db.Database) (secrets []*models.Secret, err error) {
	srv_base.Logger.Debugf("Load all full secrets")

	storedSecrets, err := secretHandler.db.GetSecrets()
	if err != nil {
		return
	}

	for _, storedSecret := range storedSecrets {
		var secret *models.Secret

		if secretHandler.encryptionIsEnabled {
			secret, err = secretHandler.DecryptSecret(storedSecret)
			if err != nil {
				return
			}
		} else {
			secret = &models.Secret{Name: storedSecret.Name, SecretType: storedSecret.SecretType, ID: storedSecret.ID, Value: string(storedSecret.Value)}
		}

		secrets = append(secrets, secret)
	}
	return
}

func (secretHandler *SecretHandler) GetSecrets() (secrets []*api_model.ShortSecret, err error) {
	srv_base.Logger.Debugf("Load all short secrets")

	storedSecrets, err := secretHandler.db.GetSecrets()
	if err != nil {
		return
	}

	for _, storedSecret := range storedSecrets {
		shortSecret := api_model.ShortSecret{Name: storedSecret.Name, SecretType: storedSecret.SecretType, ID: storedSecret.ID}
		secrets = append(secrets, &shortSecret)
	}
	return
}

func (secretHandler *SecretHandler) EncryptSecret(secret *models.Secret) (encryptedSecret *models.EncryptedSecret, err error) {
	encryptedValue, err := crypto.Encrypt([]byte(secret.Value), secretHandler.Key)
	if err != nil {
		return
	}
	encryptedSecret = &models.EncryptedSecret{
		Name:       secret.Name,
		Value:      encryptedValue,
		SecretType: secret.SecretType,
		ID:         secret.ID,
	}
	return
}

func (secretHandler *SecretHandler) DecryptSecret(secret *models.EncryptedSecret) (decryptedSecret *models.Secret, err error) {
	decryptedValue, err := crypto.Decrypt(secret.Value, secretHandler.Key)
	if err != nil {
		return
	}
	decryptedSecret = &models.Secret{
		Name:       secret.Name,
		Value:      string(decryptedValue),
		SecretType: secret.SecretType,
		ID:         secret.ID,
	}
	return
}

func (secretHandler *SecretHandler) UpdateSecret(secretRequest api_model.SecretRequest, secretID string) (err error) {
	srv_base.Logger.Debugf("Update secret %s", secretID)

	secret := models.EncryptedSecret{
		Name:       secretRequest.Name,
		Value:      []byte(secretRequest.Value),
		SecretType: secretRequest.SecretType,
		ID:         secretID,
	}
	err = secretHandler.db.UpdateSecret(&secret)
	return
}

func (secretHandler *SecretHandler) DeleteSecret(secretID string) (err error) {
	srv_base.Logger.Debugf("Delete secret %s", secretID)
	err = secretHandler.db.DeleteSecret(secretID)
	return
}
