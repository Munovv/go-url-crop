package service

import (
	"github.com/Munovv/go-url-crop/pkg/model"
	"github.com/Munovv/go-url-crop/pkg/repository"
)

type LinkService struct {
	repo repository.Link
}

func NewLinkService(r repository.Link) *LinkService {
	return &LinkService{repo: r}
}

func (s *LinkService) CropLink(link string) (string, error) {
	return s.repo.CropLink(link)
}

func (s *LinkService) GenerateCode() (string, error) {
	return "test", nil
}

func (s *LinkService) GetLink(url string) (model.Link, error) {
	link, err := s.repo.GetLink(url)
	if err != nil {
		return model.Link{}, err
	}

	return link, nil
}
