package services

import (
	"errors"

	"github.com/Jardielson-s/api-task/internal/authenticate"
	userModel "github.com/Jardielson-s/api-task/modules/users/model"
	"github.com/Jardielson-s/api-task/modules/users/repository"
	"gorm.io/gorm"
)

type UserService interface {
	CreateUserService(user *userModel.User) (userModel.User, error)
	UpdateUserService(id int, update userModel.UpdateUser) (userModel.User, error)
}

type userService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) UserService {
	return &userService{repo}
}

func (s userService) CreateUserService(user *userModel.User) (userModel.User, error) {
	userAlreadyExists, err := s.repo.FindByEmail(user.Email)
	if err == nil {
		return userAlreadyExists, errors.New(`Email already exists.`)
	}

	hash, err := authenticate.CreateHash(user.Password)
	if err != nil {
		return userAlreadyExists, err
	}
	user.Password = hash
	result, err := s.repo.Create(user)
	return result, err
}

func (s userService) UpdateUserService(id int, update userModel.UpdateUser) (userModel.User, error) {
	user, err := s.repo.FindById(id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return user, errors.New(`user not found`)
		}
		return user, errors.New(`error retrieving user`)

	}

	userAlreadyExists, err := s.repo.FindByEmail(update.Email)
	if err == nil {
		return userAlreadyExists, errors.New(`email already exists`)
	}

	user.Email = update.Email
	if update.Password != nil {
		hash, err := authenticate.CreateHash(*update.Password)
		if err != nil {
			return user, errors.New(`error to update user`)
		}
		user.Password = hash
	}
	result, err := s.repo.UpdateUser(id, user)
	return result, err
}
