package auth_repository

import (
	"medods-test/internal/auth"
	"medods-test/pkg/pg"
	"time"
)

type postgres struct {
	db *pg.Pg
}

func NewRepository(pgConn *pg.Pg) auth.Repository {
	return &postgres{
		db: pgConn,
	}
}

func (p *postgres) RefreshTokenExists(refresh string) (bool, error) {
	return false, nil
}

func (p *postgres) PutRefreshToken(guid string, refresh string, exp time.Time) error {
	err := p.db.Exec(`insert into token(user_id, refresh_token, exp) values ($1, $2, $3)`, guid, refresh, exp)
	return err
}

func (p *postgres) GetUserIp(guid string) (string, error) {
	return "", nil
}
