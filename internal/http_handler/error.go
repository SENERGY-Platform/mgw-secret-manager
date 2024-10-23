package http_handler

import (
	"errors"
	"github.com/SENERGY-Platform/mgw-secret-manager/internal/custom_errors"
	"net/http"
)

func GetStatusCode(err error) int {
	var nsf *custom_errors.NoSecretFound
	if errors.As(err, &nsf) {
		return http.StatusNotFound
	}
	var mek *custom_errors.MissingEncryptionKey
	if errors.As(err, &mek) {
		return http.StatusBadRequest
	}
	var sal *custom_errors.SecretAlreadyLoaded
	if errors.As(err, &sal) {
		return http.StatusBadRequest
	}
	var snf *custom_errors.SecretDoesNotExistsInFilesystem
	if errors.As(err, &snf) {
		return http.StatusNotFound
	}
	return http.StatusInternalServerError
}
