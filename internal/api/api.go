package api

import (
	"secret-manager/internal/config"
	"secret-manager/internal/db"

	"github.com/gin-gonic/gin"
)

type Api struct {
	config    config.Config
	dbHandler db.Database
	masterKey *[]byte
}

func New(config config.Config, dbHandler db.Database) *Api {
	return &Api{
		config:    config,
		dbHandler: dbHandler,
	}
}
func (a *Api) SetRoutes(e *gin.Engine) {
	e.POST("/secret", a.StoreSecret)
	e.POST("/load", a.LoadSecretIntoTMPFS)
	e.POST("/key", a.SetEncryptionKey)
}
