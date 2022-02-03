package repository

import (
	"Rest-Api/models"
	"gorm.io/gorm"
)

type AuthGorm struct {
	gorm gorm.DB
}

func NewAuthGorm(gorm gorm.DB) *AuthGorm {
	return &AuthGorm{gorm: gorm}
}

func (g *AuthGorm) CreateUser(user *models.User) (int, error) {
	query := models.User{
		Name:     user.Name,
		Username: user.Username,
		Password: user.Password,
	}
	g.gorm.Create(&query)
	return query.Id, nil
}

func (g *AuthGorm) GetUser(username, password string) (models.User, error) {
	var users models.User
	g.gorm.Where("username = ? AND password= ?", username, password).Find(&users)
	return users, nil
}
