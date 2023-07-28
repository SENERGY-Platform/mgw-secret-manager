package httpHandler

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/SENERGY-Platform/mgw-secret-manager/internal/api"
	"github.com/SENERGY-Platform/mgw-secret-manager/internal/customErrors"
	"github.com/SENERGY-Platform/mgw-secret-manager/internal/logger"
	"github.com/SENERGY-Platform/mgw-secret-manager/pkg/api_model"
	"github.com/SENERGY-Platform/mgw-secret-manager/pkg/interfaces"

	"github.com/gin-gonic/gin"
)

func ParseVariantRequest(gc *gin.Context) (api_model.SecretVariantRequest, error) {
	body, err := ioutil.ReadAll(gc.Request.Body)
	if err != nil {
		logger.Logger.Errorf(err.Error())
		gc.AbortWithError(http.StatusInternalServerError, err)
		return api_model.SecretVariantRequest{}, err
	}

	var secretVariantRequest api_model.SecretVariantRequest
	err = json.Unmarshal(body, &secretVariantRequest)
	if err != nil {
		logger.Logger.Errorf(err.Error())
		gc.AbortWithError(http.StatusInternalServerError, err)
		return api_model.SecretVariantRequest{}, err
	}

	return secretVariantRequest, nil
}

func ParseSecretCreateRequest(gc *gin.Context) (secretRequest api_model.SecretCreateRequest, err error) {
	body, err := ioutil.ReadAll(gc.Request.Body)
	if err != nil {
		gc.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	err = json.Unmarshal(body, &secretRequest)
	if err != nil {
		logger.Logger.Errorf(err.Error())
		gc.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	return
}

func CheckIfEncryptionKeyExists(gc *gin.Context, api *api.Api) bool {
	if api.SecretHandler.Key == nil && api.Config.EnableEncryption == true {
		gc.AbortWithError(http.StatusInternalServerError, customErrors.MissingEncryptionKey{})
		gc.Abort()
		return false
	}
	return true
}

func GetTypes(api interfaces.ApiInterface) gin.HandlerFunc {
	return func(gc *gin.Context) {
		gc.JSON(http.StatusOK, []map[string]string{{"id": "certificate", "name": "Certificate"}, {"id": "basic-auth", "name": "Credentials"}, {"id": "api-key", "name": "API Key"}})
	}
}
