package repository

import (
	"goproject/models"
	"gorm.io/gorm"
)

type PostRepository interface {
	findAll() ([]models.Post, error)
	findById(id int) (*models.Post, error)
	save(post *models.Post) error
	deleteById(id int) error
	updateById(post *models.Post) error
}

type postRepositoryStruct struct {
	db *gorm.DB
}

func NewPostRepositoryStruct(db *gorm.DB) *postRepositoryStruct {
	return &postRepositoryStruct{db: db}
}

func (postRepo *postRepositoryStruct) findAll() ([]models.Post, error) {
	var posts []models.Post
	err := postRepo.db.Find(&posts).Error
	return posts, err
}

func (postRepo *postRepositoryStruct) findById(id int) (*models.Post, error) {
	var post models.Post
	err := postRepo.db.First(&post, id).Error
	return &post, err
}

func (postRepo *postRepositoryStruct) save(post *models.Post) error {
	return postRepo.db.Create(&post).Error
}

func (postRepo *postRepositoryStruct) deleteById(id int) error {
	var post models.Post
	return postRepo.db.Delete(&post, id).Error
}

func (postRepo *postRepositoryStruct) updateById(post *models.Post) error {
	return postRepo.db.Save(&post).Error
}
