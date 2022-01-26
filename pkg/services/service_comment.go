package services

import (
	"Rest-Api/models"
	"Rest-Api/pkg/repository"
)

type CommentService struct {
	repo repository.Comment
}

func NewCommentService(repo repository.Comment) *CommentService {
	return &CommentService{repo: repo}
}

func (s *CommentService) Add(comment models.Comment) (int, error) {
	return s.repo.Add(comment)
}
