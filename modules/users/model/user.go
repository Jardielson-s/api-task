package model

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	*gorm.Model
	ID        int    `gorm:"primaryKey"`
	Username  string `gorm:"username"`
	Email     string `gorm:"index:,unique"`
	Password  string `gorm:"password"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}
