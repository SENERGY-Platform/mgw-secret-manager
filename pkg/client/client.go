package client

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"secret-manager/internal/core"
	"strings"
)

type RealClient struct {
	baseUrl string
}

func (c *RealClient) StoreSecret(name string, value string) (err error, errCode int) {
	secret := core.CreateSecret(name, value)
	body, err := json.Marshal(secret)
	if err != nil {
		return err, http.StatusInternalServerError
	}

	req, err := http.NewRequest(http.MethodPost, c.baseUrl+"/secret", strings.NewReader(string(body)))
	if err != nil {
		return err, http.StatusInternalServerError
	}
	return do(req)
}

func (c *RealClient) LoadSecretToTMPFS(secretName string) (fullTMPFSPath string, err error, errCode int) {
	req, err := http.NewRequest(http.MethodPost, c.baseUrl+"/load", nil)
	q := req.URL.Query()
	q.Add("secret", secretName)
	req.URL.RawQuery = q.Encode()

	if err != nil {
		return "", err, http.StatusInternalServerError
	}
	return doWithResponse[string](req)
}

func (c *RealClient) SetEncryptionKey(encryptionKey []byte) (err error, errCode int) {
	req, err := http.NewRequest(http.MethodPost, c.baseUrl+"/key", strings.NewReader(string(encryptionKey)))
	if err != nil {
		return err, http.StatusInternalServerError
	}
	return do(req)
}

func NewClient() (client Client) {
	return &RealClient{}
}

func do(req *http.Request) (err error, code int) {
	_, err = http.DefaultClient.Do(req)
	if err != nil {
		return err, http.StatusInternalServerError
	}
	return
}

func doWithResponse[T any](req *http.Request) (result T, err error, code int) {
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return result, err, http.StatusInternalServerError
	}
	defer resp.Body.Close()
	if resp.StatusCode > 299 {
		temp, _ := io.ReadAll(resp.Body) //read error response end ensure that resp.Body is read to EOF
		return result, fmt.Errorf("unexpected statuscode %v: %v", resp.StatusCode, string(temp)), resp.StatusCode
	}
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		_, _ = io.ReadAll(resp.Body) //ensure resp.Body is read to EOF
		return result, err, http.StatusInternalServerError
	}
	return
}
