package api

import (
	"github.com/SENERGY-Platform/mgw-secret-manager/internal/config"
	"github.com/SENERGY-Platform/mgw-secret-manager/internal/db"
	"github.com/SENERGY-Platform/mgw-secret-manager/internal/keyHandler"
	"github.com/SENERGY-Platform/mgw-secret-manager/internal/secretHandler"
)

type Api struct {
	Config        config.Config
	DbHandler     db.Database                  // sql.db is threadsafe
	SecretHandler *secretHandler.SecretHandler // secretHandler contains mutex
	KeyHandler    keyHandler.KeyHandler
}

func New(config config.Config, dbHandler db.Database, secretHandler *secretHandler.SecretHandler, keyHandler keyHandler.KeyHandler) *Api {
	return &Api{
		Config:        config,
		DbHandler:     dbHandler,
		SecretHandler: secretHandler,
		KeyHandler:    keyHandler,
	}
}
