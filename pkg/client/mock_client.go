package client

import (
	"secret-manager/internal/config"
	"secret-manager/internal/core"
	"secret-manager/internal/db"
	"secret-manager/test"
)

type MockClient struct {
	dbHandler db.Database
	masterKey *[]byte
	config    config.Config
}

func (c *MockClient) StoreSecret(name string, value string) (err error, errCode int) {
	secret := core.CreateSecret(name, value)

	err = core.StoreSecret(&secret, c.dbHandler, *c.masterKey)
	if err != nil {
		return err, 0
	}
	return err, 0
}

func (c *MockClient) LoadSecretToTMPFS(secretName string) (fullTMPFSPath string, err error, errCode int) {
	fullTMPFSPath, err = core.LoadSecretToFileSystem(secretName, c.dbHandler, c.config, *c.masterKey)
	return
}

func (c *MockClient) SetEncryptionKey(encryptionKey []byte) (err error, errCode int) {
	return nil, 0
}

func NewMockClient() (client Client, err error) {
	masterKey, err := core.GenerateMasterKey()
	client = &MockClient{
		dbHandler: db.NewMockDB(),
		masterKey: &masterKey,
		config:    test.TestConfig,
	}
	return
}
