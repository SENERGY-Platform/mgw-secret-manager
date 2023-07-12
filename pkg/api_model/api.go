package api_model

type SecretRequest struct {
	Name       string `json:"name"`
	Value      string `json:"value"`
	SecretType string `json:"type"`
}

type ShortSecret struct {
	Name       string  `json:"name"`
	SecretType string  `json:"type"`
	ID         string  `json:"id"`
	Path       string  `json:"path"`
	Item       *string `json:"item"`
}

type Secret struct {
	Name       string  `json:"name"`
	Value      string  `json:"value"`
	SecretType string  `json:"type"`
	ID         string  `json:"id"`
	Path       string  `json:"path"`
	Item       *string `json:"item"`
}

type SecretPostRequest struct {
	ID        string  `json:"id"`
	Item      *string `json:"item"`
	Reference string  `json:"ref"`
}
