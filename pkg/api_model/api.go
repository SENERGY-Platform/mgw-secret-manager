package api_model

type SecretCreateRequest struct {
	Name       string `json:"name"`
	Value      string `json:"value"`
	SecretType string `json:"type"`
}

type Secret struct {
	Name       string `json:"name"`
	SecretType string `json:"type"`
	ID         string `json:"id"`
}

type SecretVariant struct {
	Secret
	Item *string `json:"item"`
}

type SecretPathVariant struct {
	SecretVariant
	Path string `json:"path"`
}

type SecretValueVariant struct {
	SecretVariant
	Value string `json:"value"`
}

type SecretVariantRequest struct {
	ID        string  `json:"id"`
	Item      *string `json:"item"`
	Reference string  `json:"ref"`
}
