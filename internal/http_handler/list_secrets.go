package http_handler

import (
	"github.com/SENERGY-Platform/mgw-secret-manager/internal/api"
	"github.com/SENERGY-Platform/mgw-secret-manager/pkg/api_model"
	"github.com/gin-gonic/gin"
	"net/http"
)

// GetSecrets godoc
// @Summary Get Secrets
// @Description List stored secrets.
// @Tags Secrets
// @Produce	json
// @Success	200 {array} api_model.Secret "secrets"
// @Failure	500 {string} string "error message"
// @Router /secrets [get]
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

// GetShortSecret godoc
// @Summary Get secret
// @Description Get a secret.
// @Tags Secrets
// @Produce	json
// @Param id path string true "secret ID"
// @Success	200 {object} api_model.Secret "secret"
// @Failure	404 {string} string "error message"
// @Failure	500 {string} string "error message"
// @Router /secrets/{id} [get]
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
