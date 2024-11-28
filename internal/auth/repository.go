package auth

import (
	"time"
)

type Repository interface {
	GetRefreshTokenHash(guid string) (string, error)
	PutRefreshToken(guid string, refresh string, exp time.Time) error
}
