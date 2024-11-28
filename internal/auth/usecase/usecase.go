package auth_usercase

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"medods-test/config"
	"medods-test/internal/auth"
	auth_models "medods-test/internal/auth/models"
	"medods-test/pkg/logger"
	"sync"
	"time"

	"github.com/go-resty/resty/v2"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type usecase struct {
	repo    auth.Repository
	mu      sync.Mutex
	tokens  map[string]string // [refresh]access
	rClient *resty.Client
	cfg     *config.Config
	logger  *logger.Logger
}

func NewUsecase(repo auth.Repository, cfg *config.Config, logger *logger.Logger) auth.Usecase {
	return &usecase{
		tokens:  make(map[string]string),
		rClient: resty.New(),
		repo:    repo,
		cfg:     cfg,
		logger:  logger,
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

	u.tokens[string(refreshToken)] = accessString

	return &auth_models.SigninResponse{
		Access:  accessString,
		Refresh: encodedRefresh,
	}, nil
}

func (u *usecase) sendWarnIpMail(guid string) error {
	_, err := u.rClient.R().EnableTrace().Get(fmt.Sprintf("%s/user/ip_warning?guid=%s", u.cfg.MailerAddress, guid))
	return err
}

func (u *usecase) Refresh(params *auth_models.RefreshParams) (*auth_models.RefreshResponse, error) {
	refreshDecoded, err := base64.StdEncoding.DecodeString(params.Refresh)
	if err != nil {
		return nil, err
	}

	u.mu.Lock()
	accessString, ok := u.tokens[string(refreshDecoded)]
	u.mu.Unlock()

	if ok == false {
		return nil, fmt.Errorf("Tokens desync")
	}

	accessToken, err := jwt.Parse(accessString, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); ok == false {
			return nil, fmt.Errorf("Invalid token")
		}

		return []byte(u.cfg.Server.Key), nil
	})

	if err != nil {
		return nil, err
	}

	claims, ok := accessToken.Claims.(jwt.MapClaims)

	if ok == false {
		return nil, fmt.Errorf("Error while parse token")
	}

	ip, ok := claims["ip"]

	if ok == false {
		return nil, fmt.Errorf("Invalid token: ip not found")
	}

	guid, ok := claims["guid"]
	if ok == false {
		return nil, fmt.Errorf("Invalid token: guid not found")
	}

	refreshTokenHash, err := u.repo.GetRefreshTokenHash(guid.(string))
	if err != nil {
		return nil, err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(refreshTokenHash), refreshDecoded); err != nil {
		return nil, fmt.Errorf("Invalid refresh")

	}

	if ip.(string) != params.UserIp {
		if err := u.sendWarnIpMail(guid.(string)); err != nil {
			u.logger.Warnf("Error while call mailer: %s", err.Error())
		}
	}

	accessToken = jwt.NewWithClaims(jwt.SigningMethodHS512, jwt.MapClaims{
		"ip":   params.UserIp,
		"guid": guid.(string),
	})

	accessString, err = accessToken.SignedString([]byte(u.cfg.Server.Key))
	if err != nil {
		return nil, err
	}

	u.mu.Lock()
	u.tokens[params.Refresh] = accessString
	u.mu.Unlock()

	return &auth_models.RefreshResponse{
		Access: accessString,
	}, nil
}
