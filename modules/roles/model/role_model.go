package model

import "gorm.io/gorm"

type Role struct {
	*gorm.Model
	ID        int    `gorm:"primaryKey"`
	Name      string `json:"name"`
	CreatedAt []uint8
	UpdatedAt []uint8
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
