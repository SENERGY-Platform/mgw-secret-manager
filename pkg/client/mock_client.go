package client

import (
	"github.com/SENERGY-Platform/mgw-secret-manager/internal/config"
	"github.com/SENERGY-Platform/mgw-secret-manager/internal/core"
	"github.com/SENERGY-Platform/mgw-secret-manager/internal/db"
)

type MockClient struct {
	dbHandler db.Database
	masterKey *[]byte
	config    config.Config
}

func (c *MockClient) StoreSecret(name string, value string, secretType string) (err error, errCode int) {
	secret := core.CreateSecret(name, value, secretType)

	err = core.StoreSecret(&secret, &c.dbHandler, c.masterKey, c.config)
	if err != nil {
		return err, 0
	}
	return err, 0
}

func (c *MockClient) LoadSecretToTMPFS(secretName string) (fullTMPFSPath string, err error, errCode int) {
	fullTMPFSPath, err = core.LoadSecretToFileSystem(secretName, &c.dbHandler, c.config, c.masterKey)
	return
}

func (c *MockClient) SetEncryptionKey(encryptionKey []byte) (err error, errCode int) {
	return nil, 0
}

func NewMockClient() (client Client, err error) {
	masterKey, err := core.GenerateMasterKey()
	testConfig, err := config.NewConfig(nil)
	if err != nil {
		return nil, err
	}

	client = &MockClient{
		dbHandler: db.NewMockDB(),
		masterKey: &masterKey,
		config:    *testConfig,
	}
	return
}
