package repository

import (
	"Rest-Api"
	"database/sql"
)

type Post interface {
	AddP(post Rest_Api.Post) (int, error)
}

type Comments interface {
	AddC(postid int, comment Rest_Api.Comment) (int, error)
}

type Repository struct {
	Post
	Comments
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		Post:     NewPostMysql(db),
		Comments: NewCommentMysql(db)}
}
