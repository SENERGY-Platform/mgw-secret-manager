package api

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/SENERGY-Platform/mgw-secret-manager/internal/logger"
	"github.com/SENERGY-Platform/mgw-secret-manager/pkg/api_model"
	"github.com/gin-gonic/gin"
)

func (a *Api) LoadSecretIntoTMPFS(gc *gin.Context) {
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
	logger.Logger.Printf("POST Body: %s\n", body)

	var secretPostRequest api_model.SecretPostRequest
	err = json.Unmarshal(body, &secretPostRequest)
	if err != nil {
		logger.Logger.Errorf(err.Error())
		gc.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	err = a.secretHandler.LoadSecretToFileSystem(gc.Request.Context(), secretPostRequest)
	if err != nil {
		gc.AbortWithError(http.StatusInternalServerError, err)
		return
	}
}

func (a *Api) DeleteSecretFromTMPFS(gc *gin.Context) {
	body, err := ioutil.ReadAll(gc.Request.Body)
	if err != nil {
		logger.Logger.Errorf(err.Error())
		gc.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	var secretPostRequest api_model.SecretPostRequest
	err = json.Unmarshal(body, &secretPostRequest)
	if err != nil {
		logger.Logger.Errorf(err.Error())
		gc.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	a.secretHandler.RemoveSecretFromFileSystem(gc.Request.Context(), secretPostRequest)
}
