package api

import (
	"io/ioutil"
	"net/http"

	"github.com/SENERGY-Platform/mgw-secret-manager/internal/logger"
	"github.com/gin-gonic/gin"
)

func (a *Api) SetEncryptionKey(gc *gin.Context) {
	if !a.config.EnableEncryption {
		logger.Logger.Errorf("Key was posted but encryption is disabled")
		gc.AbortWithError(http.StatusInternalServerError, gin.Error{})
		return
	}

	body, err := ioutil.ReadAll(gc.Request.Body)
	if err != nil {
		logger.Logger.Errorf("Error reading the Key from request: %s", err.Error())
		gc.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	encryptionKey := body

	err = a.keyHandler.SetEncryptionKey(gc.Request.Context(), encryptionKey, a.secretHandler)
	if err != nil {
		logger.Logger.Errorf("Error setting encryption key: %s", err.Error())
		gc.AbortWithError(http.StatusInternalServerError, err)
		return
	}
}
