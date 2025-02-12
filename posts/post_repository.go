package posts

import (
	"gorm.io/gorm"
)

type PostRepository interface {
	FindAll() ([]Post, error)
	FindById(id int) (*Post, error)
	Save(post *Post) error
	DeleteById(id int) error
	UpdateById(post *Post) error
}

type postRepositoryImpl struct {
	db *gorm.DB
}

func NewPostRepository(db *gorm.DB) PostRepository {
	return &postRepositoryImpl{db: db}
}

func (postRepo *postRepositoryImpl) FindAll() ([]Post, error) {
	var posts []Post
	err := postRepo.db.Find(&posts).Error
	return posts, err
}

func (postRepo *postRepositoryImpl) FindById(id int) (*Post, error) {
	var post Post
	err := postRepo.db.First(&post, id).Error
	return &post, err
}

func (postRepo *postRepositoryImpl) Save(post *Post) error {
	return postRepo.db.Create(&post).Error
}

func (postRepo *postRepositoryImpl) DeleteById(id int) error {
	var post Post
	return postRepo.db.Delete(&post, id).Error
}

func (postRepo *postRepositoryImpl) UpdateById(post *Post) error {
	return postRepo.db.Save(&post).Error
}
