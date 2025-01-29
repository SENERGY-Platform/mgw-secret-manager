package standard

import (
	"github.com/SENERGY-Platform/mgw-secret-manager/internal/api"
	"github.com/SENERGY-Platform/mgw-secret-manager/internal/http_handler/util"
	"github.com/SENERGY-Platform/mgw-secret-manager/pkg/api_model"
	"github.com/gin-gonic/gin"
	"net/http"
)

// InitPathVariant godoc
// @Summary Init secret file
// @Description Create a placeholder file for a secret.
// @Tags Secrets
// @Accept json
// @Produce	json
// @Param request body api_model.SecretVariantRequest true "request"
// @Success	200 {object} api_model.SecretPathVariant "secret file info"
// @Failure	404 {string} string "error message"
// @Failure	500 {string} string "error message"
// @Router /path-variant/init [post]
func InitPathVariant(api *api.Api) (string, string, gin.HandlerFunc) {
	return http.MethodPost, api_model.InitPathVariantPath, func(gc *gin.Context) {

		ok := util.CheckIfEncryptionKeyExists(gc, api)
		if !ok {
			return
		}

		secretVariantRequest, err := util.ParseVariantRequest(gc)
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

// LoadPathVariant godoc
// @Summary Write secret file
// @Description Write secret value to file. File must be initialised first.
// @Tags Secrets
// @Accept json
// @Param request body api_model.SecretVariantRequest true "request"
// @Success	200
// @Failure	404 {string} string "error message"
// @Failure	500 {string} string "error message"
// @Router /path-variant/load [post]
func LoadPathVariant(api *api.Api) (string, string, gin.HandlerFunc) {
	return http.MethodPost, api_model.LoadPathVariantPath, func(gc *gin.Context) {

		ok := util.CheckIfEncryptionKeyExists(gc, api)
		if !ok {
			return
		}

		secretVariantRequest, err := util.ParseVariantRequest(gc)
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

// DeleteSecretFromTMPFS godoc
// @Summary Delete secret file
// @Description Remove a secret file.
// @Tags Secrets
// @Accept json
// @Param request body api_model.SecretVariantRequest true "request"
// @Success	200
// @Failure	500 {string} string "error message"
// @Router /path-variant/unload [delete]
func DeleteSecretFromTMPFS(api *api.Api) (string, string, gin.HandlerFunc) {
	return http.MethodPost, api_model.UnLoadPathVariantPath, func(gc *gin.Context) {
		secretVariantRequest, err := util.ParseVariantRequest(gc)
		if err != nil {
			gc.AbortWithError(http.StatusInternalServerError, err)
			return
		}

		api.SecretHandler.RemoveSecretFromFileSystem(gc.Request.Context(), secretVariantRequest)
	}
}

// CleanReferenceDirectory godoc
// @Summary Delete secret files
// @Description Remove all secret files with the same reference.
// @Tags Secrets
// @Param reference query string true "reference"
// @Success	200
// @Failure	500 {string} string "error message"
// @Router /path-variant/clean [post]
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
