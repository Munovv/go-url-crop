package repository

import (
	"github.com/Munovv/go-url-crop/pkg/model"
	"github.com/jmoiron/sqlx"
)

type Link interface {
	CropLink(link string) (string, error)
	GetLink(link string) (model.Link, error)
	GenerateCode() string
	RedirectLink(code string) (string, error)
}

type Repository struct {
	Link
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Link: NewLinkMysql(db),
	}
}
