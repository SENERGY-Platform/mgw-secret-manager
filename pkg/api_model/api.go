package api_model

type SecretRequest struct {
	Name       string `json:"name"`
	Value      string `json:"value"`
	SecretType string `json:"type"`
}

type ShortSecret struct {
	Name       string `json:"name"`
	SecretType string `json:"type"`
	ID         string `json:"id"`
}
