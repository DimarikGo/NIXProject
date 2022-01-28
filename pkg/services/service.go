package services

import (
	"Rest-Api/models"
	"Rest-Api/pkg/repository"
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

type Service struct {
	Post
	Comment
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		Post:    NewPostService(repo.Post),
		Comment: NewCommentService(repo.Comment)}
}
