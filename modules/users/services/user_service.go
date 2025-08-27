package services

import (
	"log"

	"github.com/Jardielson-s/api-task/internal/authenticate"
	userModel "github.com/Jardielson-s/api-task/modules/users/model"
	"github.com/Jardielson-s/api-task/modules/users/repository"
)

type UserService interface {
	CreateUserService(user *userModel.User) (userModel.User, error)
}

type userService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) UserService {
	return &userService{repo}
}

func (s userService) CreateUserService(user *userModel.User) (userModel.User, error) {
	hash, err := authenticate.CreateHash(user.Password)
	if err != nil {
		log.Fatal("Error to create password hash")
	}
	user.Password = hash
	result, err := s.repo.Create(user)
	return result, err
}
