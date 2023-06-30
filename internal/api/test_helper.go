package api

import (
	"testing"

	srv_base "github.com/SENERGY-Platform/go-service-base/srv-base"
	"github.com/SENERGY-Platform/mgw-secret-manager/internal/config"
	"github.com/SENERGY-Platform/mgw-secret-manager/internal/db"
	"github.com/SENERGY-Platform/mgw-secret-manager/internal/models"
	"github.com/SENERGY-Platform/mgw-secret-manager/internal/secretHandler"

	"github.com/SENERGY-Platform/mgw-secret-manager/pkg/api_model"
	"github.com/SENERGY-Platform/mgw-secret-manager/test"
	"github.com/gin-gonic/gin"
)

var testConfig, _ = config.NewConfig(nil)
var _, _ = srv_base.InitLogger(testConfig.Logger)

func GetTestRouter(enableEncryption bool) (*gin.Engine, db.Database) {
	apiEngine := gin.New()
	testConfig.EnableEncryption = enableEncryption
	var dbHandler, _ = db.NewDBHandler(testConfig)
	dbHandler.Cleanup()
	Api := New(*testConfig, dbHandler)
	Api.masterKey = &test.MasterKey
	Api.SetRoutes(apiEngine)

	return apiEngine, dbHandler
}

func SetupDummySecret(t *testing.T, name string, value string, secretType string, dbHandler db.Database) (models.Secret, api_model.ShortSecret) {
	secret := secretHandler.CreateSecret(name, value, secretType)
	err := secretHandler.StoreSecret(&secret, dbHandler, &test.MasterKey, testConfig.EnableEncryption)
	if err != nil {
		t.Errorf(err.Error())
	}

	return secret, api_model.ShortSecret{Name: name, SecretType: secretType, ID: secret.ID}
}
