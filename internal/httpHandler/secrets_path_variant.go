package httpHandler

import (
	"net/http"

	"github.com/SENERGY-Platform/mgw-secret-manager/internal/api"
	"github.com/gin-gonic/gin"
)

func InitPathVariant(api *api.Api) gin.HandlerFunc {
	return func(gc *gin.Context) {

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

func LoadPathVariant(api *api.Api) gin.HandlerFunc {
	return func(gc *gin.Context) {

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

func DeleteSecretFromTMPFS(api *api.Api) gin.HandlerFunc {
	return func(gc *gin.Context) {
		secretVariantRequest, err := ParseVariantRequest(gc)
		if err != nil {
			gc.AbortWithError(http.StatusInternalServerError, err)
			return
		}

		api.SecretHandler.RemoveSecretFromFileSystem(gc.Request.Context(), secretVariantRequest)
	}
}

func CleanReferenceDirectory(api *api.Api) gin.HandlerFunc {
	return func(gc *gin.Context) {
		reference := gc.Query("reference")
		err := api.SecretHandler.CleanReferenceDirectory(gc.Request.Context(), reference)
		if err != nil {
			gc.AbortWithError(http.StatusInternalServerError, err)
			return
		}
	}
}
