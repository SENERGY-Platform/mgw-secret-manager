package http_handler

import (
	"github.com/SENERGY-Platform/mgw-secret-manager/internal/api"
	_ "github.com/SENERGY-Platform/mgw-secret-manager/internal/http_handler/swagger_docs"
	"github.com/SENERGY-Platform/mgw-secret-manager/internal/util"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"path"
)

var routes = []func(api *api.Api) (string, string, gin.HandlerFunc){
	StoreSecret,
	GetSecrets,
	GetShortSecret,
	UpdateSecret,
	DeleteSecret,
	LoadPathVariant,
	InitPathVariant,
	DeleteSecretFromTMPFS,
	CleanReferenceDirectory,
	GetTypes,
	SetEncryptionKey,
	GetSrvInfoH,
}

// SetRoutes
// @title Secret Manager API
// @description Provides access to secret management functions.
// @license.name Apache-2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
func SetRoutes(e *gin.Engine, api *api.Api) {
	if api.Config.ExposeConfidentialEndpoints {
		routes = append(routes, GetValueVariant)
	}
	for _, route := range routes {
		m, p, hf := route(api)
		util.Logger.Debug("set route: " + m + " " + path.Join(e.BasePath(), p))
		e.Handle(m, p, hf)
	}
	e.GET("/health-check", HealthCheck(api))
	e.GET("swagger/*any", ginSwagger.WrapHandler(swaggerFiles.NewHandler(), ginSwagger.InstanceName("standard")))
}

func GetPathFilter() []string {
	return []string{
		"/health-check",
	}
}
