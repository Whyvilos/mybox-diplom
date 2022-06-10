package repository

import (
	"errors"
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

	query := fmt.Sprintf("SELECT id_user, mail, name, username, url_avatar, phone_number, rank FROM %s WHERE id_user=$1", usersTable)
	err := r.db.Get(&user, query, id_user)

	return user, err
}

func (r *ProfilePostgres) Follow(you_id, id_user int) error {

	check, err := r.CheckFollow(you_id, id_user)
	if !check {
		query := fmt.Sprintf("INSERT INTO %s (id_user, id_user_follower) VALUES ($1, $2)", followersTable)
		_, err = r.db.Exec(query, id_user, you_id)
		return err
	}
	return errors.New("you already follow")
}

func (r *ProfilePostgres) UnFollow(you_id, id_user int) error {

	check, err := r.CheckFollow(you_id, id_user)
	if check {
		query := fmt.Sprintf("DELETE FROM %s WHERE id_user=$1 AND id_user_follower=$2", followersTable)
		_, err = r.db.Exec(query, id_user, you_id)
		return err
	}
	return errors.New("you are not following")
}

func (r *ProfilePostgres) CheckFollow(you_id, id_user int) (bool, error) {
	flag_id := 0
	query := fmt.Sprintf("SELECT id_user FROM %s WHERE id_user=$1 AND id_user_follower=$2 LIMIT 1", followersTable)
	err := r.db.Get(&flag_id, query, id_user, you_id)
	if flag_id == 0 {
		return false, err
	}
	return true, err
}

func (r *ProfilePostgres) LoadLine(you_id int) ([]mybox.Post, error) {
	var line []mybox.Post
	query := fmt.Sprintf("SELECT pt.id_post,pt.url_media, pt.description, pt.creation_time, pt.id_item, pt.price, upt.id_user, ut.username FROM %s pt INNER JOIN %s upt on pt.id_post=upt.id_post INNER JOIN %s ut on ut.id_user=upt.id_user ORDER BY pt.id_post DESC", postsTable, usersPostsTable, usersTable)
	err := r.db.Select(&line, query)
	return line, err
}

func (r *ProfilePostgres) LoadFavorite(you_id int) ([]mybox.Item, error) {
	var list []mybox.Item
	query := fmt.Sprintf("SELECT it.id_item, it.title, it.url_media, it.description, it.status, it.price, uit.id_user FROM %s it INNER JOIN %s uit on uit.id_item=it.id_item INNER JOIN %s ft on it.id_item=ft.id_item WHERE ft.id_user=$1", itemsTable, usersItemsTable, favoriteTable)
	err := r.db.Select(&list, query, you_id)
	return list, err
}

func (r *ProfilePostgres) GetNotices(you_id int) ([]mybox.Notice, error) {
	var list []mybox.Notice
	query := fmt.Sprintf("SELECT nt.id_notice, nt.content, nt.status FROM %s nt INNER JOIN %s unt on unt.id_notice=nt.id_notice WHERE unt.id_user=$1 ORDER BY nt.id_notice DESC", noticesTable, usersNoticesTable)
	err := r.db.Select(&list, query, you_id)
	return list, err
}

func (r *ProfilePostgres) NoticeCheck(you_id int) error {
	query := fmt.Sprintf("UPDATE %s SET status='check' FROM %s t1 INNER JOIN %s unt ON unt.id_notice=t1.id_notice WHERE unt.id_user=$1;", noticesTable, noticesTable, usersNoticesTable)
	_, err := r.db.Exec(query, you_id)
	return err
}
