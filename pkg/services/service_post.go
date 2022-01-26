package services

import (
	"Rest-Api/models"
	"Rest-Api/pkg/repository"
)

type PostService struct {
	repo repository.Post
}

func NewPostService(repo repository.Post) *PostService {
	return &PostService{repo: repo}
}

func (s *PostService) Add(post models.Post) (int, error) {
	return s.repo.Add(post)
}
