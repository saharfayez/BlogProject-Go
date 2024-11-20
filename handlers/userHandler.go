package handlers

import (
	"encoding/json"
	"goproject/database"
	"goproject/models"
	"goproject/utils"
	"log"
	"net/http"
	"strconv"

	"golang.org/x/crypto/bcrypt"
)

func Signup(w http.ResponseWriter, r *http.Request) {

	var user models.User
	json.NewDecoder(r.Body).Decode(&user)

	var existingUser models.User
	error := database.DB.Where("username = ?", user.Username).First(&existingUser).Error

	if error == nil {
		http.Error(w, "Username already exists. Please choose another one.", http.StatusBadRequest)
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		http.Error(w, "Error hashing password", http.StatusInternalServerError)
		log.Println(err)
		return
	}
	user.Password = string(hashedPassword)

	result := database.DB.Create(&user)
	if result.Error != nil {
		http.Error(w, "Error creating user", http.StatusInternalServerError)
		log.Println(err)

		return
	}

	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(map[string]string{"id": strconv.Itoa(user.ID), "username": user.Username})
}

func Login(w http.ResponseWriter, r *http.Request) {
	var credentials models.User
	json.NewDecoder(r.Body).Decode(&credentials)

	var user models.User
	result := database.DB.Where("username = ?", credentials.Username).First(&user)
	if result.Error != nil {
		http.Error(w, "User Not Found", http.StatusNotFound)
		return
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(credentials.Password))
	if err != nil {
		http.Error(w, "Passwords are not compatible", http.StatusUnauthorized)
		return
	}

	token, err := utils.GenerateJWT(user.Username)
	if err != nil {
		http.Error(w, "Can not generate token", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{"token": token})
}
