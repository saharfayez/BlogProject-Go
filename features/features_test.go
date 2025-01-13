package features

import (
	"github.com/cucumber/godog"
	"github.com/go-testfixtures/testfixtures/v3"
	"goproject/database"
	"log"
	"os"
	"testing"
)

func InitializeScenarios(ctx *godog.ScenarioContext) {
	InitializePostManagementScenario(ctx)
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
	db, _ := database.DB.DB()
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
