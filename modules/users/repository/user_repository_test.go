package repository_test

import (
	"errors"
	"testing"

	"github.com/Jardielson-s/api-task/modules/users/model"
	"github.com/Jardielson-s/api-task/modules/users/repository"
	"github.com/stretchr/testify/assert"
)

type MockDB struct {
	users  map[int]model.User
	nextID int
}

func NewMockDB() *MockDB {
	return &MockDB{
		users:  make(map[int]model.User),
		nextID: 1,
	}
}

func (db *MockDB) Create(user *model.User) error {
	user.ID = db.nextID
	db.users[db.nextID] = *user
	db.nextID++
	return nil
}

func (db *MockDB) Where(query string, args ...interface{}) *MockDB {
	return db
}

func (db *MockDB) First(user *model.User) error {
	for _, u := range db.users {
		if u.Email == user.Email {
			*user = u
			return nil
		}
	}
	return errors.New("user not found")
}

func (db *MockDB) Save(user *model.User) error {
	if _, exists := db.users[user.ID]; exists {
		db.users[user.ID] = *user
		return nil
	}
	return errors.New("user not found")
}

func (db *MockDB) Delete(user *model.User) error {
	if _, exists := db.users[user.ID]; exists {
		delete(db.users, user.ID)
		return nil
	}
	return errors.New("user not found")
}

type mockUserRepository struct {
	db *MockDB
}

func (r *mockUserRepository) FindById(id int) (model.User, error) {
	panic("unimplemented")
}

func (r *mockUserRepository) ListUsers(page int, pageSize int, searchQuery string) ([]model.User, int64, error) {
	panic("unimplemented")
}

func NewUserRepository(db *MockDB) repository.UserRepository {
	return &mockUserRepository{db}
}

func (r *mockUserRepository) Create(input *model.User) (model.User, error) {
	err := r.db.Create(input)
	return *input, err
}

func (r *mockUserRepository) FindByEmail(email string) (model.User, error) {
	var user model.User
	user.Email = email
	err := r.db.Where("email = ?", email).First(&user)
	return user, err
}

func (r *mockUserRepository) UpdateUser(id int, update model.User) (model.User, error) {
	update.ID = id
	err := r.db.Save(&update)
	return update, err
}

func (r *mockUserRepository) DeleteUser(id int) error {
	var user model.User
	user.ID = id
	return r.db.Delete(&user)
}

func TestCreate(t *testing.T) {
	mockDB := NewMockDB()
	repo := NewUserRepository(mockDB)

	user := &model.User{
		Username: "johndoe",
		Email:    "johndoe@example.com",
		Password: "securepassword",
	}

	createdUser, err := repo.Create(user)

	assert.NoError(t, err)
	assert.Equal(t, "johndoe", createdUser.Username)
	assert.Equal(t, "johndoe@example.com", createdUser.Email)
}

func TestFindByEmail(t *testing.T) {
	mockDB := NewMockDB()
	repo := NewUserRepository(mockDB)

	mockDB.Create(&model.User{Username: "johndoe", Email: "johndoe@example.com", Password: "securepassword"})

	user, err := repo.FindByEmail("johndoe@example.com")

	assert.NoError(t, err)
	assert.Equal(t, "johndoe", user.Username)
	assert.Equal(t, "johndoe@example.com", user.Email)
}

func TestUpdateUser(t *testing.T) {
	mockDB := NewMockDB()
	repo := NewUserRepository(mockDB)

	mockDB.Create(&model.User{Username: "johndoe", Email: "johndoe@example.com", Password: "securepassword"})

	updatedUser := model.User{Username: "johnupdated", Email: "johnupdated@example.com"}
	updatedUserResult, err := repo.UpdateUser(1, updatedUser)

	assert.NoError(t, err)
	assert.Equal(t, "johnupdated", updatedUserResult.Username)
	assert.Equal(t, "johnupdated@example.com", updatedUserResult.Email)
}

func TestDeleteUser(t *testing.T) {
	mockDB := NewMockDB()
	repo := NewUserRepository(mockDB)

	mockDB.Create(&model.User{Username: "johndoe", Email: "johndoe@example.com", Password: "securepassword"})

	err := repo.DeleteUser(1)

	assert.NoError(t, err)

	_, err = repo.FindByEmail("johndoe1@example.com")
	assert.Equal(t, err.Error(), "user not found")

}
