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

	createItemQuery := fmt.Sprintf("INSERT INTO %s (description, url_media, creation_time) VALUES ($1, $2, $3) RETURNING id_post", postsTable)
	row := tx.QueryRow(createItemQuery, post.Description, post.Media, post.Creation_time)
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

	query := fmt.Sprintf("SELECT pt.id_post, pt.description, pt.url_media, pt.creation_time, upt.id_user, ut.username FROM %s pt INNER JOIN %s upt on pt.id_post = upt.id_post INNER JOIN %s ut on ut.id_user=upt.id_user WHERE upt.id_user = $1 ORDER BY pt.id_post DESC",
		postsTable, usersPostsTable, usersTable)
	err := r.db.Select(&posts, query, id_feed)

	return posts, err
}
