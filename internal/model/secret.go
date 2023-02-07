package model

type EncryptedSecret struct {
	Name  string
	Value []byte
	ID    uint64 `gorm:"primaryKey;auto_increment"`
}

type Secret struct {
	Name  string
	Value string
	ID    uint64
}
