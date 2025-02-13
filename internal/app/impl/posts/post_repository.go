package posts

import (
	"goproject/internal/app/interfaces/posts"
	"goproject/internal/app/models"
	"gorm.io/gorm"
)

type postRepositoryImpl struct {
	db *gorm.DB
}

func NewPostRepository(db *gorm.DB) posts.PostRepository {
	return &postRepositoryImpl{db: db}
}

func (postRepo *postRepositoryImpl) FindAll() ([]models.Post, error) {
	var posts []models.Post
	err := postRepo.db.Find(&posts).Error
	return posts, err
}

func (postRepo *postRepositoryImpl) FindById(id int) (*models.Post, error) {
	var post models.Post
	err := postRepo.db.First(&post, id).Error
	return &post, err
}

func (postRepo *postRepositoryImpl) Save(post *models.Post) error {
	return postRepo.db.Create(&post).Error
}

func (postRepo *postRepositoryImpl) DeleteById(id int) error {
	var post models.Post
	return postRepo.db.Delete(&post, id).Error
}

func (postRepo *postRepositoryImpl) UpdateById(post *models.Post) error {
	return postRepo.db.Save(&post).Error
}
