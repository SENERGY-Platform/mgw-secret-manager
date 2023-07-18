package api

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/SENERGY-Platform/mgw-secret-manager/internal/config"
	"github.com/SENERGY-Platform/mgw-secret-manager/pkg/api_model"

	"github.com/stretchr/testify/assert"
)

func TestStoreSecret(t *testing.T) {
	var config, _ = config.NewConfig(config.Flags.ConfPath)
	config.EnableEncryption = false
	ctx := context.Background()

	testCases := []api_model.SecretCreateRequest{
		{Name: "name1", Value: "value1", SecretType: "Type1"},
		{Name: "name1", Value: "value1", SecretType: "Type1"},
		{Name: "name2", Value: "value2", SecretType: "Type2"},
	}
	for _, tc := range testCases {
		w := httptest.NewRecorder()
		router, dbHandler, secretHandler := InitServer(config)
		defer dbHandler.Cleanup()

		body, err := json.Marshal(tc)
		if err != nil {
			t.Errorf(err.Error())
			return
		}

		req, _ := http.NewRequest("POST", "/secrets", strings.NewReader(string(body)))
		router.ServeHTTP(w, req)

		secretID := w.Body.String()

		assert.Equal(t, 200, w.Code)

		secretFromDB, err := secretHandler.GetSecret(ctx, api_model.SecretVariantRequest{ID: secretID})
		if err != nil {
			t.Errorf(err.Error())
			return
		}
		assert.Equal(t, tc.Name, secretFromDB.Name)
		assert.Equal(t, tc.SecretType, secretFromDB.SecretType)
		assert.Equal(t, secretID, secretFromDB.ID)
	}
}
