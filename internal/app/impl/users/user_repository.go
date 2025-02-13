package users

import (
	"goproject/internal/app/interfaces/users"
	"goproject/internal/app/models"
	"gorm.io/gorm"
)

type userRepositoryImpl struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) users.UserRepository {
	return &userRepositoryImpl{db: db}
}

func (userRepo *userRepositoryImpl) FindUserByUsername(username string) (*models.User, error) {
	var user models.User
	err := userRepo.db.Where("username = ?", username).First(&user).Error
	return &user, err
}

func (userRepo *userRepositoryImpl) Save(user *models.User) error {
	// gorm documentation mentions that parameter to create method should be pointer
	return userRepo.db.Create(&user).Error
}
