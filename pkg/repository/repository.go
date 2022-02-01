package repository

import (
	"Rest-Api/models"
	"gorm.io/gorm"
)

type Post interface {
	Add(post *models.Post) (*models.Post, error)
	Get(id int) (models.Post, error)
	Del(id int) (byte, error)
	Update(post *models.Post, postId int) *models.Post
}

type Comment interface {
	Get(postId int) ([]models.Comment, error)
	Del(id int) (int, error)
	Add(comment *models.Comment) (*models.Comment, error)
	Update(comment *models.Comment, commentId int) *models.Comment
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
