package server

import (
	"fmt"
	"os"

	"github.com/SENERGY-Platform/mgw-secret-manager/internal/api"

	"github.com/SENERGY-Platform/mgw-secret-manager/internal/config"
	"github.com/SENERGY-Platform/mgw-secret-manager/internal/db"
	"github.com/SENERGY-Platform/mgw-secret-manager/internal/httpHandler"
	"github.com/SENERGY-Platform/mgw-secret-manager/internal/logger"
	"github.com/SENERGY-Platform/mgw-secret-manager/internal/secretHandler"
	srv_info_hdl "github.com/SENERGY-Platform/go-service-base/srv-info-hdl"

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

func InitServer(config *config.Config, version string) (*gin.Engine, db.Database, secretHandler.SecretHandler) {
	srvInfoHdl := srv_info_hdl.New("secret-manager", version)

	dbHandler, err := db.NewDBHandler(config)
	if err != nil {
		logger.Logger.Error(err)
		os.Exit(1)
	}

	secretHandler := secretHandler.NewSecretHandler(config.EnableEncryption, dbHandler, config.TMPFSPath)
	keyHandler := keyHandler.NewKeyHandler(config.MasterKeyPath, nil)

	gin.SetMode(gin.ReleaseMode)
	engine := gin.New()

	if config.Dev {
		engine.Use(CORS())
	}

	staticHeader := map[string]string{
		"X-Api-Version": version,
	}

	engine.Use(
		gin_mw.StaticHeaderHandler(staticHeader),
		gin_mw.LoggerHandler(logger.Logger, nil, func(gc *gin.Context) string {
			return requestid.Get(gc)
		}),
		gin_mw.ErrorHandler(httpHandler.GetStatusCode, ", "),
		requestid.New(),
		gin.Recovery(),
	)
	engine.UseRawPath = true

	api := api.New(*config, dbHandler, &secretHandler, keyHandler, srvInfoHdl)
	httpHandler.SetRoutes(engine, api)

	return engine, dbHandler, secretHandler
}
