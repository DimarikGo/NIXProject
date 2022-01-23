package repository

import (
	"Rest-Api"
	"database/sql"
	"fmt"
)

const (
	uid          = "UserId"
	T            = "title"
	B            = "body"
	tablePost    = "posts"
	tablePostCom = "posts_comments"
	p_id         = "post_id"
)

type PostMysql struct {
	db *sql.DB
}

func NewPostMysql(db *sql.DB) *PostMysql {
	return &PostMysql{db: db}
}

func (p *PostMysql) AddP(post Rest_Api.Post) (int, error) {
	var postId int
	//tx, _ := p.db.Begin()
	//if err != nil {
	//	return 0, err
	//
	//}

	queryPost := fmt.Sprintf("INSERT INTO `%s` (`%v`,`%s`,`%s`) VALUES('%v', '%s','%s')", tablePost, uid, T, B, post.UserId, post.Title, post.Body)
	_ = p.db.QueryRow(queryPost)
	queryId := fmt.Sprintf("SELECT LAST_INSERT_ID(id) from posts")
	d, _ := p.db.Query(queryId)
	for d.Next() {
		var id Rest_Api.Post
		err := d.Scan(&id.Id)
		if err != nil {
			panic(err)
		}
		postId = id.Id
	}
	fmt.Println(postId)

	return postId, nil
}

//defer rows.Close()
//if err != nil {
//	return 0, fmt.Errorf("error in rows: %s", err.Error())
//}

//if err = rows.Scan(&id); err != nil {
//	//err := tx.Rollback()
//	//if err != nil {
//	//	return 0, err
//	//}
//	log.Fatalln("error", err.Error())
//	return id, err
//}
//fmt.Println(rows)
//defer rows.Close()
//queryComPost := fmt.Sprintf("INSERT INTO %s(`%v`)VALUES(%v)", tablePostCom, p_id, id)
//_, err = p.db.Query(queryComPost)
//if err != nil {
//	return 0, err
//}
//_, err = tx.Exec(queryComPost)
//if err != nil {
//	log.Fatalln("failed", err.Error())
//}
//rows.Scan(&id)

//	return 1, nil
//}
