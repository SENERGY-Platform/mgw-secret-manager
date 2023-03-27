package api

import "fmt"

type MissingQueryError struct {
	Parameter string
}

func (error MissingQueryError) Error() string {
	return fmt.Sprintf("Query parameter %s is missing", error.Parameter)
}

type SecretRequest struct {
	Name       string `json:"name"`
	Value      string `json:"value"`
	SecretType string `json:"type"`
}

type MissingEncryptionKey struct{}

func (error MissingEncryptionKey) Error() string {
	return fmt.Sprintf("Encryption Key is missing -> POST at /key")
}
