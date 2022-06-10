package repository

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

type MediaPostgres struct {
	db *sqlx.DB
}

func NewMediaPostgres(db *sqlx.DB) *MediaPostgres {
	return &MediaPostgres{db: db}
}

func (r *MediaPostgres) SaveUrlAvatar(userId int, path string) error {

	query := fmt.Sprintf("UPDATE %s SET url_avatar=$1 WHERE id_user=$2", usersTable)
	_, err := r.db.Exec(query, path, userId)
	return err

}
