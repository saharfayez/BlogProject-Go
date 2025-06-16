package posts

import (
	"goproject/internal/app/business/interfaces/posts"
	"goproject/internal/app/business/interfaces/users"
	appMiddleware "goproject/internal/app/middleware"
	"goproject/internal/app/models"
)

type postServiceImpl struct {
	postRepo    posts.PostRepository
	userService users.UserService
}

func NewPostService(postRepo posts.PostRepository, userService users.UserService) posts.PostService {
	return &postServiceImpl{postRepo: postRepo, userService: userService}
}

func (postServiceImpl *postServiceImpl) CreatePost(username string, post *models.Post) error {
	user, err := postServiceImpl.userService.FindUser(username)
	appMiddleware.ZapLogger.Info("logger in post service")
	if err != nil {
		return err
	}
	post.UserID = user.ID
	return postServiceImpl.postRepo.Create(post)
}
