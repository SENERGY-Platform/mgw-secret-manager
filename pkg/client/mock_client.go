package client

type MockClient struct {
}

func (c *MockClient) StoreSecret(name string, value string, secretType string) (err error, errCode int) {
	return err, 0
}

func (c *MockClient) LoadSecretToTMPFS(secretName string) (fullTMPFSPath string, err error, errCode int) {
	return "path", nil, 200
}

func (c *MockClient) SetEncryptionKey(encryptionKey []byte) (err error, errCode int) {
	return nil, 0
}

func NewMockClient() (client Client, err error) {
	client = &MockClient{}
	return
}
