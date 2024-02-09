package api

import (
	"github.com/SENERGY-Platform/mgw-secret-manager/internal/config"
	"github.com/SENERGY-Platform/mgw-secret-manager/internal/db"
	"github.com/SENERGY-Platform/mgw-secret-manager/internal/keyHandler"
	"github.com/SENERGY-Platform/mgw-secret-manager/internal/secretHandler"
	"github.com/SENERGY-Platform/go-service-base/srv-info-hdl"
)

type Api struct {
	Config        config.Config
	DbHandler     db.Database                  // sql.db is threadsafe
	SecretHandler *secretHandler.SecretHandler // secretHandler contains mutex
	KeyHandler    keyHandler.KeyHandler
	SrvInfoHandler srv_info_hdl.SrvInfoHandler
}

func New(config config.Config, dbHandler db.Database, secretHandler *secretHandler.SecretHandler, keyHandler keyHandler.KeyHandler, srvInfoHandler srv_info_hdl.SrvInfoHandler) *Api {
	return &Api{
		Config:        config,
		DbHandler:     dbHandler,
		SecretHandler: secretHandler,
		KeyHandler:    keyHandler,
		SrvInfoHandler:srvInfoHandler,
	}
}
