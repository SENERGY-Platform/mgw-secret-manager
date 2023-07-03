package customErrors

import "fmt"

type MissingEncryptionKey struct{}

func (error MissingEncryptionKey) Error() string {
	return fmt.Sprintf("Encryption Key is missing -> POST at /key")
}

type NoSecretFound struct {
	SecretID string
}

func (error NoSecretFound) Error() string {
	return fmt.Sprintf("Secret with ID %s not found", error.SecretID)
}
