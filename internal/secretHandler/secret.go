package secretHandler

import (
	"context"
	"encoding/json"
	"path/filepath"
	"sync"

	"github.com/SENERGY-Platform/mgw-secret-manager/internal/crypto"
	"github.com/SENERGY-Platform/mgw-secret-manager/internal/db"
	"github.com/SENERGY-Platform/mgw-secret-manager/internal/files"
	"github.com/SENERGY-Platform/mgw-secret-manager/internal/logger"
	"github.com/SENERGY-Platform/mgw-secret-manager/internal/models"
	"github.com/SENERGY-Platform/mgw-secret-manager/pkg/api_model"

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

func (secretHandler *SecretHandler) CreateSecret(name string, value string, secretType string, fileName string) api_model.Secret {
	return api_model.Secret{
		Name:       name,
		Value:      value,
		SecretType: secretType,
		ID:         uuid.New().String(),
		FileName:   fileName,
	}
}

func (secretHandler *SecretHandler) SetKey(ctx context.Context, key []byte) {
	logger.Logger.Debugf("Save encryption key in secret handler")

	secretHandler.KeyMutex.Lock()
	secretHandler.Key = key
	secretHandler.KeyMutex.Unlock()
}

func (secretHandler *SecretHandler) StoreSecret(ctx context.Context, secret *api_model.Secret) (err error) {
	logger.Logger.Debugf("Store Secret: %s", secret.ID)

	var storedSecret *models.EncryptedSecret

	if secretHandler.encryptionIsEnabled {
		storedSecret, err = secretHandler.EncryptSecret(secret)
		if err != nil {
			return
		}
	} else {
		storedSecret = &models.EncryptedSecret{
			Name:       secret.Name,
			SecretType: secret.SecretType,
			Value:      []byte(secret.Value),
			ID:         secret.ID,
			FileName:   secret.FileName,
		}
	}

	err = secretHandler.db.SetSecret(ctx, storedSecret)
	return
}

func (secretHandler *SecretHandler) GetFullSecret(ctx context.Context, secretPostRequest api_model.SecretPostRequest) (secret *api_model.Secret, err error) {
	logger.Logger.Debugf("Get full secret")
	storedSecret, err := secretHandler.db.GetSecret(ctx, secretPostRequest.ID)
	if err != nil {
		return
	}

	if secretHandler.encryptionIsEnabled {
		logger.Logger.Debugf("Decrypt Secret Value")
		secret, err = secretHandler.DecryptSecret(storedSecret)
		if err != nil {
			return nil, err
		}
	} else {
		secret = &api_model.Secret{
			Name:       storedSecret.Name,
			Value:      string(storedSecret.Value),
			SecretType: storedSecret.SecretType,
			ID:         storedSecret.ID,
			FileName:   storedSecret.FileName,
		}
	}

	if secretPostRequest.Options != nil {
		// for credentials the user can specify whether the username or password shall be used as secret value
		var credential models.Credential
		err = json.Unmarshal([]byte(secret.Value), &credential)
		if err != nil {
			logger.Logger.Errorf("Credential cant be unmarshaled %s", err.Error())
			return
		}
		if (*secretPostRequest.Options)["from"] == "user" {
			secret.Value = credential.Username
		} else if (*secretPostRequest.Options)["from"] == "password" {
			secret.Value = credential.Password
		}
	}

	return
}

func (secretHandler *SecretHandler) GetSecret(ctx context.Context, secretPostRequest api_model.SecretPostRequest) (shortSecret *api_model.ShortSecret, err error) {
	logger.Logger.Debugf("Get Secret: %s from DB", secretPostRequest.ID)

	secret, err := secretHandler.GetFullSecret(ctx, secretPostRequest)
	if err != nil {
		return
	}
	shortSecret = &api_model.ShortSecret{
		Name:       secret.Name,
		SecretType: secret.SecretType,
		ID:         secret.ID,
		FileName:   secret.FileName,
	}

	return
}

func (secretHandler *SecretHandler) LoadSecretToFileSystem(ctx context.Context, secretPostRequest api_model.SecretPostRequest) (relativeFilePath string, err error) {
	logger.Logger.Debugf("Get Secret and load into TMPFS")

	secret, err := secretHandler.GetFullSecret(ctx, secretPostRequest)
	if err != nil {
		return
	}

	relativeFilePath = filepath.Join(secret.ID, secret.FileName)
	fullOutputPath := filepath.Join(secretHandler.TMPFSPath, relativeFilePath)
	logger.Logger.Debugf("Load Secret: %s to %s", secret.ID, fullOutputPath)

	err = files.WriteToFile(string(secret.Value), fullOutputPath)
	if err != nil {
		logger.Logger.Errorf("Write to TMPFS failed: %s", err.Error())
	}
	return
}

func (secretHandler *SecretHandler) GetSecrets(ctx context.Context) (secrets []*api_model.ShortSecret, err error) {
	logger.Logger.Debugf("Load all short secrets")

	storedSecrets, err := secretHandler.db.GetSecrets(ctx)
	if err != nil {
		return
	}

	for _, storedSecret := range storedSecrets {
		shortSecret := api_model.ShortSecret{
			Name:       storedSecret.Name,
			SecretType: storedSecret.SecretType,
			ID:         storedSecret.ID,
			FileName:   storedSecret.FileName,
		}
		secrets = append(secrets, &shortSecret)
	}
	return
}

func (secretHandler *SecretHandler) EncryptSecret(secret *api_model.Secret) (encryptedSecret *models.EncryptedSecret, err error) {
	encryptedValue, err := crypto.Encrypt([]byte(secret.Value), secretHandler.Key)
	if err != nil {
		return
	}
	encryptedSecret = &models.EncryptedSecret{
		Name:       secret.Name,
		Value:      encryptedValue,
		SecretType: secret.SecretType,
		ID:         secret.ID,
		FileName:   secret.FileName,
	}
	return
}

func (secretHandler *SecretHandler) DecryptSecret(secret *models.EncryptedSecret) (decryptedSecret *api_model.Secret, err error) {
	decryptedValue, err := crypto.Decrypt(secret.Value, secretHandler.Key)
	if err != nil {
		return
	}
	decryptedSecret = &api_model.Secret{
		Name:       secret.Name,
		Value:      string(decryptedValue),
		SecretType: secret.SecretType,
		ID:         secret.ID,
		FileName:   secret.FileName,
	}
	return
}

func (secretHandler *SecretHandler) UpdateSecret(ctx context.Context, secretRequest api_model.SecretRequest, secretID string) (err error) {
	logger.Logger.Debugf("Update secret %s", secretID)

	secret := models.EncryptedSecret{
		Name:       secretRequest.Name,
		Value:      []byte(secretRequest.Value),
		SecretType: secretRequest.SecretType,
		ID:         secretID,
		FileName:   secretRequest.FileName,
	}
	err = secretHandler.db.UpdateSecret(ctx, &secret)

	// TODO if secrets exists in TMP, reload there
	return
}

func (secretHandler *SecretHandler) DeleteSecret(ctx context.Context, secretID string) (err error) {
	logger.Logger.Debugf("Delete secret %s", secretID)
	err = secretHandler.db.DeleteSecret(ctx, secretID)
	return
}
