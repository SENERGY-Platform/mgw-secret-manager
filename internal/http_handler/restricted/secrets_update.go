package restricted

import (
	"github.com/SENERGY-Platform/mgw-secret-manager/internal/api"
	"github.com/SENERGY-Platform/mgw-secret-manager/internal/http_handler/util"
	"github.com/SENERGY-Platform/mgw-secret-manager/pkg/api_model"
	"github.com/gin-gonic/gin"
	"net/http"
)

// UpdateSecret godoc
// @Summary Update secret
// @Description Update a secret.
// @Tags Secrets
// @Accept json
// @Param id path string true "secret ID"
// @Param data body api_model.SecretCreateRequest true "secret data"
// @Success	200
// @Failure	404 {string} string "error message"
// @Failure	500 {string} string "error message"
// @Router /secrets/{id} [put]
func UpdateSecret(api *api.Api) (string, string, gin.HandlerFunc) {
	return http.MethodPut, api_model.SecretPath, func(gc *gin.Context) {

		ok := util.CheckIfEncryptionKeyExists(gc, api)
		if !ok {
			return
		}

		secretRequest, err := util.ParseSecretCreateRequest(gc)

		secretID := gc.Param("id")

		err = api.SecretHandler.UpdateSecret(gc.Request.Context(), secretRequest, secretID)

		if err != nil {
			gc.AbortWithError(http.StatusInternalServerError, err)
			return
		}
		gc.JSON(http.StatusOK, nil)
	}
}
