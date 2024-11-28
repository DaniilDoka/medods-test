package pg

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type PgCredentials struct {
	Database string `env:"DB_NAME"`
	Address  string `env:"DB_ADDRESS"` //ex localhost:5432
	User     string `env:"DB_USER"`
	Password string `env:"DB_PASSWORD"`
	SslMode  string `env:"DB_SSLMODE"`
}

type Pg struct {
	db *sqlx.DB
}

func Open(creds *PgCredentials) (*Pg, error) {
	connUrl := fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=%s", creds.User, creds.Password, creds.Address, creds.Database, creds.SslMode)
	db, err := sqlx.Connect("postgres", connUrl)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return &Pg{
		db: db,
	}, nil
}

func (d Pg) Exec(query string, args ...interface{}) error {
	_, err := d.db.Exec(query, args...)
	return err
}

func (d Pg) Get(query string, dest interface{}, args ...interface{}) error {
	err := d.db.Get(dest, query, args...)
	return err
}

func (d Pg) Select(query string, dest interface{}, args ...interface{}) error {
	err := d.db.Get(dest, query, args...)
	return err
}
