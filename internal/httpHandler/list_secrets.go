package httpHandler

import (
	"net/http"

	"github.com/SENERGY-Platform/mgw-secret-manager/internal/api"
	"github.com/gin-gonic/gin"
)

func GetSecrets(api *api.Api) gin.HandlerFunc {
	return func(gc *gin.Context) {
		ok := CheckIfEncryptionKeyExists(gc, api)
		if !ok {
			return
		}

		secrets, err := api.SecretHandler.GetSecrets(gc.Request.Context())
		if err != nil {
			gc.AbortWithError(http.StatusInternalServerError, err)
			return
		}

		gc.JSON(http.StatusOK, secrets)
	}
}

func GetShortSecret(api *api.Api) gin.HandlerFunc {
	return func(gc *gin.Context) {
		ok := CheckIfEncryptionKeyExists(gc, api)
		if !ok {
			return
		}

		secretID := gc.Param("id")

		secret, err := api.SecretHandler.GetShortSecret(gc.Request.Context(), secretID)
		if err != nil {
			gc.AbortWithError(http.StatusInternalServerError, err)
			return
		}

		gc.JSON(http.StatusOK, secret)
	}
}
