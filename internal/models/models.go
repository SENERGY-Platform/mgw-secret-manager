package models

type EncryptedSecret struct {
	Name       string
	Value      []byte
	SecretType string
	ID         string `gorm:"primaryKey"`
}

type SecretValue map[string]string
