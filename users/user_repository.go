package users

import (
	"gorm.io/gorm"
)

type UserRepository interface {
	FindUserByUsername(username string) (*User, error)
	Save(user *User) error
}

type userRepositoryImpl struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepositoryImpl{db: db}
}

func (userRepo *userRepositoryImpl) FindUserByUsername(username string) (*User, error) {
	var user User
	err := userRepo.db.Where("username = ?", username).First(&user).Error
	return &user, err
}

func (userRepo *userRepositoryImpl) Save(user *User) error {
	// gorm documentation mentions that parameter to create method should be pointer
	return userRepo.db.Create(&user).Error
}
