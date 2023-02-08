package db

import (
	"fmt"
	"secret-manager/internal/config"
	"secret-manager/internal/model"

	srv_base "github.com/SENERGY-Platform/go-service-base/srv-base"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type DBHandler struct {
	db     *gorm.DB
	config config.Config
}

func (handler *DBHandler) SetSecret(secret *model.EncryptedSecret) (err error) {
	handler.db.Create(&secret)
	return
}

func (handler *DBHandler) GetSecret(secretName string) (secret *model.EncryptedSecret, err error) {
	handler.db.Where("name = ?", secretName).First(&secret)
	return
}

func (handler *DBHandler) connect() (err error) {
	dbFilePath := handler.config.DBFilePath
	connectionUrl := fmt.Sprintf("%s", dbFilePath)
	srv_base.Logger.Debugf("Connect to DB: %s", connectionUrl)
	handler.db, err = gorm.Open(sqlite.Open(connectionUrl), &gorm.Config{})
	return
}

func NewDBHandler(config config.Config) (handler *DBHandler, err error) {
	handler = &DBHandler{
		config: config,
	}

	handler.connect()
	handler.db.AutoMigrate(&model.EncryptedSecret{})

	return
}

func (handler *DBHandler) Cleanup() {
	handler.db.Session(&gorm.Session{AllowGlobalUpdate: true}).Delete(&model.Secret{})
	handler.db.Session(&gorm.Session{AllowGlobalUpdate: true}).Delete(&model.EncryptedSecret{})
}
