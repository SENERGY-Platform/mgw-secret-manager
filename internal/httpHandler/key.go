package httpHandler

import (
	"io/ioutil"
	"net/http"

	"github.com/SENERGY-Platform/mgw-secret-manager/internal/api"
	"github.com/SENERGY-Platform/mgw-secret-manager/internal/customErrors"
	"github.com/SENERGY-Platform/mgw-secret-manager/internal/logger"
	"github.com/gin-gonic/gin"
)

func SetEncryptionKey(api *api.Api) gin.HandlerFunc {
	return func(gc *gin.Context) {
		if !api.Config.EnableEncryption {
			logger.Logger.Errorf("Key was posted but encryption is disabled")
			err := customErrors.EncryptionIsDisabled{}
			gc.AbortWithError(http.StatusInternalServerError, err)
			return
		}

		body, err := ioutil.ReadAll(gc.Request.Body)
		if err != nil {
			logger.Logger.Errorf("Error reading the Key from request: %s", err.Error())
			gc.AbortWithError(http.StatusInternalServerError, err)
			return
		}
		encryptionKey := body

		err = api.KeyHandler.SetEncryptionKey(gc.Request.Context(), encryptionKey, api.SecretHandler)
		if err != nil {
			logger.Logger.Errorf("Error setting encryption key: %s", err.Error())
			gc.AbortWithError(http.StatusInternalServerError, err)
			return
		}
	}
}
