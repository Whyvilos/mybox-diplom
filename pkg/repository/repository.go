package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/whyvilos/mybox"
)

type Authorization interface {
	CreateUser(user mybox.User) (int, error)
	GetUser(username, password string) (mybox.User, error)
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

type Repository struct {
	Authorization
	UserProfile
	Catalog
	Feed
}

//db *sqlx.DB
func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
		UserProfile:   NewProfilePostgres(db),
		Catalog:       NewCatalogPostgres(db),
		Feed:          NewFeedPostgres(db),
	}
}
