package httpHandler

import (
	"errors"
	"net/http"

	"github.com/SENERGY-Platform/mgw-secret-manager/internal/customErrors"
)

func GetStatusCode(err error) int {
	var nsf *customErrors.NoSecretFound
	if errors.As(err, &nsf) {
		return http.StatusNotFound
	}
	var mek *customErrors.MissingEncryptionKey
	if errors.As(err, &mek) {
		return http.StatusBadRequest
	}
	var sal *customErrors.SecretAlreadyLoaded
	if errors.As(err, &sal) {
		return http.StatusBadRequest
	}
	var snf *customErrors.SecretDoesNotExistsInFilesystem
	if errors.As(err, &snf) {
		return http.StatusNotFound
	}
	return http.StatusInternalServerError
}
