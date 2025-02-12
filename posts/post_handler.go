package posts

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"goproject/context"
	"goproject/users"
	"goproject/utils"
	"net/http"
)

func GetPosts(c echo.Context) error {

	c.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	var posts []Post

	result := context.Context.DB.Find(&posts)
	if result.Error != nil {
		return c.String(http.StatusInternalServerError, "Error querying posts")
	}
	return c.JSON(http.StatusOK, posts)
}

func CreatePost(c echo.Context) error {
	c.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	username := utils.GetTokenFromContext(c)

	var user users.User
	if err := context.Context.DB.Where("username = ?", username).First(&user).Error; err != nil {
		return c.String(http.StatusNotFound, "User not found")
	}

	var post Post
	if err := c.Bind(&post); err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	post.UserID = user.ID

	if err := context.Context.DB.Create(&post).Error; err != nil {
		fmt.Println("Error creating post:", err)
		return c.String(http.StatusInternalServerError, "Error creating post")
	}

	return c.JSON(http.StatusCreated, post)
}

func GetPost(c echo.Context) error {

	c.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	id := c.Param("id")
	var post Post
	result := context.Context.DB.First(&post, id)
	if result.Error != nil {
		return c.String(http.StatusNotFound, "Error getting post")
	}
	return c.JSON(http.StatusFound, post)
}

func UpdatePost(c echo.Context) error {
	c.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	var user users.User
	var post Post
	user, post, err := authorizePost(c)
	if err != nil {
		return err
	}

	var updatedPost Post
	if err := c.Bind(&updatedPost); err != nil {
		return c.String(http.StatusBadRequest, "Error parsing body")
	}

	updatedPost.UserID = user.ID
	updatedPost.ID = post.ID

	if err := context.Context.DB.Save(&updatedPost).Error; err != nil {
		c.Logger().Error(err)
		return c.String(http.StatusInternalServerError, "Error updating post")
	}

	return c.JSON(http.StatusOK, updatedPost)
}

func DeletePost(c echo.Context) error {
	c.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	id := c.Param("id")
	_, _, err := authorizePost(c)
	if err != nil {
		return err
	}

	var post Post
	result := context.Context.DB.Delete(&post, id)
	if result.Error != nil {
		return c.String(http.StatusNotFound, "Error deleting post")
	}
	return c.String(http.StatusOK, "post deleted successfully")
}

func authorizePost(c echo.Context) (users.User, Post, error) {
	username := utils.GetTokenFromContext(c)
	id := c.Param("id")

	var post Post
	if err := context.Context.DB.First(&post, id).Error; err != nil {
		c.Logger().Error(err)
		return users.User{}, Post{}, echo.NewHTTPError(http.StatusNotFound, "Post not found")
	}

	var user users.User
	err := context.Context.DB.Where("username = ?", username).First(&user).Error
	if err != nil {
		c.Logger().Error(err)
		return users.User{}, Post{}, echo.NewHTTPError(http.StatusNotFound, "User not found")
	}

	if post.UserID != user.ID {
		c.Logger().Error("User not authorized to update or delete this post")
		return users.User{}, Post{}, echo.NewHTTPError(http.StatusUnauthorized, "User not authorized to update or delete this post")
	}

	return user, post, nil
}
