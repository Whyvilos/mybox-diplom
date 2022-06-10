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
	Follow(you_id, id_user int) error
	UnFollow(you_id, id_user int) error
	CheckFollow(you_id, id_user int) (bool, error)
	LoadLine(you_id int) ([]mybox.Post, error)
	LoadFavorite(you_id int) ([]mybox.Item, error)
	GetNotices(you_id int) ([]mybox.Notice, error)
	NoticeCheck(you_id int) error
}

type Catalog interface {
	GetAll(id_catalog int) ([]mybox.SimpleItem, error)
	CreateItem(userId int, input mybox.Item) (int, error)
	AddFavorite(userId, itemId int) error
	CheckFavorite(userId, itemId int) (bool, error)
	GetById(id_item int) (mybox.Item, error)
	DeleteFavorite(userId, itemId int) error
}

type Feed interface {
	CreatePost(id_user int, post mybox.Post) (int, error)
	GetAll(id_feed int) ([]mybox.Post, error)
}

type Order interface {
	CreateOrder(you_id int, input mybox.Order) (int, error)
	GetOrders(you_id int) ([]mybox.OrderResponce, error)
	GetOrdersForYou(you_id int) ([]mybox.OrderResponce, error)
	UpdateOrderStatus(you_id, id_order int, status string) error
}

type Media interface {
	SaveUrlAvatar(userId int, path string) error
}

type Chat interface {
	CreateChat(you_id, id_order int, status string) (int, error)
	SendMassage(you_id, id_chat int, input mybox.Messaage) (int, error)
	CheckYouInChat(you_id, id_chat int) (bool, error)
	GetUserInChat(id_chat int) ([]mybox.User, error)
	GetMessages(id_chat int) ([]mybox.Messaage, error)
	FindChat(you_id, id_order int) (int, error)
	FindChat2(you_id, id_order int) (int, error)
}

type Repository struct {
	Authorization
	UserProfile
	Catalog
	Feed
	Media
	Order
	Chat
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
		UserProfile:   NewProfilePostgres(db),
		Catalog:       NewCatalogPostgres(db),
		Feed:          NewFeedPostgres(db),
		Media:         NewMediaPostgres(db),
		Order:         NewOrderPostgres(db),
		Chat:          NewChatPostgres(db),
	}
}
