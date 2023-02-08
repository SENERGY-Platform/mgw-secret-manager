package model

type EncryptedSecret struct {
	Name  string
	Value []byte
	ID    string `gorm:"primaryKey"`
}

type Secret struct {
	Name  string
	Value string
	ID    string
}
