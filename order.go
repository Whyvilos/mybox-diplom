package mybox

type Order struct {
	Id_order      int    `json:"id_order" db:"id_order"`
	Id_user_owner int    `json:"id_user_owner" db:"id_user_owner"`
	Id_item       int    `json:"id_item" db:"id_item"`
	Status        string `json:"status" db:"status"`
	Description   string `json:"description" db:"description"`
}

type OrderResponce struct {
	Id_order      int     `json:"id_order" db:"id_order"`
	Id_user_owner int     `json:"id_user_owner" db:"id_user_owner"`
	Id_item       int     `json:"id_item" db:"id_item"`
	Status        string  `json:"status" db:"status"`
	Description   string  `json:"description" db:"description"`
	Title         string  `json:"title" db:"title"`
	Url_media     *string `json:"url_media" db:"url_media"`
	Price         *int    `json:"price" db:"price"`
}
