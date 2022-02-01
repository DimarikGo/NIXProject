package repository

import (
	"Rest-Api/models"
	"fmt"
	"gorm.io/gorm"
)

type CommentGorm struct {
	gorm gorm.DB
}

func NewCommentGorm(gorm gorm.DB) *CommentGorm {
	return &CommentGorm{gorm: gorm}
}

func (c *CommentGorm) Get(postId int) ([]models.Comment, error) {
	var comments []models.Comment
	query := fmt.Sprintf("select * from comments where post_id =  %v", postId)
	c.gorm.Raw(query).Scan(&comments)
	fmt.Println(comments)
	return comments, nil
}
func (c *CommentGorm) Add(comment *models.Comment) (*models.Comment, error) {
	query := models.Comment{
		PostID: comment.PostID,
		Name:   comment.Name,
		Email:  comment.Email,
		Body:   comment.Body,
	}
	c.gorm.Create(&query)
	return comment, nil
}
func (c *CommentGorm) Del(id int) (int, error) {
	fmt.Println(id)
	c.gorm.Delete(&models.Comment{}, id)
	fmt.Println(id)
	return id, nil
}
func (c *CommentGorm) Update(comment *models.Comment, commentId int) *models.Comment {
	c.gorm.Model(&models.Comment{}).Where("id=?", commentId).Updates(models.Comment{

		PostID: comment.PostID,
		Name:   comment.Name,
		Email:  comment.Email,
		Body:   comment.Body,
	})
	return comment
}
