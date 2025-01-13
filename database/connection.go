package database

import (
	"goproject/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

var DB *gorm.DB

func InitDB() (*gorm.DB, error) {
	var err error
	dsn := "abstract-programmer:example-password@tcp(127.0.0.1:3306)/BlogManager"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Error connecting to the database: ", err)
	}

	DB = db
	db.AutoMigrate(&models.User{}, &models.Post{})

	return db, nil
}
