package http_handler

import (
	"github.com/SENERGY-Platform/mgw-secret-manager/internal/api"
	"github.com/gin-gonic/gin"
	"net/http"
)

func DeleteSecret(api *api.Api) gin.HandlerFunc {
	return func(gc *gin.Context) {
		secretID := gc.Param("id")

		err := api.SecretHandler.DeleteSecret(gc.Request.Context(), secretID)
		if err != nil {
			gc.AbortWithError(http.StatusInternalServerError, err)
			return
		}

		gc.JSON(http.StatusOK, nil)
	}
}
