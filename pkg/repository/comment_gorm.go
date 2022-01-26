package repository

import (
	"Rest-Api/models"
	"gorm.io/gorm"
)

type CommentGorm struct {
	gorm gorm.DB
}

func NewCommentGorm(gorm gorm.DB) *CommentGorm {
	return &CommentGorm{gorm: gorm}
}

func (c *CommentGorm) Add(comments models.Comment) (int, error) {

	query := models.Comment{
		Id:     comments.Id,
		PostID: comments.PostID,
		Name:   comments.Name,
		Email:  comments.Email,
		Body:   comments.Body,
	}
	c.gorm.Create(&query)

	return 0, nil
}
