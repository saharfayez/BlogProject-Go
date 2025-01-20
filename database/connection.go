package database

import (
	"context"
	"fmt"
	"github.com/joho/godotenv"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
	"gorm.io/driver/postgres"
	//"github.com/glebarez/sqlite"
	"github.com/golang-migrate/migrate/v4"
	mysqlDriver "github.com/golang-migrate/migrate/v4/database/mysql"
	sqliteDriver "github.com/golang-migrate/migrate/v4/database/sqlite"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"os"
)

var DB *gorm.DB

func InitDB() (*gorm.DB, error) {
	var err error

	godotenv.Load("../.env")

	DB, err = gorm.Open(getDB(), &gorm.Config{})
	if err != nil {
		log.Fatal("Error connecting to the database: ", err)
	}
	fmt.Println(DB.Name())
	runMigrations(DB)

	return DB, nil
}

func getDB() gorm.Dialector {

	if v, ok := os.LookupEnv("MYSQL_DSN"); ok {
		return mysql.Open(v)
	}

	// Fallback to PostgreSQL Testcontainer
	ctx := context.Background()
	dsn, cleanup, err := getTestContainer(ctx)
	if err != nil {
		panic(fmt.Sprintf("Failed to start PostgreSQL container: %v", err))
	}
	defer cleanup() // Ensure the container is cleaned up

	return postgres.Open(dsn)
}

func getTestContainer(ctx context.Context) (string, func(), error) {
	var env = map[string]string{
		"POSTGRES_PASSWORD": "root",
		"POSTGRES_USER":     "root",
		"POSTGRES_DB":       "my_database",
	}
	var port = "5432/tcp"
	// Define the PostgreSQL container request
	req := testcontainers.ContainerRequest{
		Image:        "postgres:15-alpine",
		ExposedPorts: []string{port},
		Env:          env,
		WaitingFor:   wait.ForLog("database system is ready to accept connections"),
	}

	// Start the PostgreSQL container
	container, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
	})

	if err != nil {
		return "", nil, fmt.Errorf("Failed to start PostgreSQL container: %v", err)
	}

	// Get the endpoint (host:port) of the PostgreSQL container
	endpoint, err := container.Endpoint(ctx, "")
	if err != nil {
		return "", nil, fmt.Errorf("Failed to get PostgreSQL container endpoint: %v", err)
	}

	// Construct the connection string for GORM
	dsn := fmt.Sprintf("host=%s port=%s user=root password=root dbname=my_database sslmode=disable",
		endpoint, "5432")

	cleanup := func() {
		if err := container.Terminate(ctx); err != nil {
			fmt.Printf("Failed to terminate PostgreSQL container: %v\n", err)
		}
	}
	return dsn, cleanup, nil
}

func runMigrations(db *gorm.DB) {
	sqlDB, _ := db.DB()

	var migration *migrate.Migrate
	var err error

	if db.Name() == "mysql" {
		driver, err := mysqlDriver.WithInstance(sqlDB, &mysqlDriver.Config{})
		if err != nil {
			log.Fatal("error with instance", err)
		}

		migration, err = migrate.NewWithDatabaseInstance(
			"file://../database/migrations",
			"mysql",
			driver,
		)

	} else {

		driver, err := sqliteDriver.WithInstance(sqlDB, &sqliteDriver.Config{})
		if err != nil {
			log.Fatal("error with instance", err)
		}

		migration, err = migrate.NewWithDatabaseInstance(
			"file://../database/migrations",
			"sqlite",
			driver,
		)

	}
	if err != nil {
		log.Fatal("error with migrations", err)
	}
	err = migration.Up()
	if err != nil {
		log.Fatal("error with run up scripts ", err)
	}
}
