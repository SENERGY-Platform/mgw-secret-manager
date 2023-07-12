package secretHandler

import (
	"context"
	"encoding/json"

	"github.com/SENERGY-Platform/mgw-secret-manager/internal/logger"
	"github.com/SENERGY-Platform/mgw-secret-manager/internal/models"
	"github.com/SENERGY-Platform/mgw-secret-manager/pkg/api_model"
)

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
		}
		secrets = append(secrets, &shortSecret)
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
		Path:       secretHandler.BuildTMPFSOutputPath(secretPostRequest),
	}

	return
}

func (secretHandler *SecretHandler) ExtractValue(ctx context.Context, secretPostRequest api_model.SecretPostRequest, secret models.EncryptedSecret) (value string, err error) {
	if secretPostRequest.Item == nil {
		return string(secret.Value), nil
	}

	var secretValue models.SecretValue
	err = json.Unmarshal(secret.Value, &secretValue)
	if err != nil {
		logger.Logger.Errorf("Secret can not be unmarshaled: %s This can be caused by specifing an Item on a secret that is not saved in JSON", err.Error())
		return
	}

	itemKey := *secretPostRequest.Item
	val, ok := secretValue[itemKey]
	if !ok {
		logger.Logger.Errorf("Item %s does not exist as key in JSON secret %s", itemKey, secret.ID)
		err = nil
		return
	}
	return val, nil
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
		}
	}

	secret.Path = secretHandler.BuildTMPFSOutputPath(secretPostRequest)
	secret.Value, err = secretHandler.ExtractValue(ctx, secretPostRequest, *storedSecret)
	return
}
