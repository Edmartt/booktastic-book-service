package models

// I want to know if I'm reading, I read or if I need to buy that book, Status only allowed answers
const (
	Read = iota
	Reading
	ToBeRead
	PendingPurchase
)

// Books model
type Books struct {
	UUID        string `db:"uuid"`
	ISBN        string `json:"isbn" db:"isbn"`
	Title       string `json:"title" db:"title"`
	Pages       string `json:"pages" db:"pages"`
	CurrentPage string `json:"current_page" db:"current_page"`
	Author      string `json:"author" db:"author"`
	Year        string `json:"year" db:"year"`
	Status      string `json:"status" db:"status"`
}
