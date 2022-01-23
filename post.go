package Rest_Api

type Post struct {
	UserId int    `json:"userId" db:"UserId"`
	Id     int    `json:"id" db:"id"`
	Title  string `json:"title" db:"title"`
	Body   string `json:"body" db:"body"`
}
