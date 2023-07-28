package httpHandler

import (
	"net/http"

	"github.com/SENERGY-Platform/mgw-secret-manager/internal/api"
	"github.com/gin-gonic/gin"
)

func GetValueVariant(api *api.Api) gin.HandlerFunc {
	return func(gc *gin.Context) {

		ok := CheckIfEncryptionKeyExists(gc, api)
		if !ok {
			return
		}

		secretVariantRequest, err := ParseVariantRequest(gc)
		if err != nil {
			gc.AbortWithError(http.StatusInternalServerError, err)
			return
		}

		secretValueVariant, err := api.SecretHandler.GetValueVariant(gc.Request.Context(), secretVariantRequest)
		if err != nil {
			gc.AbortWithError(http.StatusInternalServerError, err)
			return
		}

		gc.JSON(http.StatusOK, secretValueVariant)
	}
}
