package models

import "gorm.io/gorm"

type Post struct {
	gorm.Model
	Id     int    `gorm:"<-"`
	UserId int    `gorm:"<-"`
	Title  string `gorm:"<-"`
	Body   string `gorm:"<-"`
}

type Comment struct {
	gorm.Model
	PostID int    `json:"postId" gorm:"<-"`
	Id     int    `json:"id" gorm:"<-"`
	Name   string `json:"name" gorm:"<-"`
	Email  string `json:"email" gorm:"<-"`
	Body   string `json:"body" gorm:"<-"`
}

type Comments []struct {
	PostId int    `json:"postId" gorm:"<-"`
	Id     int    `json:"id" gorm:"<-"`
	Name   string `json:"name" gorm:"<-"`
	Email  string `json:"email" gorm:"<-"`
	Body   string `json:"body" gorm:"<-"`
}
