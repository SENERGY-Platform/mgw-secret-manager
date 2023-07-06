package api_model

type SecretRequest struct {
	Name       string `json:"name"`
	Value      string `json:"value"`
	FileName   string `json:"file_name"`
	SecretType string `json:"type"`
}

type ShortSecret struct {
	Name       string `json:"name"`
	SecretType string `json:"type"`
	FileName   string `json:"file_name"`
	ID         string `json:"id"`
}

type Secret struct {
	Name       string `json:"name"`
	Value      string `json:"value"`
	FileName   string `json:"file_name"`
	SecretType string `json:"type"`
	ID         string `json:"id"`
}

type SecretPostRequest struct {
	ID      string             `json:"id"`
	Options *map[string]string `json:"options"`
}
