package mybox

import "time"

type Post struct {
	Id_post       int       `json:"-" db:"id_post"`
	Description   string    `json:"description" binding:"required"`
	Media         *string   `json:"url_media" db:"url_media"`
	Creation_time time.Time `json:"creation_time" binding:"required"`
	Id_item       *int      `json:"id_item"`
	Price         *int      `json:"price"`
	Id_user       int       `json:"id_user" db:"id_user"`
	Username      *string   `json:"username" db:"username"`
}
