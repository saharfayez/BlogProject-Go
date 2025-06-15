package users

import (
	"github.com/pkg/errors"
	"goproject/internal/app/business/interfaces/users"
	appMiddleware "goproject/internal/app/middleware"
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
	appMiddleware.ZapLogger.Debug("debug in user repo")
	appMiddleware.ZapLogger.Info("info in user repo")
	appMiddleware.ZapLogger.Error("error in user repo")
	return &user, errors.Wrap(err, "find user by username failed")
}

func (userRepo *userRepositoryImpl) Save(user *models.User) error {
	// gorm documentation mentions that parameter to create method should be pointer
	err := userRepo.db.Create(&user).Error
	return errors.Wrap(err, "create user failed")
}
