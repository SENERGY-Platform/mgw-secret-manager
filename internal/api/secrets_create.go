package api

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/SENERGY-Platform/mgw-secret-manager/internal/logger"
	"github.com/SENERGY-Platform/mgw-secret-manager/pkg/api_model"
	"github.com/gin-gonic/gin"
)

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
