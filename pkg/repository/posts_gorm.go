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

func (g *PostGorm) Get(id int) models.Post {
	var post models.Post
	g.gorm.First(&post, "id=?", id)
	return post
}

func (g *PostGorm) Del(id int) (byte, error) {
	g.gorm.Delete(&models.Post{}, id)
	return byte(id), nil
}

func (g *PostGorm) Update(post models.Post) models.Post {

	g.gorm.Model(&models.Post{}).Where("id=?", post.Id).Updates(models.Post{
		UserId: post.UserId,
		Title:  post.Title,
		Body:   post.Body,
	})
	return post
}
