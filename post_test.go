package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/cucumber/godog"
	"github.com/stretchr/testify/assert"
	"goproject/database"
	"goproject/models"
	"goproject/utils"
	"net/http"
	"os"
	"testing"
)

var (
	authToken string
	baseURL   = "http://localhost:8080"
)

func anAccountExistsWithUsername(username string) error {
	authToken, _ = utils.GenerateJWT(username)
	fmt.Println(authToken)
	return nil
}

func theUserCreatesPostWithTitleAndContent(title, content string) error {
	postPayload := map[string]string{
		"title":   title,
		"content": content,
	}
	payloadBytes, _ := json.Marshal(postPayload)

	req, _ := http.NewRequest("POST", baseURL+"/api/posts", bytes.NewBuffer(payloadBytes))
	req.Header.Set("Authorization", "Bearer "+authToken)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	client.Do(req)

	return nil
}
func postShouldBeCreatedSuccessfullyWithTitleAndContent(title, content string) error {

	username, err := utils.GetUsernameFromToken(authToken)
	if err != nil {
		return fmt.Errorf("failed to extract username from token: %v", err)
	}

	var user models.User
	if err := database.DB.Where("username = ?", username).First(&user).Error; err != nil {
		return fmt.Errorf("failed to find user with username '%s': %v", username, err)
	}

	var post models.Post
	if err := database.DB.Where("user_id = ?", user.ID).Last(&post).Error; err != nil {
		return fmt.Errorf("failed to find the last post for user '%s': %v", username, err)
	}

	assert.Equal(nil, title, post.Title)
	assert.Equal(&testing.T{}, content, post.Content)

	return nil
}

func userShouldBeDirectedToHomePage() error {
	return nil
}

func InitializeScenario(ctx *godog.ScenarioContext) {
	database.InitDB()
	ctx.Given(`^an account exists with username "([^"]*)"$`, anAccountExistsWithUsername)
	ctx.When(`^the user creates post with title "([^"]*)" and content "([^"]*)"$`, theUserCreatesPostWithTitleAndContent)
	ctx.Then(`^post should be created successfully with title "([^"]*)" and content "([^"]*)"$`, postShouldBeCreatedSuccessfullyWithTitleAndContent)
	ctx.Then(`^user should be directed to home page$`, userShouldBeDirectedToHomePage)
}

func TestFeatures(t *testing.T) {
	opts := godog.Options{
		Output: os.Stdout,
		Format: "pretty", // or "progress" for a more compact output
		Paths:  []string{"features"},
	}
	godog.TestSuite{
		Name:                 "godogs",
		TestSuiteInitializer: nil,
		ScenarioInitializer:  InitializeScenario,
		Options:              &opts,
	}.Run()
}
