package http_handler

import (
	"github.com/SENERGY-Platform/mgw-secret-manager/internal/api"
	"github.com/SENERGY-Platform/mgw-secret-manager/pkg/api_model"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetSecrets(api *api.Api) (string, string, gin.HandlerFunc) {
	return http.MethodGet, api_model.SecretsPath, func(gc *gin.Context) {
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

func GetShortSecret(api *api.Api) (string, string, gin.HandlerFunc) {
	return http.MethodGet, api_model.SecretPath, func(gc *gin.Context) {
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
