package services

import (
	"errors"

	"github.com/Jardielson-s/api-task/modules/tasks/model"
	"github.com/Jardielson-s/api-task/modules/tasks/repository"
	"gorm.io/gorm"
)

type TaskService interface {
	CreateTaskService(user *model.Task) (model.Task, error)
	UpdateTaskService(id int, update model.TaskUpdate, userId *int) (model.Task, error)
}

type taskService struct {
	repo repository.TaskRepository
}

func NewTaskService(repo repository.TaskRepository) TaskService {
	return &taskService{repo}
}

func (s taskService) CreateTaskService(task *model.Task) (model.Task, error) {
	taskAlreadyExists, err := s.repo.FindByName(task.Name)
	if err == nil {
		return taskAlreadyExists, errors.New(`Task already exists.`)
	}
	result, err := s.repo.Create(task)
	return result, err
}

func (s taskService) UpdateTaskService(id int, update model.TaskUpdate, userId *int) (model.Task, error) {
	task, err := s.repo.FindByQuery(repository.Query{
		ID:     id,
		UserId: userId,
	})
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return task, errors.New(`Task not found`)
		}
		return task, errors.New(`Error retrieving task`)
	}
	taskAlreadyExists, err := s.repo.FindByName(*update.Name)
	if err == nil {
		return taskAlreadyExists, errors.New(`Task already exists.`)
	}
	if update.Status != nil {
		task.Status = *update.Status
	}
	if update.Summary != nil {
		task.Summary = *update.Summary
	}
	result, err := s.repo.UpdateTask(id, task)
	return result, err
}
