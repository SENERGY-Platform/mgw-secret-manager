package db

import (
	"fmt"
	"os"

	"github.com/SENERGY-Platform/mgw-secret-manager/internal/config"
	"github.com/SENERGY-Platform/mgw-secret-manager/internal/model"
)

type Database interface {
	SetSecret(secret *model.EncryptedSecret) (err error)
	GetSecret(secretName string) (secret *model.EncryptedSecret, err error)
	GetSecrets() (secrets []*model.EncryptedSecret, err error)
	Connect() (err error)
	Cleanup()
}

func GetTestDB(config *config.Config) (Database, error) {
	testMode := os.Getenv("TEST_MODE")
	fmt.Printf(config.DBConnectionURL)
	if testMode == "INTEGRATION" {
		return NewDBHandler(*config)
	} else {
		return NewMockDB(), nil
	}
}
