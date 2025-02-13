package models

type Post struct {
	ID      uint   `gorm:"primaryKey;autoIncrement"`
	Title   string `gorm:"not null" json:"title"`
	Content string `gorm:"not null" json:"content"`
	UserID  int    `gorm:"not null" json:"user_id"`
}
