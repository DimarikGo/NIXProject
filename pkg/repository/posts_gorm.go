package repository

import (
	"Rest-Api/models"
	"github.com/go-pg/pg/v10"
	"gorm.io/gorm"
	"log"
)

type PostGorm struct {
	gorm gorm.DB
}

func NewPostGorm(gorm gorm.DB) *PostGorm {
	return &PostGorm{gorm: gorm}
}

func (g *PostGorm) Add(post *models.Post) (*models.Post, error) {

	query := models.Post{
		Id:     post.Id,
		UserId: post.UserId,
		Title:  post.Title,
		Body:   post.Body,
	}
	g.gorm.Create(&query)
	return post, nil
}

func (g *PostGorm) Get(id int) (models.Post, error) {
	var post models.Post
	if err := g.gorm.First(&post, "id=?", id); err.Error != nil {
		if g.gorm.Error == pg.ErrNoRows {
			log.Fatalf("record not found: %s", err.Error)
		}
	}
	return post, nil
}

func (g *PostGorm) Del(id int) (byte, error) {
	g.gorm.Delete(&models.Post{}, id)
	return byte(id), nil
}

func (g *PostGorm) Update(post *models.Post, postId int) *models.Post {

	g.gorm.Model(&models.Post{}).Where("id=?", postId).Updates(models.Post{
		UserId: post.UserId,
		Title:  post.Title,
		Body:   post.Body,
	})

	return post
}
