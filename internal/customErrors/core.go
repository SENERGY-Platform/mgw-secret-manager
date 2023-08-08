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

type SecretAlreadyLoaded struct {
	SecretID string
	Path     string
}

func (error SecretAlreadyLoaded) Error() string {
	return fmt.Sprintf("Secret with ID %s already loaded to %s", error.SecretID, error.Path)
}

type SecretDoesNotExistsInFilesystem struct {
	SecretID string
	Path     string
}

func (error SecretDoesNotExistsInFilesystem) Error() string {
	return fmt.Sprintf("Secret with ID %s does not exist at %s", error.SecretID, error.Path)
}

type EncryptionIsDisabled struct{}

func (error EncryptionIsDisabled) Error() string {
	return "Encryption is disabled"
}
