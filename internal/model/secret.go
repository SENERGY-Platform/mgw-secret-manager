package model

type EncryptedSecret struct {
	Name  string
	Value []byte
}

type Secret struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}
