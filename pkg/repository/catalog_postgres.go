package repository

import (
	"errors"
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

	query := fmt.Sprintf("SELECT it.id_item, it.title, it.url_media, it.description, it.status, it.price FROM %s it INNER JOIN %s uit on it.id_item = uit.id_item WHERE uit.id_user = $1 ORDER BY it.id_item DESC",

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

	createItemQuery := fmt.Sprintf("INSERT INTO %s (title, url_media, description, status, price) VALUES ($1, $2, $3, $4, $5) RETURNING id_item", itemsTable)
	row := tx.QueryRow(createItemQuery, input.Title, input.Media, input.Description, input.Status, input.Price)
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

func (r *CatalogPostgres) AddFavorite(userId, itemId int) error {
	query := fmt.Sprintf("INSERT INTO %s (id_user, id_item) VALUES($1,$2)",
		favoriteTable)
	_, err := r.db.Exec(query, userId, itemId)
	return err
}

func (r *CatalogPostgres) DeleteFavorite(userId, itemId int) error {
	check, _ := r.CheckFavorite(userId, itemId)
	if check {
		query := fmt.Sprintf("DELETE FROM %s WHERE id_user=$1 AND id_item=$2", favoriteTable)
		_, err := r.db.Exec(query, userId, itemId)
		return err
	}
	return errors.New("it isn't your favorite")
}

func (r *CatalogPostgres) CheckFavorite(userId, itemId int) (bool, error) {
	id_item := 0
	query := fmt.Sprintf("SELECT id_item FROM %s WHERE id_user=$1 AND id_item=$2",
		favoriteTable)
	err := r.db.Get(&id_item, query, userId, itemId)
	if id_item == 0 {
		return false, err
	}
	return true, err
}

func (r *CatalogPostgres) GetById(id_item int) (mybox.Item, error) {
	var item mybox.Item
	query := fmt.Sprintf("SELECT it.id_item, it.title, it.description, it.url_media, it.status,  it.count, it.price, uit.id_user FROM %s it INNER JOIN %s uit on it.id_item=uit.id_item WHERE it.id_item=$1", itemsTable, usersItemsTable)
	err := r.db.Get(&item, query, id_item)
	return item, err
}
