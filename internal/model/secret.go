package model

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

type ShortSecret struct {
	Name       string `json:"name"`
	SecretType string `json:"type"`
	ID         string `json:"id"`
}
