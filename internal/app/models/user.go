package models

type User struct {
	ID       int    `gorm:"primaryKey;autoIncrement"`
	Username string `gorm:"unique;not null" json:"username"`
	Password string `gorm:"not null" json:"password"`
	Roles    []Role `gorm:"many2many:user_roles;" json:"roles"`
	Posts    []Post `gorm:"foreignKey:UserID" json:"posts"`
}
