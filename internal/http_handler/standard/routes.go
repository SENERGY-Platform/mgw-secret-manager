package standard

import (
	gin_mw "github.com/SENERGY-Platform/gin-middleware"
	"github.com/SENERGY-Platform/mgw-secret-manager/internal/api"
	"github.com/SENERGY-Platform/mgw-secret-manager/internal/http_handler/shared"
	_ "github.com/SENERGY-Platform/mgw-secret-manager/internal/http_handler/swagger_docs"
	"github.com/SENERGY-Platform/mgw-secret-manager/internal/util"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

var routes = gin_mw.Routes[*api.Api]{
	LoadPathVariant,
	InitPathVariant,
	DeleteSecretFromTMPFS,
	CleanReferenceDirectory,
	HealthCheck,
}

// SetRoutes
// @title Secret Manager API
// @description Provides access to secret management functions.
// @license.name Apache-2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @BasePath /
func SetRoutes(e *gin.Engine, api *api.Api) error {
	rg := e.Group("")
	if api.Config.ExposeConfidentialEndpoints {
		routes = append(routes, GetValueVariant)
	}
	routes = append(routes, shared.Routes...)
	err := routes.Set(api, rg, util.Logger)
	if err != nil {
		return err
	}
	rg.GET("swagger/*any", ginSwagger.WrapHandler(swaggerFiles.NewHandler(), ginSwagger.InstanceName("standard")))
	return nil
}
