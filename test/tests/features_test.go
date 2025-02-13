package tests

import (
	"context"
	"fmt"
	"github.com/cucumber/godog"
	"github.com/go-testfixtures/testfixtures/v3"
	"goproject/internal/app/database"
	_ "goproject/internal/app/impl/context"
	contextpkg "goproject/internal/app/interfaces/context"
	"goproject/internal/app/server"
	"log"
	"os"
	"strings"
	"testing"
)

type DatabaseOperation string

const (
	insertOperation      DatabaseOperation = "insert"
	cleanInsertOperation DatabaseOperation = "clean_insert"
)

var baseURL = "http://localhost:8080"

type ScenarioState struct {
	data map[string]interface{}
}

type DatabaseSetup struct {
	fileName  string
	operation DatabaseOperation
}

func InitializeScenarios(ctx *godog.ScenarioContext) {
	state := ScenarioState{
		data: make(map[string]interface{}),
	}
	ctx.Before(func(ctx context.Context, sc *godog.Scenario) (context.Context, error) {

		state = ScenarioState{
			data: make(map[string]interface{}),
		}

		databaseSetups := extractDatabaseSetups(sc)

		fmt.Println("Before each scenario")
		for _, databaseSetup := range *databaseSetups {
			if err := loadFixtures(databaseSetup.operation, "../fixtures/"+databaseSetup.fileName+".yml"); err != nil {
				return ctx, fmt.Errorf("failed to load fixtures: %v", err)
			}
		}
		return ctx, nil
	})
	InitializePostManagementScenario(ctx, &state)
}

func extractDatabaseSetups(sc *godog.Scenario) *[]DatabaseSetup {
	var setups []DatabaseSetup
	for _, tag := range sc.Tags {
		var setup DatabaseSetup
		if strings.HasSuffix(tag.Name, "_seed") {
			seedTag := strings.TrimSuffix(tag.Name, "_seed")
			seedTag = strings.TrimPrefix(seedTag, "@")
			if strings.HasSuffix(seedTag, "_cleaninsert") {
				seedTag = strings.TrimSuffix(seedTag, "_cleaninsert")
				setup.fileName = seedTag
				setup.operation = cleanInsertOperation
				setups = append(setups, setup)
			} else if strings.HasSuffix(seedTag, "_insert") {
				seedTag = strings.TrimSuffix(seedTag, "_insert")
				setup.fileName = seedTag
				setup.operation = insertOperation
				setups = append(setups, setup)
			}

		}
	}
	return &setups
}

func loadFixtures(operation DatabaseOperation, files ...string) error {
	db, err := contextpkg.Context.GetDB().DB()
	if err != nil {
		return err
	}
	options := []func(*testfixtures.Loader) error{
		testfixtures.Database(db),
		testfixtures.Dialect(contextpkg.Context.GetDB().Dialector.Name()),
		testfixtures.DangerousSkipTestDatabaseCheck(),
		testfixtures.FilesMultiTables(files...),
	}

	if operation == insertOperation {
		options = append(options, testfixtures.DangerousSkipCleanupFixtureTables())
	}

	fixtures, err := testfixtures.New(
		options...,
	)

	if err != nil {
		log.Fatal("Error initializing testfixtures: ", err)
	}
	return fixtures.Load()
}

func InitializeTestSuite(context *godog.TestSuiteContext) {
	context.BeforeSuite(func() {
		go server.Serve()
	})
	context.AfterSuite(func() {
		database.ShutDownTestContainer()
	})
}

func TestFeatures(t *testing.T) {
	//flag.Parse()
	opts := godog.Options{
		Output: os.Stdout,
		Format: "pretty", // or "progress" for a more compact output
		Paths:  []string{"../features"},
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
