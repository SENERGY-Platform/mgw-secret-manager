package http_handler

import (
	"github.com/SENERGY-Platform/mgw-secret-manager/internal/api"
	"github.com/SENERGY-Platform/mgw-secret-manager/pkg/api_model"
	"github.com/gin-gonic/gin"
	"net/http"
)

func InitPathVariant(api *api.Api) (string, string, gin.HandlerFunc) {
	return http.MethodPost, api_model.InitPathVariantPath, func(gc *gin.Context) {

		ok := CheckIfEncryptionKeyExists(gc, api)
		if !ok {
			return
		}

		secretVariantRequest, err := ParseVariantRequest(gc)
		if err != nil {
			gc.AbortWithError(http.StatusInternalServerError, err)
			return
		}

		secretPathVariant, err := api.SecretHandler.InitPathVariant(gc.Request.Context(), secretVariantRequest)
		if err != nil {
			gc.AbortWithError(http.StatusInternalServerError, err)
			return
		}

		gc.JSON(http.StatusOK, secretPathVariant)
	}
}

func LoadPathVariant(api *api.Api) (string, string, gin.HandlerFunc) {
	return http.MethodPost, api_model.LoadPathVariantPath, func(gc *gin.Context) {

		ok := CheckIfEncryptionKeyExists(gc, api)
		if !ok {
			return
		}

		secretVariantRequest, err := ParseVariantRequest(gc)
		if err != nil {
			gc.AbortWithError(http.StatusInternalServerError, err)
			return
		}

		err = api.SecretHandler.LoadSecretToFileSystem(gc.Request.Context(), secretVariantRequest)
		if err != nil {
			gc.AbortWithError(http.StatusInternalServerError, err)
			return
		}

		gc.String(http.StatusOK, "")
	}
}

func DeleteSecretFromTMPFS(api *api.Api) (string, string, gin.HandlerFunc) {
	return http.MethodPost, api_model.UnLoadPathVariantPath, func(gc *gin.Context) {
		secretVariantRequest, err := ParseVariantRequest(gc)
		if err != nil {
			gc.AbortWithError(http.StatusInternalServerError, err)
			return
		}

		api.SecretHandler.RemoveSecretFromFileSystem(gc.Request.Context(), secretVariantRequest)
	}
}

func CleanReferenceDirectory(api *api.Api) (string, string, gin.HandlerFunc) {
	return http.MethodPost, api_model.CleanPathVariantsPath, func(gc *gin.Context) {
		reference := gc.Query("reference")
		err := api.SecretHandler.CleanReferenceDirectory(gc.Request.Context(), reference)
		if err != nil {
			gc.AbortWithError(http.StatusInternalServerError, err)
			return
		}
	}
}
