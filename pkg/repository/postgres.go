package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
)

const (
	linkTable = "links"
)

type Config struct {
	Host     string
	Port     string
	Username string
	Password string
	Db       string
	SslMode  string
}

func NewPostgresDb(cfg Config) (*sqlx.DB, error) {
	db, err := sqlx.Open("postgres", fmt.Sprintf(`host=%s port=%s user=%s dbname=%s password=%s sslmode=%s`,
		cfg.Host, cfg.Port, cfg.Username, cfg.Password, cfg.Db, cfg.SslMode))
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}
