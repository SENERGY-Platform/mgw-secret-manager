package http_handler

import (
	"errors"
	gin_mw "github.com/SENERGY-Platform/gin-middleware"
	"github.com/SENERGY-Platform/mgw-secret-manager/internal/api"
	"github.com/SENERGY-Platform/mgw-secret-manager/internal/custom_errors"
	"github.com/SENERGY-Platform/mgw-secret-manager/internal/http_handler/restricted"
	"github.com/SENERGY-Platform/mgw-secret-manager/internal/http_handler/standard"
	"github.com/SENERGY-Platform/mgw-secret-manager/internal/util"
	"github.com/SENERGY-Platform/mgw-secret-manager/pkg/api_model"
	"github.com/gin-contrib/requestid"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetPathFilter() []string {
	return []string{
		"/health-check",
	}
}

func New(a *api.Api, staticHeader map[string]string) (*gin.Engine, error) {
	gin.SetMode(gin.ReleaseMode)
	httpHandler := gin.New()
	httpHandler.Use(gin_mw.StaticHeaderHandler(staticHeader), requestid.New(requestid.WithCustomHeaderStrKey("X-Request-ID")), gin_mw.LoggerHandler(util.Logger, []string{api_model.HealthCheckPath}, func(gc *gin.Context) string {
		return requestid.Get(gc)
	}), gin_mw.ErrorHandler(GetStatusCode, ", "), gin.Recovery())
	httpHandler.UseRawPath = true
	err := standard.SetRoutes(httpHandler, a)
	if err != nil {
		return nil, err
	}
	err = restricted.SetRoutes(httpHandler, a)
	if err != nil {
		return nil, err
	}
	return httpHandler, nil
}

func GetStatusCode(err error) int {
	var nsf *custom_errors.NoSecretFound
	if errors.As(err, &nsf) {
		return http.StatusNotFound
	}
	var mek *custom_errors.MissingEncryptionKey
	if errors.As(err, &mek) {
		return http.StatusBadRequest
	}
	var sal *custom_errors.SecretAlreadyLoaded
	if errors.As(err, &sal) {
		return http.StatusBadRequest
	}
	var snf *custom_errors.SecretDoesNotExistsInFilesystem
	if errors.As(err, &snf) {
		return http.StatusNotFound
	}
	return http.StatusInternalServerError
}
