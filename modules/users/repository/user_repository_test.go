package repository

// import (
// 	"errors"
// 	"testing"

// 	entity "github.com/Jardielson-s/api-task/modules/users/entities"
// 	"github.com/stretchr/testify/assert"
// 	"github.com/stretchr/testify/mock"
// )

// // Mockando a interface TaskRepository
// type MockTaskRepository struct {
// 	mock.Mock
// }

// func (m *MockTaskRepository) GetAll() ([]entity.User, error) {
// 	args := m.Called()
// 	return args.Get(0).([]entity.User), args.Error(1)
// }

// func (m *MockTaskRepository) Create(task *entity.User) error {
// 	args := m.Called(task)
// 	return args.Error(0)
// }

// func TestTaskRepository_GetAll(t *testing.T) {
// 	mockRepo := new(MockTaskRepository)

// 	// Simulando que a função GetAll retorna um erro
// 	mockRepo.On("GetAll").Return([]entity.User{}, nil)

// 	tasks, err := mockRepo.GetAll()

// 	// Verifica se o comportamento está correto
// 	assert.NoError(t, err)
// 	assert.Len(t, tasks, 0)

// 	// Verifica se a função foi chamada
// 	mockRepo.AssertExpectations(t)
// }

// func TestTaskRepository_GetAll_Error(t *testing.T) {
// 	mockRepo := new(MockTaskRepository)

// 	// Simulando erro ao pegar as tarefas
// 	mockRepo.On("GetAll").Return(nil, errors.New("database error"))

// 	tasks, err := mockRepo.GetAll()

// 	assert.Error(t, err)
// 	assert.Nil(t, tasks)

// 	mockRepo.AssertExpectations(t)
// }
