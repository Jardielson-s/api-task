package services

import (
	"log"

	"github.com/Jardielson-s/api-task/internal/authenticate"
	entity "github.com/Jardielson-s/api-task/modules/users/entities"
	"github.com/Jardielson-s/api-task/modules/users/repository"
)

type UserService interface {
	CreateUserService(user *entity.User) (entity.User, error)
}

type userService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) UserService {
	return &userService{repo}
}

func (s userService) CreateUserService(user *entity.User) (entity.User, error) {
	hash, err := authenticate.CreateHash(user.Password)
	if err != nil {
		log.Fatal("Error to create password hash")
	}
	user.Password = hash
	result, err := s.repo.Create(user)
	return result, err
}
