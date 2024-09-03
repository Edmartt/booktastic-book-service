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
	ID          string `json:"id" db:"id"`
	ISBN        string `json:"isbn" db:"isbn"`
	Title       string `json:"title" db:"title"`
	Pages       string `json:"pages" db:"pages"`
	CurrentPage string `json:"curren_page" db:"curren_page"`
	Author      string `json:"author" db:"author"`
	Year        string `json:"year" db:"year"`
	Status      string `json:"status" db:"status"`
}
