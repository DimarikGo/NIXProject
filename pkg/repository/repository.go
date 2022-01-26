package repository

import (
	"Rest-Api/models"
	"gorm.io/gorm"
)

type Post interface {
	Add(post models.Post) (int, error)
}

type Comment interface {
	Add(comment models.Comment) (int, error)
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
