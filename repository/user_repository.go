package repository

import (
	"goproject/models"
	"gorm.io/gorm"
)

type UserRepository interface {
	findUserByUsername(username string) (*models.User, error)
	save(user *models.User) error
}

type userRepositoryStruct struct {
	db *gorm.DB
}

func NewUserRepositoryStruct(db *gorm.DB) *userRepositoryStruct {
	return &userRepositoryStruct{db: db}
}

func (userRepo *userRepositoryStruct) findUserByUsername(username string) (*models.User, error) {
	var user models.User
	err := userRepo.db.Where("username = ?", username).First(&user).Error
	return &user, err
}

func (userRepo *userRepositoryStruct) save(user *models.User) error {
	return userRepo.db.Create(&user).Error
}
