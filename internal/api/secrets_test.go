package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/SENERGY-Platform/mgw-secret-manager/internal/core"
	"github.com/SENERGY-Platform/mgw-secret-manager/internal/model"
	"github.com/SENERGY-Platform/mgw-secret-manager/test"

	srv_base "github.com/SENERGY-Platform/go-service-base/srv-base"

	"github.com/stretchr/testify/assert"
)

var _, _ = srv_base.InitLogger(test.TestConfig.Logger)

func TestLoadSecret(t *testing.T) {
	defer dbHandler.Cleanup()

	// Setup dummy secret
	secretName := "secret"
	secret := core.CreateSecret(secretName, "geheim")
	err := core.StoreSecret(&secret, dbHandler, test.MasterKey)
	if err != nil {
		t.Errorf(err.Error())
	}

	router := GetTestRouter()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", fmt.Sprintf("/load?secret=%s", secretName), nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)

	pathToSecretInTMPFS := filepath.Join(test.TestConfig.TMPFSPath, secret.ID)
	_, err = os.Stat(pathToSecretInTMPFS)
	assert.Equal(t, nil, err)

}

func TestLoadSecretMissingQuery(t *testing.T) {
	router := GetTestRouter()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/load", nil)
	router.ServeHTTP(w, req)
	assert.Equal(t, 500, w.Code)
}

func TestPostValidSecret(t *testing.T) {
	defer dbHandler.Cleanup()

	router := GetTestRouter()
	w := httptest.NewRecorder()

	secret := model.Secret{
		Name:  "test",
		Value: "value",
	}
	body, err := json.Marshal(secret)
	if err != nil {

	}
	req, _ := http.NewRequest("POST", "/secret", strings.NewReader(string(body)))
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
}
