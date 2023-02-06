package api

import (
	"secret-manager/internal/config"
	"secret-manager/internal/db"

	"github.com/gin-gonic/gin"
)

type Api struct {
	config    config.Config
	dbHandler *db.DBHandler
	masterKey []byte
}

func New(config config.Config, dbHandler *db.DBHandler, masterKey []byte) *Api {
	return &Api{
		config:    config,
		dbHandler: dbHandler,
		masterKey: masterKey,
	}
}
func (a *Api) SetRoutes(e *gin.Engine) {
	e.POST("/secret", a.StoreSecret)
	e.POST("/load", a.LoadSecretIntoTMPFS)
}
