package db

import (
	"context"
	"errors"
	"fmt"
	"github.com/SENERGY-Platform/mgw-secret-manager/internal/custom_errors"
	"github.com/SENERGY-Platform/mgw-secret-manager/internal/models"
	"github.com/SENERGY-Platform/mgw-secret-manager/internal/util"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"time"
)

type DBHandler struct {
	db     *gorm.DB
	config util.Config
}

func (handler *DBHandler) SetSecret(ctx context.Context, secret *models.EncryptedSecret) (err error) {
	err = handler.db.WithContext(ctx).Create(&secret).Error
	return
}

func (handler *DBHandler) GetSecret(ctx context.Context, secretID string) (secret *models.EncryptedSecret, err error) {
	err = handler.db.WithContext(ctx).Where("ID = ?", secretID).First(&secret).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			err = custom_errors.NoSecretFound{SecretID: secretID}
		}
	}

	return
}

func (handler *DBHandler) GetSecrets(ctx context.Context) (secrets []*models.EncryptedSecret, err error) {
	err = handler.db.WithContext(ctx).Find(&secrets).Error
	return
}

func (handler *DBHandler) UpdateSecret(ctx context.Context, secret *models.EncryptedSecret) (err error) {
	err = handler.db.WithContext(ctx).Updates(secret).Error
	return
}

func (handler *DBHandler) DeleteSecret(ctx context.Context, secretID string) (err error) {
	secret := models.EncryptedSecret{
		ID: secretID,
	}
	err = handler.db.WithContext(ctx).Delete(&secret).Error

	return
}

func (handler *DBHandler) Connect() (err error) {
	connectionUrl := fmt.Sprintf("%s", handler.config.DBConnectionURL)
	util.Logger.Debugf("Inital connect to DB: %s", connectionUrl)
	handler.db, err = gorm.Open(mysql.Open(connectionUrl), &gorm.Config{Logger: logger.Discard})
	for err != nil {
		util.Logger.Debugf("DB is not reachable -> try again in 5s")
		retries := 5
		if retries > 1 {
			retries--
			time.Sleep(5 * time.Second)
			handler.db, err = gorm.Open(mysql.Open(connectionUrl), &gorm.Config{
				Logger: logger.Discard,
			})
		}
	}

	return
}

func NewDBHandler(config *util.Config) (handler *DBHandler, err error) {
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
