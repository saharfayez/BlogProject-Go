package models

type User struct {
	ID       int    `gorm:"primaryKey;autoIncrement"`
	Username string `gorm:"unique;not null" json:"username"`
	Password string `gorm:"not null" json:"password"`
	Role     string `gorm:"not null" json:"role"`
	Posts    []Post `gorm:"foreignKey:UserID" json:"posts"`
}
