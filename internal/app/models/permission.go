package models

type Permission struct {
	ID   int    `gorm:"primary_key;auto_increment"`
	Name string `gorm:"unique;not null"`
}
