package api

import (
	"testing"

	srv_base "github.com/SENERGY-Platform/go-service-base/srv-base"
	"github.com/SENERGY-Platform/mgw-secret-manager/internal/config"
	"github.com/SENERGY-Platform/mgw-secret-manager/internal/core"
	"github.com/SENERGY-Platform/mgw-secret-manager/internal/db"
	"github.com/SENERGY-Platform/mgw-secret-manager/internal/model"
	"github.com/SENERGY-Platform/mgw-secret-manager/test"
	"github.com/gin-gonic/gin"
)

var testConfig, _ = config.NewConfig(nil)
var _, _ = srv_base.InitLogger(testConfig.Logger)

func GetTestRouter(enableEncryption bool) (*gin.Engine, *db.Database) {
	apiEngine := gin.New()
	testConfig.EnableEncryption = enableEncryption
	var dbHandler, _ = db.GetTestDB(testConfig)
	Api := New(*testConfig, dbHandler)
	Api.masterKey = &test.MasterKey
	Api.SetRoutes(apiEngine)

	return apiEngine, &dbHandler
}

func SetupDummySecret(t *testing.T, name string, value string, secretType string, dbHandler *db.Database) (model.Secret, model.ShortSecret) {
	secret := core.CreateSecret(name, value, secretType)
	err := core.StoreSecret(&secret, dbHandler, &test.MasterKey, *testConfig)
	if err != nil {
		t.Errorf(err.Error())
	}

	return secret, model.ShortSecret{Name: name, SecretType: secretType, ID: secret.ID}
}
