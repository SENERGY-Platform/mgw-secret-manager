package http_handler

import (
	"github.com/SENERGY-Platform/mgw-secret-manager/internal/api"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetSrvInfoH(api *api.Api) gin.HandlerFunc {
	return func(gc *gin.Context) {
		gc.JSON(http.StatusOK, api.SrvInfoHandler.GetInfo())
	}
}
