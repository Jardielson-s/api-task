package repository

import (
	"errors"

	"github.com/Jardielson-s/api-task/modules/tasks/model"
	"gorm.io/gorm"
)

type Query struct {
	ID     int     `json:"id"`
	UserId *int    `json:"userId"`
	Name   *string `json:"name"`
}
type TaskRepository interface {
	Create(input *model.Task) (model.Task, error)
	FindByName(email string) (model.Task, error)
	ListTasks(page, pageSize int, searchQuery string, userId *int) ([]model.Task, int64, error)
	FindById(id int) (model.Task, error)
	FindByQuery(query Query) (model.Task, error)
	UpdateTask(update model.Task) (model.Task, error)
	DeleteTask(id int) error
}
type taskRepository struct {
	db *gorm.DB
}

func (u taskRepository) Create(input *model.Task) (model.Task, error) {
	task := input
	err := u.db.Create(task).Error
	return *task, err
}

func (u taskRepository) FindByName(name string) (model.Task, error) {
	var tasks model.Task
	err := u.db.Where("name = ?", name).First(&tasks).Error
	if err != nil {
		return tasks, err
	}
	return tasks, nil
}

func (u taskRepository) ListTasks(page, pageSize int, searchQuery string, userId *int) ([]model.Task, int64, error) {
	var tasks []model.Task
	var totalCount int64

	offset := (page - 1) * pageSize
	query := u.db.Model(&model.Task{})
	if searchQuery != "" {
		query = query.Where("name LIKE ?", "%"+searchQuery+"%")
	}
	if userId != nil {
		query = query.Where("user_id LIKE ?", userId)
	}
	if err := query.Limit(pageSize).Offset(offset).Find(&tasks).Error; err != nil {
		return nil, 0, err
	}

	if err := query.Model(&model.Task{}).Count(&totalCount).Error; err != nil {
		return nil, 0, err
	}

	return tasks, totalCount, nil
}

func (u taskRepository) FindById(id int) (model.Task, error) {
	var task model.Task
	err := u.db.Where("id = ?", id).First(&task).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return task, errors.New(`task not found`)

		}
		return task, errors.New(`error retrieving task`)

	}
	if err != nil {
		return task, err
	}
	return task, nil
}

func (u taskRepository) FindByQuery(queryInput Query) (model.Task, error) {
	var task model.Task
	query := `id = ?` //u.db.Where("id = ?", queryInput.ID)
	queryPass := []interface{}{queryInput.ID}
	if queryInput.UserId != nil {
		query += ` AND user_id = ?`
		queryPass = append(queryPass, queryInput.UserId)
	}
	err := u.db.Where(query, queryPass...).First(&task).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return task, errors.New(`task not found`)

		}
		return task, errors.New(`error retrieving task`)

	}
	if err != nil {
		return task, err
	}
	return task, nil
}

func (u taskRepository) UpdateTask(update model.Task) (model.Task, error) {
	err := u.db.Save(&update).Error
	return update, err
}
func (u taskRepository) DeleteTask(id int) error {
	var task model.Task
	err := u.db.Where("id = ?", id).First(&task).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return errors.New(`task not found`)
		}
		return errors.New(`error retrieving task`)
	}
	err = u.db.Delete(&task).Error
	if err != nil {
		return err
	}
	return nil
}

func NewTaskRepository(db *gorm.DB) TaskRepository {
	return taskRepository{db}
}
