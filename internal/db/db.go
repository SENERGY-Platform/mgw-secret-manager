package db

import (
	"fmt"

	srv_base "github.com/SENERGY-Platform/go-service-base/srv-base"
	"github.com/SENERGY-Platform/mgw-secret-manager/internal/config"
	"github.com/SENERGY-Platform/mgw-secret-manager/internal/model"
	"gorm.io/driver/mysql"
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

func (handler *DBHandler) GetSecrets() (secrets []*model.EncryptedSecret, err error) {
	handler.db.Find(&secrets)
	return
}

func (handler *DBHandler) Connect() (err error) {
	connectionUrl := fmt.Sprintf("%s", handler.config.DBConnectionURL)
	srv_base.Logger.Debugf("Connect to DB: %s", connectionUrl)
	handler.db, err = gorm.Open(mysql.Open(connectionUrl), &gorm.Config{})
	return
}

func NewDBHandler(config *config.Config) (handler *DBHandler, err error) {
	handler = &DBHandler{
		config: *config,
	}

	handler.Connect()
	handler.db.AutoMigrate(&model.EncryptedSecret{})

	return
}

func (handler *DBHandler) Cleanup() {
	handler.db.Session(&gorm.Session{AllowGlobalUpdate: true}).Delete(&model.EncryptedSecret{})
}
