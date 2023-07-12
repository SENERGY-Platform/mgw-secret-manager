package secretHandler

import (
	"github.com/SENERGY-Platform/mgw-secret-manager/internal/crypto"
	"github.com/SENERGY-Platform/mgw-secret-manager/internal/models"
	"github.com/SENERGY-Platform/mgw-secret-manager/pkg/api_model"
)

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
	}
	return
}
