package repository

import (
	"fmt"

	entity "github.com/Jardielson-s/api-task/modules/users/entities"
	"gorm.io/gorm"
)

type UserRepository interface {
	Create(input *entity.User) (entity.User, error)
	FindByEmail(email string) (entity.User, error)
}
type userRepository struct {
	db *gorm.DB
}

func (u userRepository) Create(input *entity.User) (entity.User, error) {
	user := &entity.User{Username: input.Username, Email: input.Email, Password: input.Password}
	err := u.db.Create(user).Error
	return *user, err
}

func (u userRepository) FindByEmail(email string) (entity.User, error) {

	var user entity.User
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
