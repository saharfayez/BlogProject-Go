package features

import (
	"context"
	"fmt"
	"github.com/cucumber/godog"
	"github.com/go-testfixtures/testfixtures/v3"
	"goproject/database"
	"log"
	"os"
	"strings"
	"testing"
)

func InitializeScenarios(ctx *godog.ScenarioContext) {
	state := ScenarioState{
		data: make(map[string]interface{}),
	}
	ctx.Before(func(ctx context.Context, sc *godog.Scenario) (context.Context, error) {

		state = ScenarioState{
			data: make(map[string]interface{}),
		}
		fmt.Println("Before each scenario")
		seedTags := extractSeedTags(sc)
		for _, tag := range seedTags {
			if err := LoadFixtures("../fixtures/" + tag + ".yml"); err != nil {
				return ctx, fmt.Errorf("failed to load fixtures: %v", err)
			}
		}
		return ctx, nil
	})
	InitializePostManagementScenario(ctx, &state)
}

func TestFeatures(t *testing.T) {
	//flag.Parse()
	opts := godog.Options{
		Output: os.Stdout,
		Format: "pretty", // or "progress" for a more compact output
		Paths:  []string{"."},
		//Tags:     godogTags, // use parsed tags
		TestingT: t, // Integrate with go test
	}
	godog.TestSuite{
		Name:                 "posts",
		TestSuiteInitializer: InitializeTestSuite,
		ScenarioInitializer:  InitializeScenarios,
		Options:              &opts,
	}.Run()
}

func LoadFixtures(files ...string) error {
	db, err := database.DB.DB()
	if err != nil {
		return err
	}
	fixtures, err := testfixtures.New(
		testfixtures.Database(db),
		testfixtures.Dialect("mysql"),
		testfixtures.DangerousSkipTestDatabaseCheck(),
		testfixtures.FilesMultiTables(files...),
	)
	if err != nil {
		log.Fatal("Error initializing testfixtures: ", err)
	}
	return fixtures.Load()
}

func InitializeTestSuite(context *godog.TestSuiteContext) {
	context.BeforeSuite(func() {
		database.InitDB()
	})
}

func extractSeedTags(sc *godog.Scenario) []string {
	tags := []string{}
	for _, tag := range sc.Tags {
		if strings.HasSuffix(tag.Name, "_seed") {
			seedTag := strings.TrimSuffix(tag.Name, "_seed")
			seedTag = strings.TrimPrefix(seedTag, "@")
			tags = append(tags, seedTag)
		}
	}
	return tags
}
