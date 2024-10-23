package secret_handler

import (
	"context"
	"encoding/json"
	"github.com/SENERGY-Platform/mgw-secret-manager/internal/models"
	"github.com/SENERGY-Platform/mgw-secret-manager/internal/util"
	"github.com/SENERGY-Platform/mgw-secret-manager/pkg/api_model"
)

func (secretHandler *SecretHandler) ExtractValue(ctx context.Context, secretPostRequest api_model.SecretVariantRequest, value string) (extractedValue string, err error) {
	if secretPostRequest.Item == nil {
		return value, nil
	}

	var secretValue models.SecretValue
	err = json.Unmarshal([]byte(value), &secretValue)
	if err != nil {
		util.Logger.Errorf("Secret can not be unmarshaled: %s This can be caused by specifing an Item on a secret that is not saved in JSON", err.Error())
		return
	}

	itemKey := *secretPostRequest.Item
	extractedValue, ok := secretValue[itemKey]
	if !ok {
		util.Logger.Errorf("Item %s does not exist as key in JSON secret", itemKey)
		err = nil
		return
	}
	return extractedValue, nil
}

func (secretHandler *SecretHandler) GetValueVariant(ctx context.Context, secretPostRequest api_model.SecretVariantRequest) (variant api_model.SecretValueVariant, err error) {
	secret, err := secretHandler.GetSecret(ctx, secretPostRequest.ID)
	if err != nil {
		return
	}
	extractedValue, err := secretHandler.ExtractValue(ctx, secretPostRequest, secret.Value)
	if err != nil {
		return
	}

	variant = api_model.SecretValueVariant{
		SecretVariant: api_model.SecretVariant{
			Secret: api_model.Secret{
				Name:       secret.Name,
				SecretType: secret.SecretType,
				ID:         secret.ID,
			},
			Item: secretPostRequest.Item,
		},
		Value: extractedValue,
	}
	return
}
