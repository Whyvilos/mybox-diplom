package mybox

type Messaage struct {
	Id_message    int     `json:"id_message" db:"id_message"`
	Content       string  `json:"content" db:"content"`
	Id_chat       int     `json:"id_chat" db:"id_chat"`
	Id_user       int     `json:"id_user" db:"id_user"`
	Creation_time *string `json:"creation_time" db:"creation_time"`
}

type AllMessages struct {
	Massages []Messaage `json:"messages"`
	Users    []User     `json:"users"`
}
