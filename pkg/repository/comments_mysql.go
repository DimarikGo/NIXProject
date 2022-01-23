package repository

import (
	"Rest-Api"
	"database/sql"
	"fmt"
)

const (
	n        = "name"
	e        = "email"
	bc       = "body"
	pId      = "post_id"
	tableCom = "comments"
	//	p_id     = "post_id"
)

type CommentMysql struct {
	db *sql.DB
}

func NewCommentMysql(db *sql.DB) *CommentMysql {
	return &CommentMysql{db: db}
}

func (c *CommentMysql) AddC(postid int, comment Rest_Api.Comment) (i int, err error) {
	queryPost := fmt.Sprintf("INSERT INTO `%s` (`%s`,`%s`,`%s`,`%s`) VALUES('%v', '%s','%s','%v')", tableCom, n, e, bc, pId, comment.Name, comment.Email, comment.Body, postid)
	_ = c.db.QueryRow(queryPost)

	return postid, nil
}
