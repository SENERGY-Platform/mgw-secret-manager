package api

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/SENERGY-Platform/mgw-secret-manager/internal/secretHandler"
	"github.com/SENERGY-Platform/mgw-secret-manager/pkg/api_model"

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

	var secretRequest api_model.SecretRequest
	err = json.Unmarshal(body, &secretRequest)
	if err != nil {
		srv_base.Logger.Errorf(err.Error())
		gc.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	secret := secretHandler.CreateSecret(secretRequest.Name, secretRequest.Value, secretRequest.SecretType)

	err = secretHandler.StoreSecret(&secret, a.dbHandler, a.masterKey, a.config.EnableEncryption)
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

		fullTMPFSPath, err := secretHandler.LoadSecretToFileSystem(secretName, a.dbHandler, a.config, a.masterKey)
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

	secrets, err := secretHandler.GetSecrets(a.dbHandler, a.config)
	if err != nil {
		gc.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	gc.JSON(http.StatusOK, secrets)
}

func (a *Api) GetTypes(gc *gin.Context) {
	gc.JSON(http.StatusOK, []map[string]string{{"id": "certificate", "name": "Certificate"}, {"id": "basic-auth", "name": "Credentials"}, {"id": "api-key", "name": "API Key"}})
}
