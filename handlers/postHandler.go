package handlers

import (
	"encoding/json"
	"errors"
	"github.com/gorilla/mux"
	"goproject/database"
	"goproject/models"
	"goproject/utils"
	"net/http"
)

func GetPosts() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		var posts []models.Post

		result := database.DB.Find(&posts)
		if result.Error != nil {
			http.Error(w, "Error querying posts", http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(posts)
	}
}

func CreatePost() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		username, err := utils.GetUsernameFromContext(w, r)
		if err != nil {
			return
		}

		var user models.User
		if err := database.DB.Where("username = ?", username).First(&user).Error; err != nil {
			http.Error(w, "User not found", http.StatusNotFound)
			return
		}

		var post models.Post
		json.NewDecoder(r.Body).Decode(&post)
		post.UserID = user.ID

		if err := database.DB.Create(&post).Error; err != nil {
			http.Error(w, "Error creating post", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(post)
	}
}

func GetPost() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		params := mux.Vars(r)
		var post models.Post
		result := database.DB.First(&post, params["id"])
		if result.Error != nil {
			http.Error(w, "Error getting post", http.StatusNotFound)
			return
		}
		w.WriteHeader(http.StatusFound)
		json.NewEncoder(w).Encode(post)
	}
}
func UpdatePost() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		var post models.Post
		var user models.User
		user, post, err := authorizePost(w, r)
		if err != nil {
			return
		}
		var updatedPost models.Post
		json.NewDecoder(r.Body).Decode(&updatedPost)
		updatedPost.UserID = user.ID
		updatedPost.ID = post.ID
		if err := database.DB.Save(&updatedPost).Error; err != nil {
			http.Error(w, "Error updating post", http.StatusInternalServerError)
		}
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(updatedPost)

	}
}

func DeletePost() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		params := mux.Vars(r)
		_, _, err := authorizePost(w, r)
		if err != nil {
			return
		}

		var post models.Post
		result := database.DB.Delete(&post, params["id"])
		if result.Error != nil {
			http.Error(w, "Error deleting post", http.StatusNotFound)
			return
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode("Post deleted successfully")
	}

}
func authorizePost(w http.ResponseWriter, r *http.Request) (models.User, models.Post, error) {

	params := mux.Vars(r)
	username, err := utils.GetUsernameFromContext(w, r)
	if err != nil {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return models.User{}, models.Post{}, errors.New("Unauthorized")
	}
	var post models.Post
	if err := database.DB.First(&post, params["id"]).Error; err != nil {
		http.Error(w, "Post not found", http.StatusNotFound)
		return models.User{}, models.Post{}, errors.New("Post not found")
	}
	var user models.User
	err = database.DB.Where("username = ?", username).First(&user).Error
	if err != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return models.User{}, models.Post{}, errors.New("User not found")
	}

	if post.UserID != user.ID {
		http.Error(w, "User not allowed to update or delete this post", http.StatusUnauthorized)
		return models.User{}, models.Post{}, errors.New("Forbidden")
	}

	return user, post, nil
}
