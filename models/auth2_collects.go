package models

import "time"

type User2Collects struct {
	UserID       uint         `gorm:"primaryKey"`
	UserModel    UserModel    `gorm:"foreignKey:AuthID"`
	ArticleID    uint         `gorm:"primaryKey"`
	ArticleModel ArticleModel `gorm:"foreignKey:ArticleID"`
	CreatedAt    time.Time
}
