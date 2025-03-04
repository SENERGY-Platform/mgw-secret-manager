package standard

import (
	"github.com/SENERGY-Platform/mgw-secret-manager/internal/api"
	"github.com/SENERGY-Platform/mgw-secret-manager/pkg/api_model"
	"github.com/gin-gonic/gin"
	"net/http"
)

func HealthCheck(api *api.Api) (string, string, gin.HandlerFunc) {
	return http.MethodGet, api_model.HealthCheckPath, func(gc *gin.Context) {
		_, err := api.DbHandler.GetSecrets(gc.Request.Context())
		if err != nil {
			gc.JSON(http.StatusInternalServerError, nil)
		}
		gc.JSON(http.StatusOK, nil)
	}
}
