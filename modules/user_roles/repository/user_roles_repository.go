package repository

import (
	"errors"

	"github.com/Jardielson-s/api-task/modules/user_roles/model"
	"gorm.io/gorm"
)

type UserRolesRepository interface {
	FindByUserId(id int) ([]model.UserRoles, error)
	LinkUserWithRole(userId int, roleId int) error
	FindEmailsByRoleId(roleId int) ([]model.UserRoles, error)
}
type userRolesRepository struct {
	db *gorm.DB
}

func (u userRolesRepository) FindByUserId(id int) ([]model.UserRoles, error) {
	var userRoles []model.UserRoles
	err := u.db.Where("user_id = ?", id).Preload("Role").Find(&userRoles).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return userRoles, errors.New(`user not found`)

		}
		return userRoles, errors.New(`error retrieving user`)
	}

	return userRoles, nil
}

func (u userRolesRepository) FindEmailsByRoleId(roleId int) ([]model.UserRoles, error) {
	var userRoles []model.UserRoles
	err := u.db.Where("role_id = ?", roleId).Preload("User").Find(&userRoles).Error
	if err != nil {
		return userRoles, errors.New(`error retrieving user`)
	}
	return userRoles, nil
}

func (u userRolesRepository) LinkUserWithRole(userId int, roleId int) error {

	err := u.db.Save(&model.UserRoles{
		RoleId: roleId,
		UserId: userId,
	}).Error
	if err != nil {
		return errors.New(`user has linked with role`)
	}
	return nil
}

func NewUserRolesRepository(db *gorm.DB) UserRolesRepository {
	return userRolesRepository{db}
}
