package models

type EncryptedSecret struct {
	Name       string
	Value      []byte
	SecretType string
	FileName   string
	ID         string `gorm:"primaryKey"`
}

type Credential struct {
	Username string
	Password string
}
