package service

import (
	"github.com/Munovv/go-url-crop/pkg/model"
	"github.com/Munovv/go-url-crop/pkg/repository"
)

type Link interface {
	CropLink(link string) (string, error)
	GetLink(url string) (model.Link, error)
	RedirectLink(code string) (string, error)
}

type Service struct {
	Link
}

func NewService(r *repository.Repository) *Service {
	return &Service{
		Link: NewLinkService(r.Link),
	}
}
