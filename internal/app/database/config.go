package database

import (
	"context"
	"errors"
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	db_postgres "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	appcontext "goproject/internal/app/context"
	"goproject/test/utils"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
	"path/filepath"
)

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
		dsn, err := utils.StartTestContainer(ctx)
		if err != nil {
			panic(fmt.Sprintf("Failed to start PostgreSQL container: %v", err))
		}
		return postgres.Open(dsn)
	}
}

func runMigrations(db *gorm.DB) {
	sqlDB, _ := db.DB()

	var migration *migrate.Migrate
	var err error
	var migrationPath string

	currentDirectory := getCurrentDirectory()
	fmt.Println("current", currentDirectory)

	migrationPath = filepath.Join("file://", currentDirectory, "..", "..", "/internal/app/database/migrations")

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

func getCurrentDirectory() string {

	currentDirectory, err := os.Getwd()
	if err != nil {
		log.Fatal("Failed to get current directory", err)
	}

	return currentDirectory
}
