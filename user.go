package mybox

type User struct {
	Id_user  int     `json:"id_user" db:"id_user"`
	Mail     string  `json:"mail" db:"mail"`
	Name     string  `json:"name" db:"name"`
	Username string  `json:"username" db:"username"`
	Password string  `json:"password" db:"password"`
	Avatar   string  `json:"url_avatar" db:"url_avatar"`
	Phone    *string `json:"phone" db:"phone_number"`
	Rank     bool    `json:"rank" `
	IsYou    bool    `json:"isyou"`
}

type UserId struct {
	Id_user int `json:"id_user"`
}
