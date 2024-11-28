package auth

import "time"

type Repository interface {
	RefreshTokenExists(refresh string) (bool, error)
	PutRefreshToken(guid string, refresh string, exp time.Time) error
	GetUserIp(guid string) (string, error)
}
