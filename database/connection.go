package database

import (
	"fmt"
	"github.com/joho/godotenv"

	//"github.com/glebarez/sqlite"
	"github.com/golang-migrate/migrate/v4"
	mysqlDriver "github.com/golang-migrate/migrate/v4/database/mysql"
	sqliteDriver "github.com/golang-migrate/migrate/v4/database/sqlite"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"gorm.io/driver/mysql"
	"gorm.io/driver/sqlite"
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

	return sqlite.Open(":memory:")
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
