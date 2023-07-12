package secretHandler

import (
	"context"
	"fmt"
	"os"
	"path/filepath"

	"github.com/SENERGY-Platform/mgw-secret-manager/internal/customErrors"
	"github.com/SENERGY-Platform/mgw-secret-manager/internal/files"
	"github.com/SENERGY-Platform/mgw-secret-manager/internal/logger"
	"github.com/SENERGY-Platform/mgw-secret-manager/pkg/api_model"
)

func BuildTMPFSOutputPath(secretPostRequest api_model.SecretPostRequest) string {
	// Path must be unique for each deployment (Reference) and Secret Item in case of Credentials (Username/Password)
	fileName := fmt.Sprintf("%s_%s_%s", secretPostRequest.ID, secretPostRequest.Reference, secretPostRequest.Item)
	return filepath.Join(secretPostRequest.ID, fileName)
}

func (secretHandler *SecretHandler) LoadSecretToFileSystem(ctx context.Context, secretPostRequest api_model.SecretPostRequest) (err error) {
	logger.Logger.Debugf("Get Secret Value and load into TMPFS")
	relativeFilePath := BuildTMPFSOutputPath(secretPostRequest)
	fullOutputPath := filepath.Join(secretHandler.TMPFSPath, relativeFilePath)

	if _, err = os.Stat(fullOutputPath); err == nil {
		err = customErrors.SecretAlreadyLoaded{SecretID: secretPostRequest.ID, Path: relativeFilePath}
		return
	}

	err = secretHandler.SaveSecretToFileSystem(ctx, secretPostRequest, fullOutputPath)
	return
}

func (secretHandler *SecretHandler) SaveSecretToFileSystem(ctx context.Context, secretPostRequest api_model.SecretPostRequest, fullOutputPath string) (err error) {
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

func (secretHandler *SecretHandler) UpdateExistingSecretInTMPFS(ctx context.Context, secretPostRequest api_model.SecretPostRequest, override bool) (err error) {
	logger.Logger.Debugf("Update Secret Value in TMPFS")
	relativeFilePath := BuildTMPFSOutputPath(secretPostRequest)
	fullOutputPath := filepath.Join(secretHandler.TMPFSPath, relativeFilePath)

	if _, err = os.Stat(fullOutputPath); err != nil {
		return nil
	}

	err = secretHandler.SaveSecretToFileSystem(ctx, secretPostRequest, fullOutputPath)
	return
}

func (secretHandler *SecretHandler) RemoveSecretFromFileSystem(ctx context.Context, secretPostRequest api_model.SecretPostRequest) (relativeFilePath string, err error) {
	logger.Logger.Debugf("Remove secret from TMPFS")
	relativeFilePath = BuildTMPFSOutputPath(secretPostRequest)
	fullOutputPath := filepath.Join(secretHandler.TMPFSPath, relativeFilePath)

	err = os.Remove(fullOutputPath)
	return
}
