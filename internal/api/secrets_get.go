package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

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

func (a *Api) GetShortSecret(gc *gin.Context) {
	ok := a.CheckIfEncryptionKeyExists(gc)
	if !ok {
		return
	}

	secretID := gc.Param("id")

	secret, err := a.secretHandler.GetShortSecret(gc.Request.Context(), secretID)
	if err != nil {
		gc.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	gc.JSON(http.StatusOK, secret)
}
