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

func (s *PostService) Get(id int) models.Post {
	return s.repo.Get(id)
}

func (s *PostService) Del(id int) (byte, error) {
	return s.repo.Del(id)
}

func (s *PostService) Update(post models.Post) models.Post {
	return s.repo.Update(post)
}
