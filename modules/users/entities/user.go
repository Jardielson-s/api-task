package entity

import (
	"gorm.io/gorm"
)

type User struct {
	*gorm.Model
	ID        int    `gorm:"primaryKey"`
	Username  string `json:"username"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	CreatedAt []uint8
	UpdatedAt []uint8
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
