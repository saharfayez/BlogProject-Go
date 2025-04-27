package models

type Role struct {
	ID          int          `gorm:"primary_key;auto_increment"`
	Name        string       `gorm:"unique;not null"`
	Permissions []Permission `gorm:"many2many:role_permissions;"`
}
