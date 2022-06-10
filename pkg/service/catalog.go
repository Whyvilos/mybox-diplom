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

func (s *CatalogService) AddFavorite(userId, itemId int) error {
	return s.repo.AddFavorite(userId, itemId)
}

func (s *CatalogService) CheckFavorite(userId, itemId int) (bool, error) {
	return s.repo.CheckFavorite(userId, itemId)
}

func (s *CatalogService) GetById(id_item int) (mybox.Item, error) {
	return s.repo.GetById(id_item)
}

func (s *CatalogService) DeleteFavorite(userId, itemId int) error {
	return s.repo.DeleteFavorite(userId, itemId)
}
