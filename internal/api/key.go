package api

import (
	"io/ioutil"
	"net/http"

	srv_base "github.com/SENERGY-Platform/go-service-base/srv-base"
	"github.com/gin-gonic/gin"
)

func (a *Api) SetEncryptionKey(gc *gin.Context) {
	if !a.config.EnableEncryption {
		srv_base.Logger.Errorf("Key was posted but encryption is disabled")
		gc.AbortWithError(http.StatusInternalServerError, gin.Error{})
		return
	}

	body, err := ioutil.ReadAll(gc.Request.Body)
	if err != nil {
		srv_base.Logger.Errorf("Error reading the Key from request: %s", err.Error())
		gc.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	encryptionKey := body

	err = a.keyHandler.SetEncryptionKey(encryptionKey, &a.secretHandler)
	if err != nil {
		srv_base.Logger.Errorf("Error setting encryption key: %s", err.Error())
		gc.AbortWithError(http.StatusInternalServerError, err)
		return
	}
}
