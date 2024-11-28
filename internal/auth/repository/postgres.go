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
	var result bool
	err := p.db.Get(`select exists(select * from token where refresh_token=$1)`, &result, refresh)
	return result, err
}

func (p *postgres) GetRefreshTokenHash(guid string) (string, error) {
	var result string
	err := p.db.Get(`select refresh_token from token where user_id=$1`, &result, guid)
	return result, err
}

func (p *postgres) PutRefreshToken(guid string, refresh string, exp time.Time) error {
	err := p.db.Exec(`insert into token(user_id, refresh_token, exp) values ($1, $2, $3)`, guid, refresh, exp)
	return err
}
