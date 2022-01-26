package services

import (
	"Rest-Api/models"
	"Rest-Api/pkg/repository"
)

type Post interface {
	Add(post models.Post) (int, error)
}

type Comment interface {
	Add(comment models.Comment) (int, error)
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
