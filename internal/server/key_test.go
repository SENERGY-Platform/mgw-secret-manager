package server

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/SENERGY-Platform/mgw-secret-manager/internal/config"

	"github.com/SENERGY-Platform/mgw-secret-manager/internal/logger"
	"github.com/SENERGY-Platform/mgw-secret-manager/test"

	"github.com/stretchr/testify/assert"
)

var testConfig, _ = config.NewConfig(config.Flags.ConfPath)
var _, _ = logger.InitLogger(testConfig.Logger)

func TestSetKey(t *testing.T) {
	testConfig.EnableEncryption = true
	router, dbHandler, _ := InitServer(testConfig)
	defer dbHandler.Cleanup()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", fmt.Sprintf("/key"), strings.NewReader(string(test.EncryptionKey)))
	router.ServeHTTP(w, req)
	assert.Equal(t, w.Code, 200)
}

func TestSetKeyWithDisabledEncryption(t *testing.T) {
	testConfig.EnableEncryption = false
	router, dbHandler, _ := InitServer(testConfig)
	defer dbHandler.Cleanup()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", fmt.Sprintf("/key"), strings.NewReader(string(test.EncryptionKey)))
	router.ServeHTTP(w, req)
	assert.Equal(t, w.Code, 500)
}

func TestMissingKey(t *testing.T) {
	testConfig.EnableEncryption = true
	router, dbHandler, _ := InitServer(testConfig)
	defer dbHandler.Cleanup()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", fmt.Sprintf("/key"), nil)
	router.ServeHTTP(w, req)
	assert.Equal(t, w.Code, 500)
}
