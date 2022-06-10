package mybox

type Notice struct {
	Id_notice int    `json:"id_notice" db:"id_notice"`
	Content   string `json:"content" db:"content"`
	Status    string `json:"status" db:"status"`
}
