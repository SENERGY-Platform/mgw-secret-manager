package api

import (
	srv_base "github.com/SENERGY-Platform/go-service-base/srv-base"
	"github.com/SENERGY-Platform/mgw-secret-manager/internal/db"
	"github.com/SENERGY-Platform/mgw-secret-manager/test"
	"github.com/gin-gonic/gin"
)

var _, _ = srv_base.InitLogger(test.TestConfig.Logger)

func GetTestRouter() (*gin.Engine, *db.DBHandler) {
	apiEngine := gin.New()
	var dbHandler, _ = db.NewDBHandler(test.TestConfig)
	Api := New(test.TestConfig, dbHandler)
	Api.masterKey = &test.MasterKey
	Api.SetRoutes(apiEngine)

	return apiEngine, dbHandler
}
