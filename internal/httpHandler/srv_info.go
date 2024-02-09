package httpHandler

import (
	"github.com/gin-gonic/gin"
	"github.com/SENERGY-Platform/mgw-secret-manager/internal/api"
	"net/http"
)

func GetSrvInfoH(api *api.Api) gin.HandlerFunc {
	return func(gc *gin.Context) {
		gc.JSON(http.StatusOK, api.SrvInfoHandler.GetInfo())
	}
}