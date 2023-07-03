package api

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	srv_base "github.com/SENERGY-Platform/go-service-base/srv-base"
	"github.com/SENERGY-Platform/mgw-secret-manager/internal/config"
	"github.com/SENERGY-Platform/mgw-secret-manager/test"

	"github.com/stretchr/testify/assert"
)

var _, _ = srv_base.InitLogger(testConfig.Logger)

func TestPostKey(t *testing.T) {
	var config, _ = config.NewConfig(nil)
	config.EnableEncryption = true
	router, dbHandler, _ := InitServer(config)
	defer dbHandler.Cleanup()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", fmt.Sprintf("/key"), strings.NewReader(string(test.EncryptionKey)))
	router.ServeHTTP(w, req)
	assert.Equal(t, w.Code, 200)
}

func TestPostKeyWithDisabledEncryption(t *testing.T) {
	var config, _ = config.NewConfig(nil)
	config.EnableEncryption = false
	router, dbHandler, _ := InitServer(config)
	defer dbHandler.Cleanup()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", fmt.Sprintf("/key"), strings.NewReader(string(test.EncryptionKey)))
	router.ServeHTTP(w, req)
	assert.Equal(t, w.Code, 500)
}
