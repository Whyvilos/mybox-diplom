package repository

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/whyvilos/mybox"
)

type ChatPostgres struct {
	db *sqlx.DB
}

func NewChatPostgres(db *sqlx.DB) *ChatPostgres {
	return &ChatPostgres{db: db}
}

func (r *ChatPostgres) CreateChat(you_id, id_order int, status string) (int, error) {

	tx, err := r.db.Begin()
	if err != nil {
		fmt.Printf("r.db.Begin()\n")
		return 0, err
	}

	var id_user []int
	query := fmt.Sprintf("SELECT id_user FROM %s WHERE id_order=$1", usersOrdersTable)
	err = r.db.Select(&id_user, query, id_order)
	if err != nil || id_user[0] == 0 {
		tx.Rollback()
		return 0, err
	}

	var id_chat int
	query = fmt.Sprintf("INSERT INTO %s (status, type, id_order) VALUES ($1, $2, $3) RETURNING id_chat", chatsTable)
	row := tx.QueryRow(query, status, "order", id_order)
	if err := row.Scan(&id_chat); err != nil {
		tx.Rollback()
		fmt.Println(err)
		return 0, err
	}

	query = fmt.Sprintf("INSERT INTO %s (id_user, id_chat) VALUES ($1, $2)", usersChatsTable)
	_, err = tx.Exec(query, you_id, id_chat)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	query = fmt.Sprintf("INSERT INTO %s (id_user, id_chat) VALUES ($1, $2)", usersChatsTable)
	_, err = tx.Exec(query, id_user[0], id_chat)
	if err != nil {
		tx.Rollback()
		return 0, err
	}
	return 0, tx.Commit()
}

func (r *ChatPostgres) SendMassage(you_id, id_chat int, input mybox.Messaage) (int, error) {

	var id_message int
	query := fmt.Sprintf("INSERT INTO %s (id_chat, id_user, content, creation_time) VALUES ($1, $2, $3, $4) RETURNING id_message", messagesTable)
	row := r.db.QueryRow(query, id_chat, you_id, input.Content, "2022-05-23T20:04:31.475Z") //TODO добавить не вейковое время
	if err := row.Scan(&id_message); err != nil {
		return 0, err
	}

	return id_message, nil
}

func (r *ChatPostgres) CheckYouInChat(you_id, id_chat int) (bool, error) {
	flag_id := 0
	query := fmt.Sprintf("SELECT id_user FROM %s WHERE id_user=$1 AND id_chat=$2 LIMIT 1", usersChatsTable)
	err := r.db.Get(&flag_id, query, you_id, id_chat)
	if flag_id == 0 {
		return false, err
	}
	return true, err
}

func (r *ChatPostgres) GetUserInChat(id_chat int) ([]mybox.User, error) {

	var usersList []mybox.User
	var usersIds []int
	query := fmt.Sprintf("SELECT id_user FROM %s WHERE id_chat=$1", usersChatsTable)
	err := r.db.Select(&usersIds, query, id_chat)
	if err != nil {
		return usersList, err
	}
	for _, element := range usersIds {
		var user mybox.User
		query = fmt.Sprintf("SELECT id_user, name, username, url_avatar FROM %s WHERE id_user=$1", usersTable)
		err := r.db.Get(&user, query, element)
		if err != nil {
			return usersList, err
		}
		usersList = append(usersList, user)
	}
	return usersList, err
}

func (r *ChatPostgres) GetMessages(id_chat int) ([]mybox.Messaage, error) {
	var messagesList []mybox.Messaage
	query := fmt.Sprintf("SELECT * FROM %s WHERE id_chat=$1", messagesTable)
	err := r.db.Select(&messagesList, query, id_chat)
	if err != nil {
		return messagesList, err
	}
	return messagesList, err
}

func (r *ChatPostgres) FindChat(you_id, id_order int) (int, error) {
	var id_chat int
	query := fmt.Sprintf("SELECT ct.id_chat FROM %s ct INNER JOIN %s ot on ct.id_order=ot.id_order WHERE ot.id_order=$1 AND ot.id_user_owner=$2", chatsTable, ordersTable)
	err := r.db.Get(&id_chat, query, id_order, you_id)
	if err != nil {
		return 0, err
	}
	return id_chat, err
}

func (r *ChatPostgres) FindChat2(you_id, id_order int) (int, error) {
	var id_chat int
	query := fmt.Sprintf("SELECT ct.id_chat FROM %s ct INNER JOIN %s uot on ct.id_order=uot.id_order WHERE uot.id_order=$1 AND uot.id_user=$2", chatsTable, usersOrdersTable)
	err := r.db.Get(&id_chat, query, id_order, you_id)
	if err != nil {
		return 0, err
	}
	return id_chat, err
}

//SELECT ct.id_chat FROM chats ct INNER JOIN users_orders uot on ct.id_order=uot.id_order WHERE uot.id_order=2 AND uot.id_user=2;
