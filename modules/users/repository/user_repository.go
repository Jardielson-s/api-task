package repository

import (
	"errors"

	"github.com/Jardielson-s/api-task/modules/users/model"
	userModel "github.com/Jardielson-s/api-task/modules/users/model"

	"gorm.io/gorm"
)

type UserRepository interface {
	Create(input *userModel.User) (userModel.User, error)
	FindByEmail(email string) (userModel.User, error)
	ListUsers(page, pageSize int, searchQuery string) ([]userModel.User, int64, error)
	FindById(id int) (userModel.User, error)
	UpdateUser(id int, update userModel.User) (userModel.User, error)
	DeleteUser(id int) error
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
	if err != nil {
		return user, err
	}
	return user, nil
}

func (u userRepository) ListUsers(page, pageSize int, searchQuery string) ([]userModel.User, int64, error) {
	var users []userModel.User
	var totalCount int64

	offset := (page - 1) * pageSize
	query := u.db.Model(&userModel.User{})
	if searchQuery != "" {
		// Adicionando filtro de pesquisa no username e email
		query = query.Where("username LIKE ? OR email LIKE ?", "%"+searchQuery+"%", "%"+searchQuery+"%")
	}
	if err := query.Limit(pageSize).Offset(offset).Find(&users).Error; err != nil {
		return nil, 0, err
	}

	if err := query.Model(&userModel.User{}).Count(&totalCount).Error; err != nil {
		return nil, 0, err
	}

	return users, totalCount, nil
}

func (u userRepository) FindById(id int) (userModel.User, error) {
	var user userModel.User
	err := u.db.Where("id = ?", id).First(&user).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return user, errors.New(`User not found`)

		}
		return user, errors.New(`Error retrieving user`)

	}
	if err != nil {
		return user, err
	}
	return user, nil
}

func (u userRepository) UpdateUser(id int, update model.User) (userModel.User, error) {
	u.db.Save(&update)
	return update, nil
}
func (u userRepository) DeleteUser(id int) error {
	var user userModel.User
	err := u.db.Where("id = ?", id).First(&user).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return errors.New(`User not found`)
		}
		return errors.New(`Error retrieving user`)
	}
	err = u.db.Delete(&user).Error
	if err != nil {
		return err
	}
	return nil
}

// db.Where("age = ?", 20).Delete(&User{})
func NewUserRepository(db *gorm.DB) UserRepository {
	return userRepository{db}
}
