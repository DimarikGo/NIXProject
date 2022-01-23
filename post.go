package Rest_Api

type Post struct {
	Id     int    `json:"id" db:"id"`
	UserId int    `json:"UserId" db:"UserId"`
	Title  string `json:"title" db:"title"`
	Body   string `json:"body" db:"body"`
}
