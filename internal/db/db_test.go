package db

import (
	"context"
	"testing"

	"github.com/SENERGY-Platform/mgw-secret-manager/internal/models"
	"github.com/SENERGY-Platform/mgw-secret-manager/internal/util"

	"github.com/stretchr/testify/assert"
)

var testConfig, _ = util.NewConfig(util.Flags.ConfPath)
var _, _ = util.InitLogger(testConfig.Logger)

func TestSetSecret(t *testing.T) {
	testHandler, err := NewDBHandler(testConfig)
	defer testHandler.Cleanup()
	if err != nil {
		t.Errorf(err.Error())
	}
	secret := &models.EncryptedSecret{
		Name:  "test",
		Value: make([]byte, 2),
	}
	ctx := context.Background()
	err = testHandler.SetSecret(ctx, secret)
	assert.Equal(t, err, nil)
}

func TestGetSecret(t *testing.T) {
	secretID := "id"
	testHandler, err := NewDBHandler(testConfig)
	ctx := context.Background()
	defer testHandler.Cleanup()
	if err != nil {
		t.Errorf(err.Error())
	}
	secret := &models.EncryptedSecret{
		Name:  "name",
		Value: make([]byte, 2),
		ID:    secretID,
	}
	err = testHandler.SetSecret(ctx, secret)

	storedSecret, err := testHandler.GetSecret(ctx, secretID)
	assert.Nil(t, err)
	assert.Equal(t, secret, storedSecret)
}
