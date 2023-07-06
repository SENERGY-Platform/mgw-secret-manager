package api

import (
	"github.com/SENERGY-Platform/mgw-secret-manager/internal/config"
	"github.com/SENERGY-Platform/mgw-secret-manager/internal/db"
	"github.com/SENERGY-Platform/mgw-secret-manager/internal/keyHandler"
	"github.com/SENERGY-Platform/mgw-secret-manager/internal/secretHandler"

	"github.com/gin-gonic/gin"
)

type Api struct {
	config        config.Config
	dbHandler     db.Database // sql.db is threadsafe
	secretHandler *secretHandler.SecretHandler
	keyHandler    keyHandler.KeyHandler
}

func New(config config.Config, dbHandler db.Database, secretHandler *secretHandler.SecretHandler, keyHandler keyHandler.KeyHandler) *Api {
	return &Api{
		config:        config,
		dbHandler:     dbHandler,
		secretHandler: secretHandler,
		keyHandler:    keyHandler,
	}
}

func (a *Api) SetRoutes(e *gin.Engine) {
	e.POST("/secrets", a.StoreSecret)
	e.GET("/secrets", a.GetSecrets)
	e.PUT("/secrets/:id", a.UpdateSecret)
	e.GET("/secrets/:id", a.GetSecret)
	e.DELETE("/secrets/:id", a.DeleteSecret)
	e.GET("/types", a.GetTypes)
	e.POST("/load", a.LoadSecretIntoTMPFS)
	e.POST("/key", a.SetEncryptionKey)

	if a.config.EexposeConfidentialEndpoints {
		e.GET("/confidential/secrets/:id", a.GetFullSecret)
	}
}
