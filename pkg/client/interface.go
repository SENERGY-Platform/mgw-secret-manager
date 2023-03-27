package client

type Client interface {
	StoreSecret(name string, value string, secretType string) (err error, errCode int)
	LoadSecretToTMPFS(secretName string) (fullTMPFSPath string, err error, errCode int)
	SetEncryptionKey(encryptionKey []byte) (err error, errCode int)
}
