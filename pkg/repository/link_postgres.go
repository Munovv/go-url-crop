package repository

import (
	"github.com/Munovv/go-url-crop/pkg/model"
	"github.com/jmoiron/sqlx"
)

type LinkPostgres struct {
	db *sqlx.DB
}

func NewLinkPostgres(db *sqlx.DB) *LinkPostgres {
	return &LinkPostgres{db: db}
}

func (r *LinkPostgres) CropLink(link string) (string, error) {
	return "link.ru/test", nil
}

func (r *LinkPostgres) GetLink(link string) (model.Link, error) {
	return model.Link{}, nil
}
