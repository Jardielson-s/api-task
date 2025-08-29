package repository

import (
	"errors"
	"testing"
	"time"

	"github.com/Jardielson-s/api-task/modules/users/model"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// A mock struct for userModel.User, based on the provided code snippet
type User struct {
	ID        int `gorm:"primaryKey"`
	Username  string
	Email     string `gorm:"uniqueIndex"`
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}

func (User) TableName() string {
	return "users"
}

// setupTestDB initializes an in-memory SQLite database for testing
func setupTestDB() (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	db.AutoMigrate(&User{})
	return db, nil
}

func TestCreate(t *testing.T) {
	db, err := setupTestDB()
	assert.NoError(t, err)

	repo := NewUserRepository(db)

	inputUser := model.User{Username: "techuser", Email: "tech@example.com", Password: "password123"}
	createdUser, err := repo.Create(&inputUser)

	assert.NoError(t, err)
	assert.Equal(t, "techuser", createdUser.Username)
	assert.Equal(t, "tech@example.com", createdUser.Email)
	assert.NotEqual(t, 0, createdUser.ID)
}

func TestFindByEmail(t *testing.T) {
	db, err := setupTestDB()
	assert.NoError(t, err)
	repo := NewUserRepository(db)

	user := &User{Username: "findbyemailuser", Email: "find@example.com", Password: "password123"}
	db.Create(user)

	foundUser, err := repo.FindByEmail("find@example.com")
	assert.NoError(t, err)
	assert.Equal(t, "findbyemailuser", foundUser.Username)
	assert.Equal(t, "find@example.com", foundUser.Email)

	_, err = repo.FindByEmail("nonexistent@example.com")
	assert.Error(t, err)
	assert.True(t, errors.Is(err, gorm.ErrRecordNotFound))
}

func TestFindById(t *testing.T) {
	db, err := setupTestDB()
	assert.NoError(t, err)
	repo := NewUserRepository(db)

	user := &User{Username: "findbyiduser", Email: "id@example.com"}
	db.Create(user)

	foundUser, err := repo.FindById(user.ID)
	assert.NoError(t, err)
	assert.Equal(t, user.ID, foundUser.ID)
	assert.Equal(t, "findbyiduser", foundUser.Username)

	_, err = repo.FindById(999)
	assert.Error(t, err)
	assert.EqualError(t, err, "user not found")
}

func TestUpdateUser(t *testing.T) {
	db, err := setupTestDB()
	assert.NoError(t, err)
	repo := NewUserRepository(db)

	user := &User{Username: "olduser", Email: "old@example.com"}
	db.Create(user)

	updatedData := model.User{ID: user.ID, Username: "updateduser", Email: "updated@example.com"}
	updatedUser, err := repo.UpdateUser(user.ID, updatedData)

	assert.NoError(t, err)
	assert.Equal(t, "updateduser", updatedUser.Username)
	assert.Equal(t, "updated@example.com", updatedUser.Email)

	var persistedUser User
	db.First(&persistedUser, user.ID)
	assert.Equal(t, "updateduser", persistedUser.Username)
}

func TestDeleteUser(t *testing.T) {
	db, err := setupTestDB()
	assert.NoError(t, err)
	repo := NewUserRepository(db)

	user := &User{Username: "deleteuser", Email: "delete@example.com"}
	db.Create(user)

	err = repo.DeleteUser(user.ID)
	assert.NoError(t, err)

	var deletedUser User
	result := db.First(&deletedUser, user.ID)
	assert.True(t, errors.Is(result.Error, gorm.ErrRecordNotFound))

	err = repo.DeleteUser(999)
	assert.Error(t, err)
	assert.EqualError(t, err, "user not found")
}
