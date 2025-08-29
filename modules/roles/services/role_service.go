package services

import (
	"github.com/Jardielson-s/api-task/modules/roles/repositories"
	UserRolesRepo "github.com/Jardielson-s/api-task/modules/user_roles/repository"
)

type RolesService interface {
	FindByRoleByName(name string) ([]string, error)
}
type rolesService struct {
	roleRepo      repositories.RolesRepository
	usersRoleRepo UserRolesRepo.UserRolesRepository
}

func (s rolesService) FindByRoleByName(name string) ([]string, error) {
	var emails []string

	role, err := s.roleRepo.FindByRoleByName(name)
	if err != nil {
		return emails, nil
	}

	userRoles, err := s.usersRoleRepo.FindEmailsByRoleId(role.ID)
	if err != nil {
		return emails, nil
	}
	for _, userRole := range userRoles {
		emails = append(emails, userRole.User.Email)
	}

	return emails, nil
}

func NewRolesRepository(
	roleRepo repositories.RolesRepository,
	usersRoleRepo UserRolesRepo.UserRolesRepository,
) RolesService {
	return rolesService{roleRepo, usersRoleRepo}
}
