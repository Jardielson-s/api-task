package services

import (
	"errors"

	"github.com/Jardielson-s/api-task/internal/authenticate"
	entity "github.com/Jardielson-s/api-task/modules/auth/entities"
	"github.com/Jardielson-s/api-task/modules/users/repository"
)

type AuthService interface {
	Login(user *entity.Login) (string, error)
}

type authService struct {
	repo repository.UserRepository
}

func NewAuthService(repo repository.UserRepository) AuthService {
	return &authService{repo}
}

func (s authService) Login(login *entity.Login) (string, error) {
	user, err := s.repo.FindByEmail(login.Email)
	if err != nil {
		return string(user.ID), nil
	}

	if !authenticate.CompareHash(login.Password, user.Password) {
		return "", errors.New(`Email or password incorrect`)
	}
	token, err := authenticate.CreateToken(user.Username)
	if err != nil {
		return "", err
	}
	return token, nil
}
