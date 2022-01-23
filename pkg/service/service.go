package service

import (
	"Rest-Api"
	"Rest-Api/pkg/repository"
)

type Post interface {
	AddP(post Rest_Api.Post) (int, error)
}

type Comments interface {
	AddC(postid int, comment Rest_Api.Comment) (int, error)
}

type Service struct {
	Post
	Comments
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		Post:     NewPostService(repo.Post),
		Comments: NewCommentService(repo.Comments)}
}
