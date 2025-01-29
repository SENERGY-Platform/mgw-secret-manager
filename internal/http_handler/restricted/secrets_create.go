package restricted

import (
	"github.com/SENERGY-Platform/mgw-secret-manager/internal/api"
	"github.com/SENERGY-Platform/mgw-secret-manager/internal/http_handler/util"
	"github.com/SENERGY-Platform/mgw-secret-manager/pkg/api_model"
	"github.com/gin-gonic/gin"
	"net/http"
)

// StoreSecret godoc
// @Summary Create secret
// @Description Store a secret.
// @Tags Secrets
// @Accept json
// @Produce	plain
// @Param data body api_model.SecretCreateRequest true "secret data"
// @Success	200 {string} string "secret ID"
// @Failure	500 {string} string "error message"
// @Router /secrets [post]
func StoreSecret(api *api.Api) (string, string, gin.HandlerFunc) {
	return http.MethodPost, api_model.SecretsPath, func(gc *gin.Context) {
		ok := util.CheckIfEncryptionKeyExists(gc, api)
		if !ok {
			return
		}

		secretRequest, err := util.ParseSecretCreateRequest(gc)

		secret := api.SecretHandler.CreateSecret(secretRequest.Name, secretRequest.Value, secretRequest.SecretType)

		err = api.SecretHandler.StoreSecret(gc.Request.Context(), &secret)
		if err != nil {
			gc.AbortWithError(http.StatusInternalServerError, err)
			return
		}
		gc.String(http.StatusOK, secret.ID)
	}
}
