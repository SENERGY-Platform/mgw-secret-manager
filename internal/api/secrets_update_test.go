package api

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"path/filepath"
	"strings"
	"testing"

	"github.com/SENERGY-Platform/mgw-secret-manager/internal/config"
	"github.com/SENERGY-Platform/mgw-secret-manager/pkg/api_model"
	"github.com/stretchr/testify/assert"
)

type a struct {
	ExistingSecret api_model.SecretRequest
	ChangedSecret  api_model.SecretRequest
	LoadIntoTMPFS  bool
	CaseName       string
}

func TestUpdateSecret(t *testing.T) {
	var config, _ = config.NewConfig(config.Flags.ConfPath)
	config.EnableEncryption = false
	w := httptest.NewRecorder()
	ctx := context.Background()
	reference := "ref"

	testCases := []a{
		{
			LoadIntoTMPFS: false,
			ExistingSecret: api_model.SecretRequest{
				Name:       "name1",
				Value:      "value1",
				SecretType: "type1",
			},
			ChangedSecret: api_model.SecretRequest{
				Name:       "name2",
				Value:      "value2",
				SecretType: "type2",
			},
			CaseName: "Secret was not loaded into TMPFS before update",
		},
		{
			LoadIntoTMPFS: true,
			ExistingSecret: api_model.SecretRequest{
				Name:       "name3",
				Value:      "value3",
				SecretType: "type3",
			},
			ChangedSecret: api_model.SecretRequest{
				Name:       "name4",
				Value:      "value4",
				SecretType: "type4",
			},
			CaseName: "Secret was loaded into TMPFS before update",
		},
		// TODO add case with credentials
	}

	for _, tc := range testCases {
		t.Run(tc.CaseName, func(t *testing.T) {
			router, dbHandler, secretHandler := InitServer(config)
			defer dbHandler.Cleanup()

			_, shortSecret := SetupDummySecret(t, tc.ExistingSecret.Name, tc.ExistingSecret.Value, tc.ExistingSecret.SecretType, secretHandler)
			secretID := shortSecret.ID

			if tc.LoadIntoTMPFS {
				// Load the secret into TMPFS and check whether the value is new
				secretHandler.LoadSecretToFileSystem(context.Background(), api_model.SecretPostRequest{ID: secretID, Reference: reference})
			}

			body, err := json.Marshal(tc.ChangedSecret)
			if err != nil {
				t.Errorf(err.Error())
			}

			req, _ := http.NewRequest("PUT", "/secrets/"+secretID, strings.NewReader(string(body)))
			router.ServeHTTP(w, req)

			var response string

			json.NewDecoder(w.Body).Decode(&response)

			assert.Equal(t, 200, w.Code)

			secretFromDB, err := secretHandler.GetSecret(ctx, api_model.SecretPostRequest{ID: secretID})
			if err != nil {
				t.Errorf(err.Error())
			}
			assert.Equal(t, tc.ChangedSecret.Name, secretFromDB.Name)
			assert.Equal(t, tc.ChangedSecret.SecretType, secretFromDB.SecretType)
			assert.Equal(t, secretID, secretFromDB.ID)

			if tc.LoadIntoTMPFS {
				pathToSecretInTMPFS := secretHandler.BuildTMPFSOutputPath(api_model.SecretPostRequest{ID: secretID, Reference: reference})
				fullSecretPath := filepath.Join(config.TMPFSPath, pathToSecretInTMPFS)
				fileContent, err := ioutil.ReadFile(fullSecretPath)
				if err != nil {
					t.Errorf(err.Error())
					return
				}

				assert.Equal(t, tc.ChangedSecret.Value, string(fileContent))
			}
		})
	}
}
