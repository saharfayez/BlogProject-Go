package posts

import "goproject/internal/app/models"

type PostService interface {
	CreatePost(username string, post *models.Post) error
}
