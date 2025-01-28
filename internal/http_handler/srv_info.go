package http_handler

import (
	_ "github.com/SENERGY-Platform/go-service-base/srv-info-hdl/lib"
	"github.com/SENERGY-Platform/mgw-secret-manager/internal/api"
	"github.com/gin-gonic/gin"
	"net/http"
)

// GetSrvInfoH godoc
// @Summary Get service info
// @Description	Get basic service and runtime information.
// @Tags Info
// @Produce	json
// @Success	200 {object} lib.SrvInfo "info"
// @Failure	500 {string} string "error message"
// @Router /info [get]
func GetSrvInfoH(api *api.Api) (string, string, gin.HandlerFunc) {
	return http.MethodGet, "/info", func(gc *gin.Context) {
		gc.JSON(http.StatusOK, api.SrvInfoHandler.GetInfo())
	}
}
