package service

import (
	"Rest-Api"
	"Rest-Api/pkg/repository"
)

type CommentService struct {
	repo repository.Comments
}

func NewCommentService(repo repository.Comments) *CommentService {
	return &CommentService{repo: repo}
}

func (s *CommentService) AddC(comment Rest_Api.Comment) (int, error) {
	return s.repo.AddC(comment)
}
