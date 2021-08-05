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
}

func NewMysqlDb(cfg Config) (*sqlx.DB, error) {
	db, err := sqlx.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		cfg.Username, cfg.Password, cfg.Host, cfg.Port, cfg.Db))
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}
