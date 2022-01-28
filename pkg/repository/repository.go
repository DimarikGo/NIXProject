package repository

import (
	"Rest-Api/models"
	"gorm.io/gorm"
)

type Post interface {
	Add(post models.Post) (int, error)
	Get(id int) models.Post
	Del(id int) (byte, error)
	Update(post models.Post) models.Post
}

type Comment interface {
	Get(postId int) []models.Comment
	Del(id int) (int, error)
	Add(comment models.Comment) (int, error)
	Update(comment models.Comment) models.Comment
}

type Repository struct {
	Post
	Comment
}

func NewRepository(gorm gorm.DB) *Repository {
	return &Repository{
		Post:    NewPostGorm(gorm),
		Comment: NewCommentGorm(gorm)}
}
