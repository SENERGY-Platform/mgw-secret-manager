package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (a *Api) DeleteSecret(gc *gin.Context) {
	secretID := gc.Param("id")

	err := a.secretHandler.DeleteSecret(gc.Request.Context(), secretID)
	if err != nil {
		gc.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	gc.JSON(http.StatusOK, nil)
}
