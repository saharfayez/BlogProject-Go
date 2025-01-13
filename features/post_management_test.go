package features

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/cucumber/godog"
	"github.com/stretchr/testify/assert"
	"goproject/database"
	"goproject/models"
	"goproject/utils"
	"net/http"
)

var baseURL = "http://localhost:8080"

type ScenarioState struct {
	authToken  string
	userName   string
	statusCode int
}

func (state *ScenarioState) anAccountExistsWithUsername(username string) error {
	return nil
}

func (state *ScenarioState) userIsLoggedInWithUsername(username string) error {
	state.authToken, _ = utils.GenerateJWT(username)
	state.userName = username
	return nil
}

func (state *ScenarioState) theUserCreatesPostWithTitleAndContent(title, content string) error {
	postPayload := map[string]string{
		"title":   title,
		"content": content,
	}
	payloadBytes, _ := json.Marshal(postPayload)

	req, _ := http.NewRequest("POST", baseURL+"/api/posts", bytes.NewBuffer(payloadBytes))
	req.Header.Set("Authorization", "Bearer "+state.authToken)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	response, err := client.Do(req)
	if err != nil {
		return err
	}
	state.statusCode = response.StatusCode
	return nil
}

func (state *ScenarioState) postShouldBeCreatedSuccessfullyWithTitleAndContent(title, content string) error {
	var user models.User
	if err := database.DB.Where("username = ?", state.userName).First(&user).Error; err != nil {
		return fmt.Errorf("failed to find user with username '%s': %v", state.userName, err)
	}

	var post models.Post
	if err := database.DB.Where("user_id = ?", user.ID).Last(&post).Error; err != nil {
		return fmt.Errorf("failed to find the last post for user '%s': %v", state.userName, err)
	}

	assert.Equal(nil, title, post.Title)
	assert.Equal(nil, content, post.Content)
	assert.Equal(nil, state.statusCode, 201)

	return nil
}

func (state *ScenarioState) userShouldBeReDirectedToHomePage() error {
	return nil
}

//func userShouldBeDirectedToLandingPage() error {
//	return nil
//}

//var godogTags string // Variable to hold tags

//func init() {
//	flag.StringVar(&godogTags, "godog.tags", "", "Tags to filter scenarios")
//}

func InitializePostManagementScenario(ctx *godog.ScenarioContext) {
	state := &ScenarioState{}
	ctx.Before(func(ctx context.Context, sc *godog.Scenario) (context.Context, error) {
		//fmt.Println("Before each scenario")
		*state = ScenarioState{}
		if sc.Uri == "post_management.feature" && sc.Name == "Create a post successfully" {
			if err := LoadFixtures("create_post_successfully.yml"); err != nil {
				return ctx, fmt.Errorf("failed to load fixtures: %v", err)
			}
		}
		return ctx, nil
	})
	ctx.Given(`^an account exists with username "([^"]*)"$`, state.anAccountExistsWithUsername)
	ctx.Given(`^user is logged in with username "([^"]*)"$`, state.userIsLoggedInWithUsername)
	ctx.When(`^the user creates a post with title "([^"]*)" and content "([^"]*)"$`, state.theUserCreatesPostWithTitleAndContent)
	ctx.Then(`^post should be created successfully with title "([^"]*)" and content "([^"]*)"$`, state.postShouldBeCreatedSuccessfullyWithTitleAndContent)
	ctx.Then(`^user should be redirected to home page$`, state.userShouldBeReDirectedToHomePage)
	//ctx.Step(`^user should be directed to landing page$`, userShouldBeDirectedToLandingPage)
}
