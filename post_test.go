package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/cucumber/godog"
	"github.com/stretchr/testify/assert"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"testing"
)

var (
	response     *http.Response
	responseBody []byte
	authToken    string
	baseURL      = "http://localhost:8080"
)

func aUserIsLoggedInWithUsernameAndPassword(username, password string) error {

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
	responseBody, _ = io.ReadAll(response.Body)
	return nil
}

func theResponseStatusShouldBe(expectedStatus int) error {
	fmt.Println(expectedStatus)
	fmt.Println(response.StatusCode)
	assert.Equal(nil, expectedStatus, response.StatusCode, fmt.Sprintf("Expected status %d but got %d", expectedStatus, response.StatusCode))
	return nil
}

func theResponseBodyShouldBe(expectedMessage string) error {
	assert.Equal(nil, expectedMessage, string(responseBody), fmt.Sprintf("Expected message %s but got %s", expectedMessage, response.Body))
	return nil
}

func InitializeScenario(ctx *godog.ScenarioContext) {
	ctx.Step(`^a user is logged in with username "([^"]*)" and password "([^"]*)"$`, aUserIsLoggedInWithUsernameAndPassword)
	ctx.Step(`^the user sends a POST request to "([^"]*)" with title "([^"]*)" and content "([^"]*)"$`, theUserSendsAPOSTRequestToWithTitleAndContent)
	ctx.Step(`^the response status should be (\d+)$`, theResponseStatusShouldBe)
	ctx.Step(`^the response body should be "([^"]*)"$`, theResponseBodyShouldBe)

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
