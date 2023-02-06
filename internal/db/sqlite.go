package db

import (
	"fmt"
	"secret-manager/internal/config"
	"secret-manager/internal/model"

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

	//handler.db, _ = sql.Open("sqlite3", connectionUrl)
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
