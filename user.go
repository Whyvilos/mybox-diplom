package mybox

type User struct {
	Id_user  int     `json:"-" db:"id_user"`
	Mail     string  `json:"mail" binding:"required"`
	Name     string  `json:"name" binding:"required"`
	Username string  `json:"username" binding:"required"`
	Password string  `json:"password" binding:"required"`
	Phone    *string `json:"phone" db:"phone_number"`
	Rank     bool    `json:"rank" `
}
