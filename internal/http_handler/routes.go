package http_handler

import (
	"github.com/SENERGY-Platform/mgw-secret-manager/internal/api"
	"github.com/SENERGY-Platform/mgw-secret-manager/internal/util"
	"github.com/gin-gonic/gin"
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
}

func GetPathFilter() []string {
	return []string{
		"/health-check",
	}
}
