package repositories

import (
	"github.com/Jardielson-s/api-task/modules/roles/model"
	"gorm.io/gorm"
)

type RolesRepository interface {
	FindByRoleByName(name string) (model.Role, error)
}
type rolesRepository struct {
	db *gorm.DB
}

func (u rolesRepository) FindByRoleByName(name string) (model.Role, error) {
	var role model.Role
	err := u.db.Where("name = ?", name).First(&role).Error
	if err != nil {
		return role, nil
	}
	return role, nil
}

func NewRolesRepository(db *gorm.DB) RolesRepository {
	return rolesRepository{db}
}
