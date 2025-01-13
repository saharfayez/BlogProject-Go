package features

import (
	"github.com/cucumber/godog"
	"github.com/go-testfixtures/testfixtures/v3"
	"goproject/database"
	"log"
	"os"
	"testing"
)

func TestFeature(t *testing.T) {
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
		ScenarioInitializer:  InitializeScenario,
		Options:              &opts,
	}.Run()
}

func LoadFixtures(files ...string) error {
	con, _ := database.DB.DB()
	fixtures, err := testfixtures.New(
		testfixtures.Database(con), // Use the underlying *sql.DB
		testfixtures.Dialect("mysql"),
		//testfixtures.Paths(files...), // Path to your fixture files
		testfixtures.DangerousSkipTestDatabaseCheck(),
		testfixtures.FilesMultiTables(files...),
	)
	if err != nil {
		log.Fatal("Error initializing testfixtures: ", err)
	}
	return fixtures.Load()
}
