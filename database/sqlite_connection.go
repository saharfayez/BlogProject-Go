package database

import (
	"fmt"
	"github.com/glebarez/sqlite"
	"goproject/models"
	"gorm.io/gorm"
	"log"
)

var SqliteDB *gorm.DB

func InitSqliteDB() (*gorm.DB, error) {
	var err error
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})

	if err != nil {
		log.Fatal("Error connecting to the database: ", err)
	}

	SqliteDB = db
	db.AutoMigrate(&models.User{}, &models.Post{})
	if db.Migrator().HasTable(&models.User{}) {
		fmt.Println("Users table exists!")
	}

	return db, nil
}
