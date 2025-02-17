package posts

import "goproject/internal/app/models"

type PostRepository interface {
	FindAll() ([]models.Post, error)
	FindById(id int) (*models.Post, error)
	Create(post *models.Post) error
	DeleteById(id int) error
	UpdateById(post *models.Post) error
}
