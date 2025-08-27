package model

import (
	"time"

	permissionModel "github.com/Jardielson-s/api-task/modules/permissions/model"
	roleModel "github.com/Jardielson-s/api-task/modules/roles/model"

	"gorm.io/gorm"
)

type RolePermissions struct {
	*gorm.Model
	RoleId       int `gorm:"not null"`
	PermissionId int `gorm:"not null"`
	CreatedAt    time.Time
	DeletedAt    gorm.DeletedAt
	Role         roleModel.Role             `gorm:"foreignKey:RoleId;references:ID"`
	Permission   permissionModel.Permission `gorm:"foreignKey:PermissionId;references:ID"`
}
