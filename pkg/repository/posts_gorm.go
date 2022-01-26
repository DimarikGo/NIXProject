package repository

import (
	"Rest-Api/models"
	"gorm.io/gorm"
)

type PostGorm struct {
	gorm gorm.DB
}

func NewPostGorm(gorm gorm.DB) *PostGorm {
	return &PostGorm{gorm: gorm}
}

func (g *PostGorm) Add(post models.Post) (int, error) {
	query := models.Post{
		Id:     post.Id,
		UserId: post.UserId,
		Title:  post.Title,
		Body:   post.Body,
	}
	g.gorm.Create(&query)
	return query.Id, nil
}
