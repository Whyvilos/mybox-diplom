package service

import (
	"github.com/whyvilos/mybox"
	"github.com/whyvilos/mybox/pkg/repository"
)

type CatalogService struct {
	repo repository.Catalog
}

func NewCatalogService(repo repository.Catalog) *CatalogService {
	return &CatalogService{repo: repo}
}

func (s *CatalogService) GetAll(id_catalog int) ([]mybox.SimpleItem, error) {
	return s.repo.GetAll(id_catalog)
}

func (s *CatalogService) CreateItem(userId int, input mybox.Item) (int, error) {
	return s.repo.CreateItem(userId, input)
}
