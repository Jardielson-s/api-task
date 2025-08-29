package services

import (
	"errors"
	"testing"

	"github.com/Jardielson-s/api-task/modules/tasks/model"
	"github.com/Jardielson-s/api-task/modules/tasks/repository"
	mock_repository "github.com/Jardielson-s/api-task/modules/tasks/repository/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func TestCreateTaskService(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mock_repository.NewMockTaskRepository(ctrl)
	taskService := NewTaskService(mockRepo)

	taskInput := &model.Task{Name: "New Task", Summary: "Some summary"}
	taskOutput := model.Task{ID: 1, Name: "New Task", Summary: "Some summary"}
	existingTask := model.Task{ID: 2, Name: "Existing Task", Summary: "Existing summary"}

	tests := []struct {
		name          string
		inputTask     *model.Task
		setupMock     func()
		expectedTask  model.Task
		expectedError error
	}{
		{
			name:      "Criação de tarefa com sucesso",
			inputTask: taskInput,
			setupMock: func() {
				mockRepo.EXPECT().FindByName(taskInput.Name).Return(model.Task{}, gorm.ErrRecordNotFound).Times(1)
				mockRepo.EXPECT().Create(taskInput).Return(taskOutput, nil).Times(1)
			},
			expectedTask:  taskOutput,
			expectedError: nil,
		},
		{
			name:      "Falha na criação: tarefa já existe",
			inputTask: &existingTask,
			setupMock: func() {
				mockRepo.EXPECT().FindByName(existingTask.Name).Return(existingTask, nil).Times(1)
			},
			expectedTask:  existingTask,
			expectedError: errors.New(`Task already exists.`),
		},
		{
			name:      "Falha na criação: erro no repositório",
			inputTask: taskInput,
			setupMock: func() {
				mockRepo.EXPECT().FindByName(taskInput.Name).Return(model.Task{}, gorm.ErrRecordNotFound).Times(1)
				mockRepo.EXPECT().Create(taskInput).Return(model.Task{}, errors.New("database error")).Times(1)
			},
			expectedTask:  model.Task{},
			expectedError: errors.New("database error"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.setupMock()
			result, err := taskService.CreateTaskService(tt.inputTask)
			assert.Equal(t, tt.expectedTask, result)
			if tt.expectedError != nil {
				assert.EqualError(t, err, tt.expectedError.Error())
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestUpdateTaskService(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mock_repository.NewMockTaskRepository(ctrl)
	taskService := NewTaskService(mockRepo)

	taskID := 1
	userID := 10
	taskFound := model.Task{ID: 1, Name: "Old Name", Summary: "Old Summary", Status: "pending", UserId: userID}

	stringPtr := func(s string) *string { return &s }

	tests := []struct {
		name          string
		id            int
		update        model.TaskUpdate
		userID        *int
		setupMock     func()
		expectedTask  model.Task
		expectedError error
	}{

		{
			name:   "Falha na atualização: tarefa não encontrada",
			id:     taskID,
			update: model.TaskUpdate{},
			userID: &userID,
			setupMock: func() {
				mockRepo.EXPECT().FindByQuery(repository.Query{ID: taskID, UserId: &userID}).Return(model.Task{}, gorm.ErrRecordNotFound).Times(1)
			},
			expectedTask:  model.Task{},
			expectedError: errors.New(`task not found`),
		},
		{
			name:   "Falha na atualização: erro ao buscar a tarefa",
			id:     taskID,
			update: model.TaskUpdate{},
			userID: &userID,
			setupMock: func() {
				mockRepo.EXPECT().FindByQuery(repository.Query{ID: taskID, UserId: &userID}).Return(model.Task{}, errors.New("database error")).Times(1)
			},
			expectedTask:  model.Task{},
			expectedError: errors.New(`error retrieving task`),
		},
		{
			name:   "Falha na atualização: novo nome já existe",
			id:     taskID,
			update: model.TaskUpdate{Name: stringPtr("Existing Name")},
			userID: &userID,
			setupMock: func() {
				mockRepo.EXPECT().FindByQuery(repository.Query{ID: taskID, UserId: &userID}).Return(taskFound, nil).Times(1)
				mockRepo.EXPECT().FindByName("Existing Name").Return(model.Task{}, nil).Times(1)
			},
			expectedTask:  model.Task{},
			expectedError: errors.New(`task already exists`),
		},
		{
			name:   "Atualização: nenhum campo para atualizar",
			id:     taskID,
			update: model.TaskUpdate{},
			userID: &userID,
			setupMock: func() {
				mockRepo.EXPECT().FindByQuery(repository.Query{ID: taskID, UserId: &userID}).Return(taskFound, nil).Times(1)
				mockRepo.EXPECT().UpdateTask(taskFound).Return(taskFound, nil).Times(1)
			},
			expectedTask:  taskFound,
			expectedError: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.setupMock()
			result, err := taskService.UpdateTaskService(tt.id, tt.update, tt.userID)

			assert.Equal(t, tt.expectedTask, result)
			if tt.expectedError != nil {
				assert.EqualError(t, err, tt.expectedError.Error())
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
