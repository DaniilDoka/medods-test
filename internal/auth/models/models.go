package auth_models

type SigninParams struct {
	Guid   string
	UserIp string
}

type SigninResponse struct {
	Access  string `json:"access"`
	Refresh string `json:"refresh"`
}

type RefreshParams struct {
	Refresh string
	UserIp  string
}

type RefreshResponse struct {
	Access string `json:"access"`
}
