package services

import (
	"Rest-Api/models"
	"Rest-Api/pkg/repository"
)

type Authorization interface {
	CreateUser(user *models.User) (int, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(token string) (int, error)
}

type Post interface {
	Add(post *models.Post, id int) (*models.Post, error)
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

type Service struct {
	Authorization
	Post
	Comment
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repo.Authorization),
		Post:          NewPostService(repo.Post),
		Comment:       NewCommentService(repo.Comment)}
}
