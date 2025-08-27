package model

import (
	"time"

	"gorm.io/gorm"
)

type Role struct {
	*gorm.Model
	ID        int    `gorm:"primaryKey"`
	Name      string `gorm:"index:,unique"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}
