package standard

import (
	"github.com/SENERGY-Platform/mgw-secret-manager/internal/api"
	"github.com/SENERGY-Platform/mgw-secret-manager/internal/http_handler/util"
	"github.com/SENERGY-Platform/mgw-secret-manager/pkg/api_model"
	"github.com/gin-gonic/gin"
	"net/http"
)

// GetValueVariant godoc
// @Summary Get secret value
// @Description Get the value of a secret.
// @Tags Secrets
// @Accept json
// @Produce	json
// @Param request body api_model.SecretVariantRequest true "secret request"
// @Success	200 {object} api_model.SecretValueVariant "secret with value"
// @Failure	404 {string} string "error message"
// @Failure	500 {string} string "error message"
// @Router /value-variant [post]
func GetValueVariant(api *api.Api) (string, string, gin.HandlerFunc) {
	return http.MethodPost, api_model.ValueVariantPath, func(gc *gin.Context) {

		ok := util.CheckIfEncryptionKeyExists(gc, api)
		if !ok {
			return
		}

		secretVariantRequest, err := util.ParseVariantRequest(gc)
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
