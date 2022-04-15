package repository

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/whyvilos/mybox"
)

type AuthPostgres struct {
	db *sqlx.DB
}

func NewAuthPostgres(db *sqlx.DB) *AuthPostgres {
	return &AuthPostgres{db: db}
}

func (r *AuthPostgres) CreateUser(user mybox.User) (int, error) {

	var id int

	query := fmt.Sprintf("INSERT INTO %s (mail,name,username,password_hash,phone_number, rank) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id_user", usersTable)

	row := r.db.QueryRow(query, user.Mail, user.Name, user.Username, user.Password, user.Phone, true)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (r *AuthPostgres) GetUser(username, password string) (mybox.User, error) {
	var user mybox.User

	query := fmt.Sprintf("SELECT id_user FROM %s WHERE username=$1 AND password_hash=$2", usersTable)

	err := r.db.Get(&user, query, username, password)

	return user, err
}
