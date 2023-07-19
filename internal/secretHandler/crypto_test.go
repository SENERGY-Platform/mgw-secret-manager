package secretHandler

import (
	"context"
	"testing"

	"github.com/SENERGY-Platform/mgw-secret-manager/internal/config"
	"github.com/SENERGY-Platform/mgw-secret-manager/internal/db"
	"github.com/SENERGY-Platform/mgw-secret-manager/internal/logger"
	"github.com/SENERGY-Platform/mgw-secret-manager/internal/models"

	"github.com/SENERGY-Platform/mgw-secret-manager/test"

	"github.com/stretchr/testify/assert"
)

var testConfig, _ = config.NewConfig(config.Flags.ConfPath)
var _, _ = logger.InitLogger(testConfig.Logger)

func TestEncryptDecryptSecret(t *testing.T) {
	var dbHandler, _ = db.NewDBHandler(testConfig)
	ctx := context.Background()

	secretHandler := NewSecretHandler(false, dbHandler, ".")
	secretHandler.SetKey(ctx, test.MasterKey)
	secret := &models.Secret{
		Name:  "Test",
		Value: "value",
	}
	encryptedSecret, err := secretHandler.EncryptSecret(secret)
	if err != nil {
		t.Errorf(err.Error())
	}
	assert.Equal(t, encryptedSecret.Name, secret.Name)

	decryptedSecret, err := secretHandler.DecryptSecret(encryptedSecret)
	if err != nil {
		t.Errorf(err.Error())
	}

	assert.Equal(t, decryptedSecret.Name, secret.Name)
	assert.Equal(t, decryptedSecret.Value, secret.Value)
}
