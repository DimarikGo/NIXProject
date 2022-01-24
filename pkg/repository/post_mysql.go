package repository

import (
	"Rest-Api"
	"database/sql"
	"fmt"
)

const (
	uid       = "UserId"
	id_id     = "id"
	T         = "title"
	B         = "body"
	tablePost = "posts"
)

type PostMysql struct {
	db *sql.DB
}

func NewPostMysql(db *sql.DB) *PostMysql {
	return &PostMysql{db: db}
}

func (p *PostMysql) AddP(post Rest_Api.Post) (int, error) {

	queryPost := fmt.Sprintf("INSERT INTO `%s` (`%s`,`%s`,`%s`,`%s`) VALUES('%v',%v, '%s','%s')", tablePost, uid, id_id, T, B, post.UserId, post.Id, post.Title, post.Body)
	_ = p.db.QueryRow(queryPost)
	

	return post.Id, nil
}


