package secretHandler

import (
	"context"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/SENERGY-Platform/mgw-secret-manager/internal/customErrors"
	"github.com/SENERGY-Platform/mgw-secret-manager/internal/files"
	"github.com/SENERGY-Platform/mgw-secret-manager/internal/logger"
	"github.com/SENERGY-Platform/mgw-secret-manager/pkg/api_model"
)

func (secretHandler *SecretHandler) BuildTMPFSOutputPath(secretPostRequest api_model.SecretVariantRequest) string {
	// Path must be unique for each deployment (Reference) and Secret Item in case of Credentials (Username/Password)
	fileName := "value"
	if secretPostRequest.Item != nil {
		fileName = *secretPostRequest.Item
	}
	return filepath.Join(secretPostRequest.Reference, secretPostRequest.ID, fileName)
}

func (secretHandler *SecretHandler) LoadSecretToFileSystem(ctx context.Context, secretPostRequest api_model.SecretVariantRequest) (err error) {
	// Load secret to TMPFS if it does not exist already

	logger.Logger.Debugf("Get Secret Value and load into TMPFS")
	relativeFilePath := secretHandler.BuildTMPFSOutputPath(secretPostRequest)
	fullOutputPath := filepath.Join(secretHandler.TMPFSPath, relativeFilePath)

	if _, err = os.Stat(fullOutputPath); err == nil {
		err = customErrors.SecretAlreadyLoaded{SecretID: secretPostRequest.ID, Path: relativeFilePath}
		return
	}

	err = secretHandler.SaveSecretToFileSystem(ctx, secretPostRequest, fullOutputPath)
	return
}

func (secretHandler *SecretHandler) SaveSecretToFileSystem(ctx context.Context, secretPostRequest api_model.SecretVariantRequest, fullOutputPath string) (err error) {
	// Get secret value and write file to TMPFS

	secret, errGet := secretHandler.GetFullSecret(ctx, secretPostRequest)
	if errGet != nil {
		return errGet
	}

	logger.Logger.Debugf("Load Secret: %s to %s", secret.ID, fullOutputPath)
	err = files.WriteToFile(string(secret.Value), fullOutputPath)
	if err != nil {
		logger.Logger.Errorf("Write to TMPFS failed: %s", err.Error())
	}
	return
}

func (secretHandler *SecretHandler) UpdateExistingSecretInTMPFS(ctx context.Context, secretID string) (err error) {
	// Reload existing secrets to TMPFS so that services have access to the newest value

	logger.Logger.Debugf("Update existing secret files in TMPFS")

	referenceDirectories, _ := ioutil.ReadDir(secretHandler.TMPFSPath)
	for _, referenceDirectory := range referenceDirectories {
		pathToFiles := filepath.Join(secretHandler.TMPFSPath, referenceDirectory.Name(), secretID)
		files, _ := ioutil.ReadDir(pathToFiles)
		logger.Logger.Debugf(pathToFiles)
		// TODO if exists
		for _, fileName := range files {
			logger.Logger.Debugf(fileName.Name())
			secretPostRequest := api_model.SecretVariantRequest{ID: secretID, Reference: referenceDirectory.Name()}

			// "Value" is the reserved secret key for single value secrets
			if fileName.Name() != "value" {
				fileNameValue := fileName.Name()
				secretPostRequest.Item = &fileNameValue
			}
			relativeFilePath := secretHandler.BuildTMPFSOutputPath(secretPostRequest)
			fullOutputPath := filepath.Join(secretHandler.TMPFSPath, relativeFilePath)
			err = secretHandler.SaveSecretToFileSystem(ctx, secretPostRequest, fullOutputPath)
		}

	}
	return
}

func (secretHandler *SecretHandler) RemoveSecretFromFileSystem(ctx context.Context, secretPostRequest api_model.SecretVariantRequest) (relativeFilePath string, err error) {
	logger.Logger.Debugf("Remove secret from TMPFS")
	relativeFilePath = secretHandler.BuildTMPFSOutputPath(secretPostRequest)
	fullOutputPath := filepath.Join(secretHandler.TMPFSPath, relativeFilePath)

	err = os.Remove(fullOutputPath)
	return
}
