package repository

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

const (
	usersTable        = "users"
	itemsTable        = "items"
	usersItemsTable   = "users_items"
	postsTable        = "posts"
	usersPostsTable   = "users_posts"
	followersTable    = "followers"
	favoriteTable     = "favorite"
	ordersTable       = "orders"
	noticesTable      = "notifications"
	usersOrdersTable  = "users_orders"
	usersNoticesTable = "users_notifications"
	chatsTable        = "chats"
	usersChatsTable   = "users_chats"
	messagesTable     = "messages"
)

type Config struct {
	Host     string
	Port     string
	Username string
	Password string
	DBName   string
	SSLMode  string
}

func NewPostgresDB(cfg Config) (*sqlx.DB, error) {
	db, err := sqlx.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.Username, cfg.DBName, cfg.Password, cfg.SSLMode))
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	return db, nil
}
