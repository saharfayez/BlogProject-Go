package database

import (
	"github.com/glebarez/sqlite"
	"goproject/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"os"
)

var DB *gorm.DB

func InitDB() (*gorm.DB, error) {
	var err error

	DB, err = gorm.Open(getDB(), &gorm.Config{})
	if err != nil {
		log.Fatal("Error connecting to the database: ", err)
	}

	DB.AutoMigrate(&models.User{}, &models.Post{})
	return DB, nil
}

func getDB() gorm.Dialector {

	if v, ok := os.LookupEnv("MYSQL_DSN"); ok {
		return mysql.Open(v)
	}

	return sqlite.Open(":memory:")
}
