package database

import (
	"errors"
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	db_postgres "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	appcontext "goproject/internal/app/context"
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

	RunMigrations(db)

	return db, nil
}

func getDB() gorm.Dialector {
	return postgres.Open(appcontext.Context.GetPropertiesConfig().GetDatabaseUrl())
}

func RunMigrations(db *gorm.DB) {
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
