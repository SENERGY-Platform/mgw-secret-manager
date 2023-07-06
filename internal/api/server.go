package api

import (
	"fmt"
	"os"

	"github.com/SENERGY-Platform/mgw-secret-manager/internal/config"
	"github.com/SENERGY-Platform/mgw-secret-manager/internal/db"
	"github.com/SENERGY-Platform/mgw-secret-manager/internal/logger"
	"github.com/SENERGY-Platform/mgw-secret-manager/internal/secretHandler"
	"github.com/gin-contrib/requestid"

	gin_mw "github.com/SENERGY-Platform/gin-middleware"
	"github.com/gin-gonic/gin"

	"github.com/SENERGY-Platform/mgw-secret-manager/internal/keyHandler"
)

func CORS() gin.HandlerFunc {
	return func(c *gin.Context) {

		fmt.Println(c.Request.Header)
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "DELETE, PUT, POST")
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
		logger.Logger.Error(err)
		os.Exit(1)
	}

	gin.SetMode(gin.ReleaseMode)
	engine := gin.New()

	if config.Dev {
		engine.Use(CORS())
	}

	engine.Use(gin_mw.LoggerHandler(logger.Logger), gin_mw.ErrorHandler, requestid.New(), gin.Recovery())

	engine.UseRawPath = true
	secretHandler := secretHandler.NewSecretHandler(config.EnableEncryption, dbHandler, config.TMPFSPath)
	keyHandler := keyHandler.NewKeyHandler(config.MasterKeyPath, nil)
	Api := New(*config, dbHandler, &secretHandler, keyHandler)
	Api.SetRoutes(engine)

	return engine, dbHandler, secretHandler
}
