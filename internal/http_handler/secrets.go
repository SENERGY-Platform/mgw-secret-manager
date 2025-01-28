package http_handler

import (
	"encoding/json"
	"github.com/SENERGY-Platform/mgw-secret-manager/internal/api"
	"github.com/SENERGY-Platform/mgw-secret-manager/internal/custom_errors"
	"github.com/SENERGY-Platform/mgw-secret-manager/internal/util"
	"github.com/SENERGY-Platform/mgw-secret-manager/pkg/api_model"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
)

func ParseVariantRequest(gc *gin.Context) (api_model.SecretVariantRequest, error) {
	body, err := io.ReadAll(gc.Request.Body)
	if err != nil {
		util.Logger.Errorf(err.Error())
		gc.AbortWithError(http.StatusInternalServerError, err)
		return api_model.SecretVariantRequest{}, err
	}

	util.Logger.Debugf("Payload: %s", string(body))

	var secretVariantRequest api_model.SecretVariantRequest
	err = json.Unmarshal(body, &secretVariantRequest)
	if err != nil {
		util.Logger.Errorf(err.Error())
		gc.AbortWithError(http.StatusInternalServerError, err)
		return api_model.SecretVariantRequest{}, err
	}

	return secretVariantRequest, nil
}

func ParseSecretCreateRequest(gc *gin.Context) (secretRequest api_model.SecretCreateRequest, err error) {
	body, err := io.ReadAll(gc.Request.Body)
	if err != nil {
		gc.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	err = json.Unmarshal(body, &secretRequest)
	if err != nil {
		util.Logger.Errorf(err.Error())
		gc.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	return
}

func CheckIfEncryptionKeyExists(gc *gin.Context, api *api.Api) bool {
	if api.SecretHandler.Key == nil && api.Config.EnableEncryption == true {
		gc.AbortWithError(http.StatusInternalServerError, custom_errors.MissingEncryptionKey{})
		gc.Abort()
		return false
	}
	return true
}

func GetTypes(api *api.Api) (string, string, gin.HandlerFunc) {
	return http.MethodGet, "/types", func(gc *gin.Context) {
		gc.JSON(http.StatusOK, []map[string]string{{"id": "certificate", "name": "Certificate"}, {"id": "basic-auth", "name": "Credentials"}, {"id": "api-key", "name": "API Key"}, {"id": "client-id", "name": "Client ID"}, {"id": "private-key", "name": "Private Key"}})
	}
}
