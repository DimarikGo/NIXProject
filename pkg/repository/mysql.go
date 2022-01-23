package repository

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

type Config struct {
	Name     string
	Password string
	DBName   string
}

func NewMysqlDb(cfg Config) (*sql.DB, error) {

	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@/%s", cfg.Name, cfg.Password, cfg.DBName))

	if err != nil {
		log.Fatal("no database connection")
	}
	log.Println("db connection")
	//defer db.Close().Error()
	return db, err
}
