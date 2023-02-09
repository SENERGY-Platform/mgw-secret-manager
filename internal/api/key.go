package api

import (
	"io/ioutil"
	"net/http"
	"os"
	"secret-manager/internal/core"

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

	var masterKey []byte
	if _, err := os.Stat(a.config.MasterKeyPath); err == nil {
		srv_base.Logger.Debug(("Master Encryption Key found -> Decrypt and Load"))
		masterKey, err = core.GetMasterKey(a.config, encryptionKey)
		if err != nil {
			srv_base.Logger.Error(err)
		}
	} else {
		srv_base.Logger.Debug(("Master Encryption Key not found -> Create, Encrypt and Store"))
		masterKey, err = core.CreateAndStoreMasterKey(a.config, encryptionKey)
		if err != nil {
			srv_base.Logger.Error(err)
		}
	}

	a.masterKey = &masterKey

}
