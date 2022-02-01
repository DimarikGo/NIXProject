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

func (s *CommentService) Get(postId int) ([]models.Comment, error) {
	return s.repo.Get(postId)
}

func (s *CommentService) Add(comment *models.Comment) (*models.Comment, error) {
	return s.repo.Add(comment)
}
func (s *CommentService) Del(id int) (int, error) {
	return s.repo.Del(id)
}
func (s *CommentService) Update(comment *models.Comment, commentId int) *models.Comment {
	return s.repo.Update(comment, commentId)
}
