package api

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/SENERGY-Platform/mgw-secret-manager/internal/logger"
	"github.com/SENERGY-Platform/mgw-secret-manager/pkg/api_model"
	"github.com/gin-gonic/gin"
)

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
