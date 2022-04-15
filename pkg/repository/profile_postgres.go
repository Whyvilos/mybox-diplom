package repository

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/whyvilos/mybox"
)

type ProfilePostgres struct {
	db *sqlx.DB
}

func NewProfilePostgres(db *sqlx.DB) *ProfilePostgres {
	return &ProfilePostgres{db: db}
}

func (r *ProfilePostgres) GetById(you_id, id_user int) (mybox.User, error) {

	var user mybox.User

	query := fmt.Sprintf("SELECT id_user, mail, name, username, phone_number, rank FROM %s WHERE id_user=$1", usersTable)
	err := r.db.Get(&user, query, id_user)

	return user, err
}
