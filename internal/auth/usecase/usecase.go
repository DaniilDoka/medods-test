package auth_usercase

import (
	"crypto/rand"
	"encoding/base64"
	"medods-test/config"
	"medods-test/internal/auth"
	auth_models "medods-test/internal/auth/models"
	"medods-test/pkg/logger"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type usecase struct {
	repo   auth.Repository
	cfg    *config.Config
	logger *logger.Logger
}

func NewUsecase(repo auth.Repository, cfg *config.Config, logger *logger.Logger) auth.Usecase {
	return &usecase{
		repo:   repo,
		cfg:    cfg,
		logger: logger,
	}
}

func generateRandomRefreshToken() ([]byte, error) {
	bytes := make([]byte, 72)
	if _, err := rand.Read(bytes); err != nil {
		return nil, err
	}
	return bytes, nil
}

func (u *usecase) Signin(params *auth_models.SigninParams) (*auth_models.SigninResponse, error) {
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS512, jwt.MapClaims{
		"ip":   params.UserIp,
		"guid": params.Guid,
	})
	accessString, err := accessToken.SignedString([]byte(u.cfg.Server.Key))
	if err != nil {
		return nil, err
	}
	exp := time.Now().Add(time.Hour * 24)

	refreshToken, err := generateRandomRefreshToken()
	if err != nil {
		return nil, err
	}

	hash, err := bcrypt.GenerateFromPassword(refreshToken, bcrypt.DefaultCost)

	if err != nil {
		return nil, err
	}

	if err := u.repo.PutRefreshToken(params.Guid, string(hash), exp); err != nil {
		return nil, err
	}
	encodedRefresh := base64.StdEncoding.EncodeToString(refreshToken)

	return &auth_models.SigninResponse{
		Access:  accessString,
		Refresh: encodedRefresh,
	}, nil
}

func (u *usecase) Refresh(params *auth_models.RefreshParams) (*auth_models.RefreshResponse, error) {
	return nil, nil
}
