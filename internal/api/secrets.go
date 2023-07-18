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

	var secretRequest api_model.SecretCreateRequest
	err = json.Unmarshal(body, &secretRequest)
	if err != nil {
		logger.Logger.Errorf(err.Error())
		gc.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	secret := a.secretHandler.CreateSecret(secretRequest.Name, secretRequest.Value, secretRequest.SecretType)

	err = a.secretHandler.StoreSecret(gc.Request.Context(), &secret)
	if err != nil {
		gc.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	gc.String(http.StatusOK, secret.ID)
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

	var secretRequest api_model.SecretCreateRequest
	err = json.Unmarshal(body, &secretRequest)
	if err != nil {
		logger.Logger.Errorf(err.Error())
		gc.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	secretID := gc.Param("id")

	err = a.secretHandler.UpdateSecret(gc.Request.Context(), secretRequest, secretID)

	if err != nil {
		gc.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	gc.JSON(http.StatusOK, nil)
}

func (a *Api) GetSecret(gc *gin.Context) {
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

	var secretPostRequest api_model.SecretVariantRequest
	err = json.Unmarshal(body, &secretPostRequest)
	if err != nil {
		logger.Logger.Errorf(err.Error())
		gc.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	secret, err := a.secretHandler.GetSecret(gc.Request.Context(), secretPostRequest)
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

	secrets, err := a.secretHandler.GetSecrets(gc.Request.Context())
	if err != nil {
		gc.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	gc.JSON(http.StatusOK, secrets)
}

func (a *Api) GetFullSecret(gc *gin.Context) {
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

	var secretPostRequest api_model.SecretVariantRequest
	err = json.Unmarshal(body, &secretPostRequest)
	if err != nil {
		logger.Logger.Errorf(err.Error())
		gc.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	secrets, err := a.secretHandler.GetFullSecret(gc.Request.Context(), secretPostRequest)
	if err != nil {
		gc.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	gc.JSON(http.StatusOK, secrets)
}

func (a *Api) DeleteSecret(gc *gin.Context) {
	secretID := gc.Param("id")

	err := a.secretHandler.DeleteSecret(gc.Request.Context(), secretID)
	if err != nil {
		gc.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	gc.JSON(http.StatusOK, nil)
}

func (a *Api) GetTypes(gc *gin.Context) {
	gc.JSON(http.StatusOK, []map[string]string{{"id": "certificate", "name": "Certificate"}, {"id": "basic-auth", "name": "Credentials"}, {"id": "api-key", "name": "API Key"}})
}
