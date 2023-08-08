package httpHandler

import (
	"net/http"

	"github.com/SENERGY-Platform/mgw-secret-manager/internal/api"
	"github.com/gin-gonic/gin"
)

func HealthCheck(api *api.Api) gin.HandlerFunc {
	return func(gc *gin.Context) {
		_, err := api.DbHandler.GetSecrets(gc.Request.Context())
		if err != nil {
			gc.JSON(http.StatusInternalServerError, nil)
		}
		gc.JSON(http.StatusOK, nil)
	}
}
