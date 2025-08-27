package model

import (
	"time"

	roleModel "github.com/Jardielson-s/api-task/modules/roles/model"
	"github.com/Jardielson-s/api-task/modules/users/model"
	"gorm.io/gorm"
)

type UserRoles struct {
	*gorm.Model
	UserId    int `gorm:"not null"`
	RoleId    int `gorm:"not null"`
	CreatedAt time.Time
	DeletedAt gorm.DeletedAt
	User      model.User     `gorm:"foreignKey:UserId;references:ID"`
	Role      roleModel.Role `gorm:"foreignKey:RoleId;references:ID"`
}
