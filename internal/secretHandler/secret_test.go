package secretHandler

import (
	"testing"

	"github.com/SENERGY-Platform/mgw-secret-manager/internal/config"
	"github.com/SENERGY-Platform/mgw-secret-manager/internal/db"
	"github.com/SENERGY-Platform/mgw-secret-manager/internal/models"

	"github.com/SENERGY-Platform/mgw-secret-manager/test"

	srv_base "github.com/SENERGY-Platform/go-service-base/srv-base"

	"github.com/stretchr/testify/assert"
)

var testConfig, _ = config.NewConfig(nil)
var _, err = srv_base.InitLogger(testConfig.Logger)

func TestEncryptDecryptSecret(t *testing.T) {
	var dbHandler, _ = db.NewDBHandler(testConfig)

	secretHandler := NewSecretHandler(false, dbHandler, ".")
	secretHandler.SetKey(test.MasterKey)
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
