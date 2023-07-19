package models

type EncryptedSecret struct {
	Name       string
	Value      []byte
	SecretType string
	ID         string `gorm:"primaryKey"`
}

type Secret struct {
	Name       string
	Value      string
	SecretType string
	ID         string
}

type SecretValue map[string]string
