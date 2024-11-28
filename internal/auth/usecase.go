package auth

import auth_models "medods-test/internal/auth/models"

type Usecase interface {
	Signin(*auth_models.SigninParams) (*auth_models.SigninResponse, error)
	Refresh(*auth_models.RefreshParams) (*auth_models.RefreshResponse, error)
}
