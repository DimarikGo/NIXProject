package repository

import (
	"Rest-Api"
	"database/sql"
	"fmt"
)

const (
	n        = "name"
	c_id     = "id"
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

func (c *CommentMysql) AddC(comment Rest_Api.Comment) (i int, err error) {
	queryPost := fmt.Sprintf("INSERT INTO `%s` (`%s`,`%s`,`%s`,`%s`,`%s`) VALUES('%s','%v', '%s','%s','%v')", tableCom, n, c_id, e, bc, pId, comment.Name, comment.Id, comment.Email, comment.Body, comment.PostId)
	_ = c.db.QueryRow(queryPost)

	return comment.PostId, nil
}
