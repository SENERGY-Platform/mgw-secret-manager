package api

import (
	"github.com/SENERGY-Platform/go-service-base/srv-info-hdl"
	"github.com/SENERGY-Platform/mgw-secret-manager/internal/db"
	"github.com/SENERGY-Platform/mgw-secret-manager/internal/key_handler"
	"github.com/SENERGY-Platform/mgw-secret-manager/internal/secret_handler"
	"github.com/SENERGY-Platform/mgw-secret-manager/internal/util"
)

type Api struct {
	Config         util.Config
	DbHandler      db.Database                   // sql.db is threadsafe
	SecretHandler  *secret_handler.SecretHandler // secretHandler contains mutex
	KeyHandler     key_handler.KeyHandler
	SrvInfoHandler srv_info_hdl.SrvInfoHandler
}

func New(config util.Config, dbHandler db.Database, secretHandler *secret_handler.SecretHandler, keyHandler key_handler.KeyHandler, srvInfoHandler srv_info_hdl.SrvInfoHandler) *Api {
	return &Api{
		Config:         config,
		DbHandler:      dbHandler,
		SecretHandler:  secretHandler,
		KeyHandler:     keyHandler,
		SrvInfoHandler: srvInfoHandler,
	}
}
