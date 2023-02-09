package api

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"secret-manager/test"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPostKey(t *testing.T) {
	defer dbHandler.Cleanup()

	router := GetTestRouter()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", fmt.Sprintf("/key"), strings.NewReader(string(test.EncryptionKey)))
	router.ServeHTTP(w, req)
	assert.Equal(t, w.Code, 200)
}
