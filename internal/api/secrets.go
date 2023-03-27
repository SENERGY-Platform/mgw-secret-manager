package api

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/SENERGY-Platform/mgw-secret-manager/internal/core"

	srv_base "github.com/SENERGY-Platform/go-service-base/srv-base"
	"github.com/gin-gonic/gin"
)

func (a *Api) StoreSecret(gc *gin.Context) {
	if a.masterKey == nil && a.config.EnableEncryption == true {
		gc.AbortWithError(http.StatusInternalServerError, MissingEncryptionKey{})
	}

	body, err := ioutil.ReadAll(gc.Request.Body)
	if err != nil {
		gc.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	var secretRequest SecretRequest
	err = json.Unmarshal(body, &secretRequest)
	if err != nil {
		srv_base.Logger.Errorf(err.Error())
		gc.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	secret := core.CreateSecret(secretRequest.Name, secretRequest.Value, secretRequest.SecretType)

	err = core.StoreSecret(&secret, a.dbHandler, *a.masterKey, a.config)
	if err != nil {
		gc.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	gc.JSON(http.StatusOK, secret.ID)
}

func (a *Api) LoadSecretIntoTMPFS(gc *gin.Context) {
	if a.masterKey == nil && a.config.EnableEncryption == true {
		gc.AbortWithError(http.StatusInternalServerError, MissingEncryptionKey{})
	}

	if secretNames, ok := gc.Request.URL.Query()["secret"]; ok {
		secretName := secretNames[0]

		fullTMPFSPath, err := core.LoadSecretToFileSystem(secretName, a.dbHandler, a.config, *a.masterKey)
		if err != nil {
			gc.AbortWithError(http.StatusInternalServerError, err)
			return
		}
		gc.JSON(http.StatusOK, fullTMPFSPath)
	} else {
		gc.AbortWithError(http.StatusInternalServerError, MissingQueryError{Parameter: "secret"})
	}
}

func (a *Api) GetSecrets(gc *gin.Context) {
	if a.masterKey == nil && a.config.EnableEncryption == true {
		gc.AbortWithError(http.StatusInternalServerError, MissingEncryptionKey{})
		return
	}

	secrets, err := core.GetSecrets(a.dbHandler, a.config, *a.masterKey)
	if err != nil {
		gc.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	gc.JSON(http.StatusOK, secrets)
}
