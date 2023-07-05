package client

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/SENERGY-Platform/mgw-secret-manager/pkg/api_model"
)

type HttpClient interface {
	Do(req *http.Request) (*http.Response, error)
}

type RealClient struct {
	BaseUrl    string
	HTTPClient HttpClient
}

func (c *RealClient) StoreSecret(name string, value string, secretType string) (err error, errCode int) {
	secretRequest := api_model.SecretRequest{
		Name:       name,
		Value:      value,
		SecretType: secretType,
	}
	body, err := json.Marshal(secretRequest)
	if err != nil {
		return err, http.StatusInternalServerError
	}

	req, err := http.NewRequest(http.MethodPost, c.BaseUrl+"/secrets", strings.NewReader(string(body)))
	if err != nil {
		return err, http.StatusInternalServerError
	}
	return do(req, c.HTTPClient)
}

func (c *RealClient) LoadSecretToTMPFS(secretID string) (fullTMPFSPath string, err error, errCode int) {
	req, err := http.NewRequest(http.MethodPost, c.BaseUrl+"/load", nil)
	q := req.URL.Query()
	q.Add("secret", secretID)
	req.URL.RawQuery = q.Encode()

	if err != nil {
		return "", err, http.StatusInternalServerError
	}
	return doWithResponse[string](req, c.HTTPClient)
}

func (c *RealClient) SetEncryptionKey(encryptionKey []byte) (err error, errCode int) {
	req, err := http.NewRequest(http.MethodPost, c.BaseUrl+"/key", strings.NewReader(string(encryptionKey)))
	if err != nil {
		return err, http.StatusInternalServerError
	}
	return do(req, c.HTTPClient)
}

func (c *RealClient) GetSecrets() (secrets []api_model.ShortSecret, err error, errCode int) {
	req, err := http.NewRequest(http.MethodGet, c.BaseUrl+"/secrets", nil)
	if err != nil {
		return nil, err, http.StatusInternalServerError
	}
	return doWithResponse[[]api_model.ShortSecret](req, c.HTTPClient)
}

func (c *RealClient) UpdateSecret(name string, value string, secretType string, id string) (err error, errCode int) {
	secretRequest := api_model.SecretRequest{
		Name:       name,
		Value:      value,
		SecretType: secretType,
	}
	body, err := json.Marshal(secretRequest)
	if err != nil {
		return err, http.StatusInternalServerError
	}

	req, err := http.NewRequest(http.MethodPut, c.BaseUrl+"/secrets/"+id, strings.NewReader(string(body)))
	if err != nil {
		return err, http.StatusInternalServerError
	}
	return do(req, c.HTTPClient)
}

func (c *RealClient) DeleteSecret(id string) (err error, errCode int) {
	req, err := http.NewRequest(http.MethodDelete, c.BaseUrl+"/secrets/"+id, nil)
	if err != nil {
		return err, http.StatusInternalServerError
	}
	return do(req, c.HTTPClient)
}

func NewClient(url string, httpClient HttpClient) (client Client) {
	return &RealClient{
		BaseUrl:    url,
		HTTPClient: httpClient,
	}
}

func do(req *http.Request, client HttpClient) (err error, code int) {
	_, err = client.Do(req)
	if err != nil {
		return err, http.StatusInternalServerError
	}
	return
}

func doWithResponse[T any](req *http.Request, client HttpClient) (result T, err error, code int) {
	resp, err := client.Do(req)
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
