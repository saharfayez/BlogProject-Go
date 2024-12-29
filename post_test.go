package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/cucumber/godog"
	"github.com/stretchr/testify/assert"
	"goproject/models"
	"io/ioutil"
	"net/http"
)

var (
	response      *http.Response
	responseBody  []byte
	loggedInUser  models.User
	authToken     string
	baseURL       = "http://localhost:8080"
	postTitle     string
	postContent   string
	createdPostID uint
)

func aUserIsLoggedInWithUsernameAndPassword(username, password string) error {
	loggedInUser = models.User{Username: username, Password: password}

	// Sign up user if they don't exist
	// Log in the user
	loginPayload := map[string]string{
		"username": username,
		"password": password,
	}
	payloadBytes, _ := json.Marshal(loginPayload)
	resp, err := http.Post(baseURL+"/login", "application/json", bytes.NewBuffer(payloadBytes))
	if err != nil {
		return fmt.Errorf("error logging in user: %w", err)
	}
	body, _ := ioutil.ReadAll(resp.Body)
	var loginResponse map[string]string
	json.Unmarshal(body, &loginResponse)

	authToken = loginResponse["Token"]
	return nil
}

func theUserSendsAPOSTRequestToWithTitleAndContent(api, title, content string) error {
	postTitle = title
	postContent = content

	postPayload := map[string]string{
		"title":   title,
		"content": content,
	}
	payloadBytes, _ := json.Marshal(postPayload)

	req, _ := http.NewRequest("POST", baseURL+api, bytes.NewBuffer(payloadBytes))
	req.Header.Set("Authorization", "Bearer "+authToken)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	response, _ = client.Do(req)
	responseBody, _ = ioutil.ReadAll(response.Body)
	return nil
}

func theResponseStatusShouldBe(expectedStatus int) error {
	fmt.Println(expectedStatus)
	fmt.Println(response.StatusCode)
	assert.Equal(nil, expectedStatus, response.StatusCode, fmt.Sprintf("Expected status %d but got %d", expectedStatus, response.StatusCode))
	return nil
}

func theResponseBodyShouldBe(expectedMessage string) error {
	fmt.Println(responseBody)
	assert.Equal(nil, expectedMessage, string(responseBody), fmt.Sprintf("Expected message %s but got %s", expectedMessage, response.Body))
	return nil
}

func InitializeScenario(ctx *godog.ScenarioContext) {
	ctx.Step(`^a user is logged in with username "([^"]*)" and password "([^"]*)"$`, aUserIsLoggedInWithUsernameAndPassword)
	ctx.Step(`^the user sends a POST request to "([^"]*)" with title "([^"]*)" and content "([^"]*)"$`, theUserSendsAPOSTRequestToWithTitleAndContent)
	ctx.Step(`^the response status should be (\d+)$`, theResponseStatusShouldBe)
	ctx.Step(`^the response body should be "([^"]*)"$`, theResponseBodyShouldBe)

}
