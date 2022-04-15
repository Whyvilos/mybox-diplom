package service

import (
	"github.com/whyvilos/mybox"
	"github.com/whyvilos/mybox/pkg/repository"
)

type Authorization interface {
	CreateUser(user mybox.User) (int, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(token string) (int, error)
}

type UserProfile interface {
	GetById(you_id, id_user int) (mybox.User, error)
}

type Catalog interface {
	GetAll(id_catalog int) ([]mybox.SimpleItem, error)
	CreateItem(userId int, input mybox.Item) (int, error)
}

type Feed interface {
	CreatePost(id_user int, post mybox.Post) (int, error)
	GetAll(id_feed int) ([]mybox.Post, error)
}

type Service struct {
	Authorization
	UserProfile
	Catalog
	Feed
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repo.Authorization),
		UserProfile:   NewUserProfileService(repo.UserProfile),
		Catalog:       NewCatalogService(repo.Catalog),
		Feed:          NewFeedService(repo.Feed),
	}
}
