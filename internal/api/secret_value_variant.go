package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (a *Api) GetValueVariant(gc *gin.Context) {
	ok := a.CheckIfEncryptionKeyExists(gc)
	if !ok {
		return
	}

	secretVariantRequest, err := ParseVariantRequest(gc)
	if err != nil {
		gc.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	secretValueVariant, err := a.secretHandler.GetValueVariant(gc.Request.Context(), secretVariantRequest)
	if err != nil {
		gc.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	gc.JSON(http.StatusOK, secretValueVariant)

}
