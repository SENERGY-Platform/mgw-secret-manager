package api

import (
	"github.com/SENERGY-Platform/mgw-secret-manager/internal/db"
	"github.com/SENERGY-Platform/mgw-secret-manager/test"

	srv_base "github.com/SENERGY-Platform/go-service-base/srv-base"

	"github.com/gin-gonic/gin"
)

var _, _ = srv_base.InitLogger(test.TestConfig.Logger)

var dbHandler, _ = db.NewDBHandler(test.TestConfig)

func GetTestRouter() *gin.Engine {
	apiEngine := gin.New()
	Api := New(test.TestConfig, dbHandler)
	Api.masterKey = &test.MasterKey
	Api.SetRoutes(apiEngine)

	return apiEngine
}
