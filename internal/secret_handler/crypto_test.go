package secret_handler

import (
	"context"
	"github.com/SENERGY-Platform/mgw-secret-manager/internal/db"
	"github.com/SENERGY-Platform/mgw-secret-manager/internal/models"
	"github.com/SENERGY-Platform/mgw-secret-manager/internal/util"
	"github.com/stretchr/testify/assert"
	"testing"
)

var MasterKey []byte = make([]byte, 32)

var testConfig, _ = util.NewConfig(util.Flags.ConfPath)
var _, _ = util.InitLogger(testConfig.Logger)

func TestEncryptDecryptSecret(t *testing.T) {
	var dbHandler, _ = db.NewDBHandler(testConfig)
	ctx := context.Background()

	secretHandler := NewSecretHandler(false, dbHandler, ".")
	secretHandler.SetKey(ctx, MasterKey)
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
