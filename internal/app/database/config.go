package database

import (
	"context"
	"errors"
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	db_postgres "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
	appcontext "goproject/internal/app/interfaces/context"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

var ShutDownTestContainer func()

func InitDB() (*gorm.DB, error) {
	var err error

	db, err := gorm.Open(getDB(), &gorm.Config{})
	if err != nil {
		log.Fatal("Error connecting to the database: ", err)
	}

	runMigrations(db)

	return db, nil
}

func getDB() gorm.Dialector {
	if appcontext.Context.GetPropertiesConfig().GetProfile() != "test" {
		return postgres.Open(appcontext.Context.GetPropertiesConfig().GetDatabaseUrl())
	} else {
		// Fallback to PostgreSQL Testcontainer
		ctx := context.Background()
		dsn, err := startTestContainer(ctx)
		if err != nil {
			panic(fmt.Sprintf("Failed to start PostgreSQL container: %v", err))
		}
		return postgres.Open(dsn)
	}
}

func startTestContainer(ctx context.Context) (string, error) {
	var env = map[string]string{
		"POSTGRES_PASSWORD": "postgres",
		"POSTGRES_USER":     "postgres",
		"POSTGRES_DB":       "postgres",
	}
	var port = "5432/tcp"
	// Define the PostgreSQL container request
	req := testcontainers.ContainerRequest{
		Image:        "postgres:15-alpine",
		ExposedPorts: []string{port},
		Env:          env,
		WaitingFor:   wait.ForLog("database system is ready to accept connections").WithOccurrence(2),
	}

	// Start the PostgreSQL container
	container, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
	})

	if err != nil {
		return "", fmt.Errorf("Failed to start PostgreSQL container: %v", err)
	}

	// Get the host and port of the PostgreSQL container
	_, err = container.Host(ctx)
	if err != nil {
		return "", fmt.Errorf("failed to get PostgreSQL container host: %v", err)
	}

	mappedPort, err := container.MappedPort(ctx, "5432")
	if err != nil {
		return "", fmt.Errorf("failed to get PostgreSQL container port: %v", err)
	}
	fmt.Printf("PORT NUMBER: %d", mappedPort.Port())

	//time.Sleep(5 * time.Second)
	// Construct the connection string for GORM

	dsn := fmt.Sprintf("host=127.0.0.1 user=postgres password=postgres dbname=postgres port=%s sslmode=disable TimeZone=Asia/Jakarta",
		mappedPort.Port())

	ShutDownTestContainer = func() {
		if err := container.Terminate(ctx); err != nil {
			fmt.Printf("Failed to terminate PostgreSQL container: %v\n", err)
		}
	}
	return dsn, nil
}

func runMigrations(db *gorm.DB) {
	sqlDB, _ := db.DB()

	var migration *migrate.Migrate
	var err error
	var migrationPath string

	currentDirectory, currentDirName := getCurrentDirectory()
	fmt.Println("current", currentDirectory)

	if currentDirName == "tests" {
		migrationPath = filepath.Join("file://", currentDirectory, "..", "..", "/internal/app/database/migrations")
	} else {
		migrationPath = filepath.Join("file://", currentDirectory, "/internal/app/database/migrations")
	}

	driver, err := db_postgres.WithInstance(sqlDB, &db_postgres.Config{})
	if err != nil {
		log.Fatal("Failed to create migration driver:", err)
	}
	migration, err = migrate.NewWithDatabaseInstance(
		migrationPath,
		"postgres",
		driver,
	)
	if err != nil {
		log.Fatal("Failed to initialize migrate instance:", err)
	}

	err = migration.Up()
	if err != nil && !errors.Is(migrate.ErrNoChange, err) {
		log.Fatal("Failed to apply migrations:", err)
	}

	log.Println("Migrations applied successfully!")
}

func getCurrentDirectory() (string, string) {

	currentDirectory, err := os.Getwd()
	if err != nil {
		log.Fatal("Failed to get current directory", err)
	}

	var directories []string
	if runtime.GOOS == "windows" {
		directories = strings.Split(currentDirectory, "\\")
	} else {
		directories = strings.Split(currentDirectory, "/")
	}

	currentDirName := directories[len(directories)-1]

	return currentDirectory, currentDirName
}
