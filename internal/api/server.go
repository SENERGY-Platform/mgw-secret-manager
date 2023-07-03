package api

import (
	"fmt"
	"os"

	srv_base "github.com/SENERGY-Platform/go-service-base/srv-base"
	"github.com/SENERGY-Platform/mgw-secret-manager/internal/config"
	"github.com/SENERGY-Platform/mgw-secret-manager/internal/db"
	"github.com/SENERGY-Platform/mgw-secret-manager/internal/secretHandler"

	gin_mw "github.com/SENERGY-Platform/gin-middleware"
	"github.com/gin-gonic/gin"

	"github.com/SENERGY-Platform/mgw-secret-manager/internal/keyHandler"
)

func CORS() gin.HandlerFunc {
	return func(c *gin.Context) {

		fmt.Println(c.Request.Header)
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, Origin, Cache-Control, X-Requested-With")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

func InitServer(config *config.Config) (*gin.Engine, db.Database, secretHandler.SecretHandler) {
	dbHandler, err := db.NewDBHandler(config)
	if err != nil {
		srv_base.Logger.Error(err)
		os.Exit(1)
	}

	gin.SetMode(gin.ReleaseMode)
	apiEngine := gin.New()

	if config.Dev {
		apiEngine.Use(CORS())
	}

	apiEngine.Use(gin_mw.LoggerHandler(srv_base.Logger), gin_mw.ErrorHandler, gin.Recovery())

	apiEngine.UseRawPath = true
	secretHandler := secretHandler.NewSecretHandler(config.EnableEncryption, dbHandler, config.TMPFSPath)
	keyHandler := keyHandler.NewKeyHandler(config.MasterKeyPath, nil)
	Api := New(*config, dbHandler, secretHandler, keyHandler)
	Api.SetRoutes(apiEngine)

	return apiEngine, dbHandler, secretHandler
}
