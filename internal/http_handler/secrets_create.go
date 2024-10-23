package http_handler

import (
	"github.com/SENERGY-Platform/mgw-secret-manager/internal/api"
	"github.com/gin-gonic/gin"
	"net/http"
)

func StoreSecret(api *api.Api) gin.HandlerFunc {
	return func(gc *gin.Context) {
		ok := CheckIfEncryptionKeyExists(gc, api)
		if !ok {
			return
		}

		secretRequest, err := ParseSecretCreateRequest(gc)

		secret := api.SecretHandler.CreateSecret(secretRequest.Name, secretRequest.Value, secretRequest.SecretType)

		err = api.SecretHandler.StoreSecret(gc.Request.Context(), &secret)
		if err != nil {
			gc.AbortWithError(http.StatusInternalServerError, err)
			return
		}
		gc.String(http.StatusOK, secret.ID)
	}
}
