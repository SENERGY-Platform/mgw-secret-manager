package httpHandler

import (
	"github.com/SENERGY-Platform/mgw-secret-manager/internal/api"
	"github.com/SENERGY-Platform/mgw-secret-manager/pkg/api_model"
	"github.com/gin-gonic/gin"
)

func SetRoutes(e *gin.Engine, api *api.Api) {
	e.POST(api_model.SecretsPath, StoreSecret(api))

	e.GET(api_model.SecretsPath, GetSecrets(api))

	e.GET(api_model.SecretPath, GetShortSecret(api))
	e.PUT(api_model.SecretPath, UpdateSecret(api))
	e.DELETE(api_model.SecretPath, DeleteSecret(api))

	e.POST(api_model.LoadPathVariantPath, LoadPathVariant(api))
	e.POST(api_model.InitPathVariantPath, InitPathVariant(api))
	e.POST(api_model.UnLoadPathVariantPath, DeleteSecretFromTMPFS(api))
	e.POST(api_model.CleanPathVariantsPath, CleanReferenceDirectory(api))

	e.GET("/types", GetTypes(api))

	e.POST("/key", SetEncryptionKey(api))

	if api.Config.ExposeConfidentialEndpoints {
		e.POST(api_model.ValueVariantPath, GetValueVariant(api))
	}

	e.GET("/health-check", HealthCheck(api))
	e.GET("/info", GetSrvInfoH(api))
}
