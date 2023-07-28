package httpHandler

import (
	"net/http"

	"github.com/SENERGY-Platform/mgw-secret-manager/internal/api"
	"github.com/gin-gonic/gin"
)

func UpdateSecret(api *api.Api) gin.HandlerFunc {
	return func(gc *gin.Context) {

		ok := CheckIfEncryptionKeyExists(gc, api)
		if !ok {
			return
		}

		secretRequest, err := ParseSecretCreateRequest(gc)

		secretID := gc.Param("id")

		err = api.SecretHandler.UpdateSecret(gc.Request.Context(), secretRequest, secretID)

		if err != nil {
			gc.AbortWithError(http.StatusInternalServerError, err)
			return
		}
		gc.JSON(http.StatusOK, nil)
	}
}
