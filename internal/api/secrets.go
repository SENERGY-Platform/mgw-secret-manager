package api

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/SENERGY-Platform/mgw-secret-manager/internal/customErrors"
	"github.com/SENERGY-Platform/mgw-secret-manager/internal/logger"
	"github.com/SENERGY-Platform/mgw-secret-manager/pkg/api_model"

	"github.com/gin-gonic/gin"
)

func (a *Api) CheckIfEncryptionKeyExists(gc *gin.Context) bool {
	if a.secretHandler.Key == nil && a.config.EnableEncryption == true {
		gc.AbortWithError(http.StatusInternalServerError, customErrors.MissingEncryptionKey{})
		gc.Abort()
		return false
	}
	return true
}

func (a *Api) StoreSecret(gc *gin.Context) {
	ok := a.CheckIfEncryptionKeyExists(gc)
	if !ok {
		return
	}

	body, err := ioutil.ReadAll(gc.Request.Body)
	if err != nil {
		logger.Logger.Errorf(err.Error())
		gc.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	var secretRequest api_model.SecretRequest
	err = json.Unmarshal(body, &secretRequest)
	if err != nil {
		logger.Logger.Errorf(err.Error())
		gc.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	secret := a.secretHandler.CreateSecret(secretRequest.Name, secretRequest.Value, secretRequest.SecretType)

	err = a.secretHandler.StoreSecret(&secret)
	if err != nil {
		gc.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	gc.JSON(http.StatusOK, secret.ID)
}

func (a *Api) UpdateSecret(gc *gin.Context) {
	ok := a.CheckIfEncryptionKeyExists(gc)
	if !ok {
		return
	}

	body, err := ioutil.ReadAll(gc.Request.Body)
	if err != nil {
		gc.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	var secretRequest api_model.SecretRequest
	err = json.Unmarshal(body, &secretRequest)
	if err != nil {
		logger.Logger.Errorf(err.Error())
		gc.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	secretID := gc.Param("id")

	err = a.secretHandler.UpdateSecret(secretRequest, secretID)

	if err != nil {
		gc.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	gc.JSON(http.StatusOK, nil)
}

func (a *Api) LoadSecretIntoTMPFS(gc *gin.Context) {
	ok := a.CheckIfEncryptionKeyExists(gc)
	if !ok {
		return
	}

	if secretIDs, ok := gc.Request.URL.Query()["secret"]; ok {
		secretID := secretIDs[0]

		fullTMPFSPath, err := a.secretHandler.LoadSecretToFileSystem(secretID)
		if err != nil {
			gc.AbortWithError(http.StatusInternalServerError, err)
			return
		}
		gc.String(http.StatusOK, fullTMPFSPath)
	} else {
		gc.AbortWithError(http.StatusInternalServerError, customErrors.MissingQueryError{Parameter: "secret"})
	}
}

func (a *Api) GetSecret(gc *gin.Context) {
	ok := a.CheckIfEncryptionKeyExists(gc)
	if !ok {
		return
	}

	secretID := gc.Param("id")
	secret, err := a.secretHandler.GetSecret(secretID)
	if err != nil {
		gc.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	gc.JSON(http.StatusOK, secret)
}

func (a *Api) GetSecrets(gc *gin.Context) {
	ok := a.CheckIfEncryptionKeyExists(gc)
	if !ok {
		return
	}

	secrets, err := a.secretHandler.GetSecrets()
	if err != nil {
		gc.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	gc.JSON(http.StatusOK, secrets)
}

func (a *Api) DeleteSecret(gc *gin.Context) {
	secretID := gc.Param("id")

	err := a.secretHandler.DeleteSecret(secretID)
	if err != nil {
		gc.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	gc.JSON(http.StatusOK, nil)
}

func (a *Api) GetTypes(gc *gin.Context) {
	gc.JSON(http.StatusOK, []map[string]string{{"id": "certificate", "name": "Certificate"}, {"id": "basic-auth", "name": "Credentials"}, {"id": "api-key", "name": "API Key"}})
}
