package repository

import (
	"gorm.io/gorm"
)

type PermissionRepository interface {
}
type permissionRepository struct {
	db *gorm.DB
}

func NewPermissionRepository(db *gorm.DB) permissionRepository {
	return permissionRepository{db}
}
