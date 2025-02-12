package users

import "goproject/posts"

type User struct {
	ID       int    `gorm:"primaryKey;autoIncrement"`
	Username string `gorm:"unique;not null" json:"username"`
	Password string `gorm:"not null" json:"password"`
	Posts    []posts.Post `gorm:"foreignKey:UserID" json:"posts"`
}
