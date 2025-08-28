package model

import (
	"time"

	"github.com/Jardielson-s/api-task/modules/users/model"
	"gorm.io/gorm"
)

type Task struct {
	*gorm.Model
	ID        int    `gorm:"primaryKey"`
	UserId    int    `gorm:"foreignKey:,unique"`
	Name      string `gorm:"index:,unique"`
	Summary   string `gorm:"summary"`
	Status    string `gorm:"status"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
	User      model.User `gorm:"foreignKey:UserId;references:ID"`
}

type TaskUpdate struct {
	Name    *string `json:"name"`
	Summary *string `json:"summary"`
	Status  *string `json:"status"`
}
