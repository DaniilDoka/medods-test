package auth_models

type SigninParams struct {
	Guid   string `json:"guid"`
	UserIp string `json:"ip"`
}

type SigninResponse struct {
	Access  string `json:"access"`
	Refresh string `json:"refresh"`
}

type RefreshParams struct {
	Refresh string `json:"refresh"`
}

type RefreshResponse struct {
	Access string `json:"access"`
}
