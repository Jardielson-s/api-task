package services

import (
	"errors"

	"github.com/Jardielson-s/api-task/internal/authenticate"
	entity "github.com/Jardielson-s/api-task/modules/auth/entities"
	RolePermissionsRepo "github.com/Jardielson-s/api-task/modules/role_permissions/repository"
	userRoles "github.com/Jardielson-s/api-task/modules/user_roles/repository"

	"github.com/Jardielson-s/api-task/modules/users/repository"
)

type AuthService interface {
	Login(user *entity.Login) (string, error)
}

type authService struct {
	repo                repository.UserRepository
	userRolesRepo       userRoles.UserRolesRepository
	rolePermissionsRepo RolePermissionsRepo.RolePermissionsRepository
}

func NewAuthService(
	repo repository.UserRepository,
	userRolesRepo *userRoles.UserRolesRepository,
	rolePermissionsRepo *RolePermissionsRepo.RolePermissionsRepository,
) AuthService {
	return &authService{repo, *userRolesRepo, *rolePermissionsRepo}
}

func (s authService) Login(login *entity.Login) (string, error) {
	user, err := s.repo.FindByEmail(login.Email)
	if err != nil {
		return string(user.ID), err
	}

	if !authenticate.CompareHash(login.Password, user.Password) {
		return "", errors.New(`email or password incorrect`)
	}
	roles := []interface{}{}
	var rolesIds []int
	result, err := s.userRolesRepo.FindByUserId(user.ID)

	if err != nil {
		return "", err
	}
	for _, userRole := range result {
		roles = append(roles, userRole.Role.Name)
		rolesIds = append(rolesIds, userRole.RoleId)
	}
	rolePermissions, err := s.rolePermissionsRepo.FindByRoleIds(rolesIds)
	if err != nil {
		return "", err
	}
	permissions := []interface{}{}
	for _, rolePermission := range rolePermissions {
		permissions = append(permissions, rolePermission.Permission.Name)
	}
	token, err := authenticate.CreateToken(authenticate.TokenInfo{
		ID:          user.ID,
		Username:    user.Username,
		Email:       user.Email,
		Roles:       roles,
		Permissions: permissions,
	})
	if err != nil {
		return "", err
	}
	return token, nil
}
