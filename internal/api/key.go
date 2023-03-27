package api

import (
	"io/ioutil"
	"net/http"

	"github.com/SENERGY-Platform/mgw-secret-manager/internal/core"

	srv_base "github.com/SENERGY-Platform/go-service-base/srv-base"
	"github.com/gin-gonic/gin"
)

func (a *Api) SetEncryptionKey(gc *gin.Context) {
	if a.config.EnableEncryption != true {
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

	masterKey, err := core.SetEncryptionKey(encryptionKey, a.config)
	if err != nil {
		srv_base.Logger.Errorf("Error setting encryption key: %s", err.Error())
		gc.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	a.masterKey = &masterKey

}
