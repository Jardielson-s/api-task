package repository

import (
	"github.com/Jardielson-s/api-task/modules/role_permissions/model"
	"gorm.io/gorm"
)

type RolePermissionsRepository interface {
	FindByRoleIds(input []int) ([]model.RolePermissions, error)
}
type rolePermissionsRepository struct {
	db *gorm.DB
}

func NewPermissionRepository(db *gorm.DB) RolePermissionsRepository {
	return rolePermissionsRepository{db}
}

func (r rolePermissionsRepository) FindByRoleIds(input []int) ([]model.RolePermissions, error) {
	var rolePermission []model.RolePermissions
	err := r.db.Where("role_id IN ?", input).Preload("Permission").Find(&rolePermission).Error
	if err != nil {
		return rolePermission, err
	}
	return rolePermission, nil
}
