package repository

import (
	"fmt"

	userModel "github.com/Jardielson-s/api-task/modules/users/model"

	"gorm.io/gorm"
)

type UserRepository interface {
	Create(input *userModel.User) (userModel.User, error)
	FindByEmail(email string) (userModel.User, error)
}
type userRepository struct {
	db *gorm.DB
}

func (u userRepository) Create(input *userModel.User) (userModel.User, error) {
	user := &userModel.User{Username: input.Username, Email: input.Email, Password: input.Password}
	err := u.db.Create(user).Error
	return *user, err
}

func (u userRepository) FindByEmail(email string) (userModel.User, error) {

	var user userModel.User
	err := u.db.Where("email = ?", email).First(&user).Error
	fmt.Print(user)
	if err != nil {
		// log.Fatalln("Password or email wrong")
		return user, err
	}
	return user, nil
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return userRepository{db}
}
