package repository

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Init() *gorm.DB {

	dsn := "root:11111111@tcp(127.0.0.1:3306)/NIXdb?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("not db connect")
	}
	return db
}
