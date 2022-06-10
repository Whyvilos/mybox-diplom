package mybox

type Item struct {
	Id_item     int     `json:"id_item" db:"id_item"`
	Id_user     *int    `json:"id_user" db:"id_user"`
	Title       string  `json:"title" binding:"required"`
	Media       *string `json:"url_media" db:"url_media"`
	Description string  `json:"description" binding:"required"`
	Status      string  `json:"status" binding:"required"`
	Count       *int    `json:"count"`
	Price       int     `json:"price" binding:"required"`
}

type SimpleItem struct {
	Id_item     int     `json:"id_item" db:"id_item"`
	Id_user     *int    `json:"id_user" db:"id_user"`
	Title       string  `json:"title" binding:"required"`
	Media       *string `json:"url_media" db:"url_media"`
	Description string  `json:"description" binding:"required"`
	Status      string  `json:"status" binding:"required"`
	Price       *int    `json:"price"`
}
