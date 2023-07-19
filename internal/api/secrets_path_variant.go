package api

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/SENERGY-Platform/mgw-secret-manager/internal/logger"
	"github.com/SENERGY-Platform/mgw-secret-manager/pkg/api_model"
	"github.com/gin-gonic/gin"
)

func (a *Api) InitPathVariant(gc *gin.Context) {
	ok := a.CheckIfEncryptionKeyExists(gc)
	if !ok {
		return
	}

	secretVariantRequest, err := ParseVariantRequest(gc)
	if err != nil {
		gc.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	err = a.secretHandler.InitPathVariant(gc.Request.Context(), secretVariantRequest)
	if err != nil {
		gc.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	gc.String(http.StatusOK, "")
}

func (a *Api) LoadPathVariant(gc *gin.Context) {
	ok := a.CheckIfEncryptionKeyExists(gc)
	if !ok {
		return
	}

	secretVariantRequest, err := ParseVariantRequest(gc)
	if err != nil {
		gc.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	err = a.secretHandler.LoadSecretToFileSystem(gc.Request.Context(), secretVariantRequest)
	if err != nil {
		gc.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	gc.String(http.StatusOK, "")
}

func (a *Api) DeleteSecretFromTMPFS(gc *gin.Context) {
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

	a.secretHandler.RemoveSecretFromFileSystem(gc.Request.Context(), secretPostRequest)
}
