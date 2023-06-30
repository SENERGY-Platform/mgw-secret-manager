package models

type EncryptedSecret struct {
	Name       string
	Value      []byte
	SecretType string
	ID         string `gorm:"primaryKey"`
}

type Secret struct {
	Name       string `json:"name"`
	Value      string `json:"value"`
	SecretType string `json:"type"`
	ID         string `json:"id"`
}
