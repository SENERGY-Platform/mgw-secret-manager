package db

import (
	"errors"
	"fmt"
	"time"

	"github.com/SENERGY-Platform/mgw-secret-manager/internal/config"
	"github.com/SENERGY-Platform/mgw-secret-manager/internal/customErrors"
	"github.com/SENERGY-Platform/mgw-secret-manager/internal/logger"
	"github.com/SENERGY-Platform/mgw-secret-manager/internal/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type DBHandler struct {
	db     *gorm.DB
	config config.Config
}

func (handler *DBHandler) SetSecret(secret *models.EncryptedSecret) (err error) {
	err = handler.db.Create(&secret).Error
	return
}

func (handler *DBHandler) GetSecret(secretID string) (secret *models.EncryptedSecret, err error) {
	err = handler.db.Where("ID = ?", secretID).First(&secret).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			err = customErrors.NoSecretFound{SecretID: secretID}
		}
	}

	return
}

func (handler *DBHandler) GetSecrets() (secrets []*models.EncryptedSecret, err error) {
	err = handler.db.Find(&secrets).Error
	return
}

func (handler *DBHandler) UpdateSecret(secret *models.EncryptedSecret) (err error) {
	err = handler.db.Save(secret).Error
	return
}

func (handler *DBHandler) DeleteSecret(secretID string) (err error) {
	secret := models.EncryptedSecret{
		ID: secretID,
	}
	err = handler.db.Delete(&secret).Error

	return
}

func (handler *DBHandler) Connect() (err error) {
	connectionUrl := fmt.Sprintf("%s", handler.config.DBConnectionURL)
	logger.Logger.Debugf("Connect to DB: %s", connectionUrl)
	handler.db, err = gorm.Open(mysql.Open(connectionUrl), &gorm.Config{})
	for err != nil {
		retries := 5
		if retries > 1 {
			retries--
			time.Sleep(5 * time.Second)
			handler.db, err = gorm.Open(mysql.Open(connectionUrl), &gorm.Config{})
		}
	}

	return
}

func NewDBHandler(config *config.Config) (handler *DBHandler, err error) {
	handler = &DBHandler{
		config: *config,
	}

	handler.Connect()
	handler.db.AutoMigrate(&models.EncryptedSecret{})

	return
}

func (handler *DBHandler) Cleanup() {
	handler.db.Session(&gorm.Session{AllowGlobalUpdate: true}).Delete(&models.EncryptedSecret{})
}
