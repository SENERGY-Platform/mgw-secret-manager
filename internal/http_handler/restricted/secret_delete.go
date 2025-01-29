package restricted

import (
	"github.com/SENERGY-Platform/mgw-secret-manager/internal/api"
	"github.com/SENERGY-Platform/mgw-secret-manager/pkg/api_model"
	"github.com/gin-gonic/gin"
	"net/http"
)

// DeleteSecret godoc
// @Summary Delete secret
// @Description Remove a secret.
// @Tags Secrets
// @Param id path string true "secret ID"
// @Success	200
// @Failure	404 {string} string "error message"
// @Failure	500 {string} string "error message"
// @Router /secrets/{id} [delete]
func DeleteSecret(api *api.Api) (string, string, gin.HandlerFunc) {
	return http.MethodDelete, api_model.SecretPath, func(gc *gin.Context) {
		secretID := gc.Param("id")

		err := api.SecretHandler.DeleteSecret(gc.Request.Context(), secretID)
		if err != nil {
			gc.AbortWithError(http.StatusInternalServerError, err)
			return
		}

		gc.JSON(http.StatusOK, nil)
	}
}
