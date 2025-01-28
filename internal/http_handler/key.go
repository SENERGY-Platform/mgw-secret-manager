package http_handler

import (
	"github.com/SENERGY-Platform/mgw-secret-manager/internal/api"
	"github.com/SENERGY-Platform/mgw-secret-manager/internal/custom_errors"
	"github.com/SENERGY-Platform/mgw-secret-manager/internal/util"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
)

func SetEncryptionKey(api *api.Api) (string, string, gin.HandlerFunc) {
	return http.MethodPost, "/key", func(gc *gin.Context) {
		if !api.Config.EnableEncryption {
			util.Logger.Errorf("Key was posted but encryption is disabled")
			err := custom_errors.EncryptionIsDisabled{}
			gc.AbortWithError(http.StatusInternalServerError, err)
			return
		}

		body, err := io.ReadAll(gc.Request.Body)
		if err != nil {
			util.Logger.Errorf("Error reading the Key from request: %s", err.Error())
			gc.AbortWithError(http.StatusInternalServerError, err)
			return
		}
		encryptionKey := body

		err = api.KeyHandler.SetEncryptionKey(gc.Request.Context(), encryptionKey, api.SecretHandler)
		if err != nil {
			util.Logger.Errorf("Error setting encryption key: %s", err.Error())
			gc.AbortWithError(http.StatusInternalServerError, err)
			return
		}
	}
}
