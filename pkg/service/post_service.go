package service

import (
	"Rest-Api"
	"Rest-Api/pkg/repository"
)

type PostService struct {
	repo repository.Post
}

func NewPostService(repo repository.Post) *PostService {
	return &PostService{repo: repo}
}

func (s *PostService) AddP(post Rest_Api.Post) (int, error) {
	return s.repo.AddP(post)
}
