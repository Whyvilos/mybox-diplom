package mybox

type Item struct {
	Id_item     int    `json:"-" db:"id_item"`
	Title       string `json:"title" binding:"required"`
	Description string `json:"description" binding:"required"`
	Status      string `json:"status" binding:"required"`
	Count       *int   `json:"count"`
	Price       int    `json:"price" binding:"required"`
}

type SimpleItem struct {
	Id_item int    `json:"-" db:"id_item"`
	Title   string `json:"title" binding:"required"`
	Status  string `json:"status" binding:"required"`
	Price   *int   `json:"price"`
}
