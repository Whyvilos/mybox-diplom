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

type Media interface {
	SaveUrlAvatar(userId int, path string) error
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
	DeleteFavorite(userId, itemId int) error
	CheckFavorite(userId, itemId int) (bool, error)
	GetById(id_item int) (mybox.Item, error)
}

type Feed interface {
	CreatePost(id_user int, post mybox.Post) (int, error)
	GetAll(id_feed int) ([]mybox.Post, error)
}

type Chat interface {
	CreateChat(you_id, id_order int, status string) (int, error)
	SendMassage(you_id, id_chat int, input mybox.Messaage) (int, error)
	GetAllMessage(you_id, id_chat int) (mybox.AllMessages, error)
	FindChat(you_id, id_order int) (int, error)
	FindChat2(you_id, id_order int) (int, error)
}

type Order interface {
	CreateOrder(you_id int, input mybox.Order) (int, error)
	GetOrders(you_id int) ([]mybox.OrderResponce, error)
	GetOrdersForYou(you_id int) ([]mybox.OrderResponce, error)
	UpdateOrderStatus(you_id, id_order int, status string) error
}

type Service struct {
	Authorization
	UserProfile
	Catalog
	Feed
	Media
	Order
	Chat
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repo.Authorization),
		UserProfile:   NewUserProfileService(repo.UserProfile),
		Catalog:       NewCatalogService(repo.Catalog),
		Feed:          NewFeedService(repo.Feed),
		Media:         NewMediaService(repo.Media),
		Order:         NewOrderService(repo.Order),
		Chat:          NewCharService(repo.Chat),
	}
}
