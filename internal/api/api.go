package api

import (
	"github.com/SENERGY-Platform/mgw-secret-manager/internal/config"
	"github.com/SENERGY-Platform/mgw-secret-manager/internal/db"

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
	e.POST("/secrets", a.StoreSecret)
	e.GET("/secrets", a.GetSecrets)
	e.GET("/types", a.GetTypes)
	e.POST("/load", a.LoadSecretIntoTMPFS)
	e.POST("/key", a.SetEncryptionKey)
}
