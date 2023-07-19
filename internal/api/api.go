package api

import (
	"github.com/SENERGY-Platform/mgw-secret-manager/internal/config"
	"github.com/SENERGY-Platform/mgw-secret-manager/internal/db"
	"github.com/SENERGY-Platform/mgw-secret-manager/internal/keyHandler"
	"github.com/SENERGY-Platform/mgw-secret-manager/internal/secretHandler"
	"github.com/SENERGY-Platform/mgw-secret-manager/pkg/api_model"

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
	e.POST(api_model.SecretsPath, a.StoreSecret)
	e.GET(api_model.SecretsPath, a.GetSecrets)

	e.PUT(api_model.SecretPath, a.UpdateSecret)
	e.DELETE(api_model.SecretPath, a.DeleteSecret)

	e.POST(api_model.LoadPathVariantPath, a.LoadPathVariant)
	e.POST(api_model.InitPathVariantPath, a.InitPathVariant)
	e.POST(api_model.UnLoadPathVariantPath, a.DeleteSecretFromTMPFS)

	e.GET("/types", a.GetTypes)

	e.POST("/key", a.SetEncryptionKey)

	if a.config.ExposeConfidentialEndpoints {
		e.POST(api_model.ValueVariantPath, a.GetValueVariant)
	}
}
