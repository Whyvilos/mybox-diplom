package repository

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/whyvilos/mybox"
)

type CatalogPostgres struct {
	db *sqlx.DB
}

func NewCatalogPostgres(db *sqlx.DB) *CatalogPostgres {
	return &CatalogPostgres{db: db}
}

func (r *CatalogPostgres) GetAll(id_catalog int) ([]mybox.SimpleItem, error) {

	var simpleItems []mybox.SimpleItem

	query := fmt.Sprintf("SELECT it.id_item, it.title, it.status, it.price FROM %s it INNER JOIN %s uit on it.id_item = uit.id_item WHERE uit.id_user = $1",
		itemsTable, usersItemsTable)
	err := r.db.Select(&simpleItems, query, id_catalog)

	return simpleItems, err

}

func (r *CatalogPostgres) CreateItem(userId int, input mybox.Item) (int, error) {
	tx, err := r.db.Begin()
	if err != nil {
		fmt.Printf("r.db.Begin()\n")
		return 0, err
	}

	var id_item int

	createItemQuery := fmt.Sprintf("INSERT INTO %s (title, description, status, price) VALUES ($1, $2, $3, $4) RETURNING id_item", itemsTable)
	row := tx.QueryRow(createItemQuery, input.Title, input.Description, input.Status, input.Price)
	if err := row.Scan(&id_item); err != nil {
		tx.Rollback()
		return 0, err
	}

	createUsersListQuery := fmt.Sprintf("INSERT INTO %s (id_user, id_item) VALUES ($1, $2)", usersItemsTable)
	_, err = tx.Exec(createUsersListQuery, userId, id_item)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	return id_item, tx.Commit()
}
