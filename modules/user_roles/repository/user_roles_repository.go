package repository

import (
	"errors"

	"github.com/Jardielson-s/api-task/modules/user_roles/model"
	"gorm.io/gorm"
)

type UserRolesRepository interface {
	FindByUserId(id int) ([]model.UserRoles, error)
}
type userRolesRepository struct {
	db *gorm.DB
}

func (u userRolesRepository) FindByUserId(id int) ([]model.UserRoles, error) {
	var userRoles []model.UserRoles
	err := u.db.Where("user_id = ?", id).Preload("Role").Find(&userRoles).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return userRoles, errors.New(`User not found`)

		}
		return userRoles, errors.New(`Error retrieving user`)
	}

	return userRoles, nil
}

func NewUserRolesRepository(db *gorm.DB) UserRolesRepository {
	return userRolesRepository{db}
}
