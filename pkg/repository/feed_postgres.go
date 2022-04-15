package repository

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/whyvilos/mybox"
)

type FeedPostgres struct {
	db *sqlx.DB
}

func NewFeedPostgres(db *sqlx.DB) *FeedPostgres {
	return &FeedPostgres{db: db}
}

func (r *FeedPostgres) CreatePost(id_user int, post mybox.Post) (int, error) {

	tx, err := r.db.Begin()
	if err != nil {
		fmt.Printf("r.db.Begin()\n")
		return 0, err
	}

	var id_post int

	createItemQuery := fmt.Sprintf("INSERT INTO %s (description, creation_time) VALUES ($1, $2) RETURNING id_post", postsTable)
	row := tx.QueryRow(createItemQuery, post.Description, post.Creation_time)
	if err := row.Scan(&id_post); err != nil {
		tx.Rollback()
		return 0, err
	}

	createUsersListQuery := fmt.Sprintf("INSERT INTO %s (id_user, id_post) VALUES ($1, $2)", usersPostsTable)
	_, err = tx.Exec(createUsersListQuery, id_user, id_post)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	return id_post, tx.Commit()
}

func (r *FeedPostgres) GetAll(id_feed int) ([]mybox.Post, error) {

	var posts []mybox.Post

	query := fmt.Sprintf("SELECT pt.id_post, pt.description, pt.creation_time FROM %s pt INNER JOIN %s upt on pt.id_post = upt.id_post WHERE upt.id_user = $1",
		postsTable, usersPostsTable)
	err := r.db.Select(&posts, query, id_feed)

	return posts, err
}
