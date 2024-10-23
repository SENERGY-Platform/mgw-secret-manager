package secret_handler

import (
	"github.com/SENERGY-Platform/mgw-secret-manager/internal/crypto"
	"github.com/SENERGY-Platform/mgw-secret-manager/internal/models"
)

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

func (secretHandler *SecretHandler) DecryptSecret(secret *models.EncryptedSecret) (decryptedSecret models.Secret, err error) {
	decryptedValue, err := crypto.Decrypt(secret.Value, secretHandler.Key)
	if err != nil {
		return
	}
	decryptedSecret = models.Secret{
		Name:       secret.Name,
		SecretType: secret.SecretType,
		ID:         secret.ID,
		Value:      string(decryptedValue),
	}
	return
}
