package http_handler

import (
	"github.com/SENERGY-Platform/mgw-secret-manager/internal/api"
	"github.com/SENERGY-Platform/mgw-secret-manager/pkg/api_model"
	"github.com/gin-gonic/gin"
	"net/http"
)

func UpdateSecret(api *api.Api) (string, string, gin.HandlerFunc) {
	return http.MethodPut, api_model.SecretPath, func(gc *gin.Context) {

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
