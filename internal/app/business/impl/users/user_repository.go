package users

import (
	"github.com/pkg/errors"
	"goproject/internal/app/business/interfaces/users"
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
	if err != nil {
		return nil, errors.Wrap(err, "user not found")
	}
	return &user, nil
}

func (userRepo *userRepositoryImpl) Save(user *models.User) error {
	// gorm documentation mentions that parameter to create method should be pointer
	return userRepo.db.Create(&user).Error
}
