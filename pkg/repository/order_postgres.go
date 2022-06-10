package repository

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/whyvilos/mybox"
)

type OrderPostgres struct {
	db *sqlx.DB
}

func NewOrderPostgres(db *sqlx.DB) *OrderPostgres {
	return &OrderPostgres{db: db}
}

func (r *OrderPostgres) CreateOrder(you_id int, input mybox.Order) (int, error) {

	tx, err := r.db.Begin()
	if err != nil {
		fmt.Printf("r.db.Begin()\n")
		return 0, err
	}

	var id_order int

	createOrderQuery := fmt.Sprintf("INSERT INTO %s (id_user_owner, id_item, status, description) VALUES ($1, $2, $3, $4) RETURNING id_order", ordersTable)
	row := tx.QueryRow(createOrderQuery, input.Id_user_owner, input.Id_item, "new", input.Description)
	if err := row.Scan(&id_order); err != nil {
		tx.Rollback()
		return 0, err
	}

	createUsersOrdersQuery := fmt.Sprintf("INSERT INTO %s (id_user, id_order) VALUES ($1, $2)", usersOrdersTable)
	_, err = tx.Exec(createUsersOrdersQuery, you_id, id_order)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	var id_notice int

	createNoticeQuery := fmt.Sprintf("INSERT INTO %s (content, status) VALUES ($1, $2) RETURNING id_notice", noticesTable)
	row = tx.QueryRow(createNoticeQuery, fmt.Sprintf("Получен новый заказ %d", id_order), "new")
	if err := row.Scan(&id_notice); err != nil {
		tx.Rollback()
		return 0, err
	}

	createUsersNoticesQuery := fmt.Sprintf("INSERT INTO %s (id_user, id_notice) VALUES ($1, $2)", usersNoticesTable)
	_, err = tx.Exec(createUsersNoticesQuery, input.Id_user_owner, id_notice)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	return id_order, tx.Commit()

}

func (r *OrderPostgres) GetOrders(you_id int) ([]mybox.OrderResponce, error) {
	var list []mybox.OrderResponce
	query := fmt.Sprintf("SELECT ot.id_order, ot.id_user_owner, ot.id_item, ot.status, ot.description, it.title, it.url_media, it.price FROM %s ot INNER JOIN %s it on it.id_item=ot.id_item INNER JOIN %s uot on uot.id_order=ot.id_order WHERE uot.id_user=$1", ordersTable, itemsTable, usersOrdersTable)
	err := r.db.Select(&list, query, you_id)
	return list, err
}

func (r *OrderPostgres) GetOrderOffer(you_id int) ([]mybox.OrderResponce, error) {
	var list []mybox.OrderResponce
	query := fmt.Sprintf("SELECT ot.id_order, ot.id_user_owner, ot.id_item, ot.status, ot.description, it.title, it.url_media, it.price FROM %s ot INNER JOIN %s it on it.id_item=ot.id_item INNER JOIN %s uot on uot.id_order=ot.id_order WHERE uot.id_user=$1", ordersTable, itemsTable, usersOrdersTable)
	err := r.db.Select(&list, query, you_id)
	return list, err
}

func (r *OrderPostgres) GetOrdersForYou(you_id int) ([]mybox.OrderResponce, error) {
	var list []mybox.OrderResponce
	query := fmt.Sprintf("SELECT ot.id_order, ot.id_user_owner, ot.id_item, ot.status, ot.description, it.title, it.url_media, it.price FROM %s ot INNER JOIN %s it on it.id_item=ot.id_item INNER JOIN %s uot on uot.id_order=ot.id_order WHERE ot.id_user_owner=$1", ordersTable, itemsTable, usersOrdersTable)
	err := r.db.Select(&list, query, you_id)
	return list, err
}

func (r *OrderPostgres) UpdateOrderStatus(you_id, id_order int, status string) error {
	var id int
	query := fmt.Sprintf("UPDATE %s SET status=$1 WHERE id_user_owner=$2 AND id_order=$3 RETURNING id_order", ordersTable)
	a := r.db.QueryRow(query, status, you_id, id_order)
	if err := a.Scan(&id); err != nil {
		return err
	}
	return nil
}
