package api

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"secret-manager/internal/core"
	"secret-manager/internal/model"

	srv_base "github.com/SENERGY-Platform/go-service-base/srv-base"
	"github.com/gin-gonic/gin"
)

func (a *Api) StoreSecret(gc *gin.Context) {
	body, err := ioutil.ReadAll(gc.Request.Body)
	if err != nil {
		gc.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	var secret model.Secret
	err = json.Unmarshal(body, &secret)
	if err != nil {
		srv_base.Logger.Errorf(err.Error())
		gc.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	srv_base.Logger.Debugf(secret.Value)

	err = core.StoreSecret(&secret, a.dbHandler, a.masterKey)
	if err != nil {
		gc.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	gc.JSON(http.StatusOK, "Secret was stored successfully")
}

func (a *Api) LoadSecretIntoTMPFS(gc *gin.Context) {
	if secretNames, ok := gc.Request.URL.Query()["secret"]; ok {
		secretName := secretNames[0]

		err := core.LoadSecretToFileSystem(secretName, a.dbHandler, a.config, a.masterKey)
		if err != nil {
			gc.AbortWithError(http.StatusInternalServerError, err)
			return
		}
	} else {
		gc.AbortWithError(http.StatusInternalServerError, MissingQueryError{Parameter: "secret"})
	}
}
