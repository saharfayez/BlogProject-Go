package interfaces

import "goproject/models"

type PostRepository interface {
	FindAll() ([]models.Post, error)
	FindById(id int) (*models.Post, error)
	Save(post *models.Post) error
	DeleteById(id int) error
	UpdateById(post *models.Post) error
}
