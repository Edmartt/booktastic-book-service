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
	ID     string `json:"id"`
	ISBN   string `json:"isbn"`
	Title  string `json:"title"`
	Author string `json:"author"`
	Year   string `json:"year"`
	Status string `json:"status"`
}
