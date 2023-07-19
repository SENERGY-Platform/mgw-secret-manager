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

const SecretsPath = "/secrets"
const PathVariantPath = "/path-variant"
const SecretPath = SecretsPath + "/:id"
const LoadPathVariantPath = PathVariantPath + "/load"
const InitPathVariantPath = PathVariantPath + "/init"
const UnLoadPathVariantPath = PathVariantPath + "/unload"

const ConfidentialPath = "/confidential"
const ValueVariantPath = ConfidentialPath + "/value-variant"

func (a *Api) SetRoutes(e *gin.Engine) {
	e.POST(SecretsPath, a.StoreSecret)
	e.GET(SecretsPath, a.GetSecrets)

	e.PUT(SecretPath, a.UpdateSecret)
	e.DELETE(SecretPath, a.DeleteSecret)

	e.POST(LoadPathVariantPath, a.LoadPathVariant)
	e.POST(InitPathVariantPath, a.InitPathVariant)
	e.POST(UnLoadPathVariantPath, a.DeleteSecretFromTMPFS)

	e.GET("/types", a.GetTypes)

	e.POST("/key", a.SetEncryptionKey)

	if a.config.ExposeConfidentialEndpoints {
		e.POST(ValueVariantPath, a.GetValueVariant)
	}
}
