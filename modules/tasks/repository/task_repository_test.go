package repository

import (
	"errors"
	"testing"

	"github.com/Jardielson-s/api-task/modules/tasks/model"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func setupTestDB() (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	db.AutoMigrate(&model.Task{})
	return db, nil
}

func TestCreateTask(t *testing.T) {
	db, err := setupTestDB()
	assert.NoError(t, err)

	repo := NewTaskRepository(db)

	inputTask := model.Task{Name: "New Task", UserId: 1, Summary: "summary"}
	createdTask, err := repo.Create(&inputTask)

	assert.NoError(t, err)
	assert.Equal(t, "New Task", createdTask.Name)
	assert.Equal(t, 1, createdTask.UserId)
}

func TestFindByName(t *testing.T) {
	db, err := setupTestDB()
	assert.NoError(t, err)
	repo := NewTaskRepository(db)

	task := model.Task{Name: "Task", UserId: 1, Summary: "Summary"}
	db.Create(&task)

	foundTask, err := repo.FindByName("Task")
	assert.NoError(t, err)
	assert.Equal(t, "Task", foundTask.Name)

	_, err = repo.FindByName("task not found")
	assert.Error(t, err)
	assert.True(t, errors.Is(err, gorm.ErrRecordNotFound))
}

func TestListTasks(t *testing.T) {
	db, err := setupTestDB()
	assert.NoError(t, err)
	repo := NewTaskRepository(db)

	db.Create(&model.Task{Name: "Task 1", UserId: 1, Summary: "Summary", Status: "pending"})
	db.Create(&model.Task{Name: "Task 2", UserId: 2, Summary: "Summary", Status: "pending"})
	db.Create(&model.Task{Name: "Task Filter", UserId: 1, Summary: "Summary", Status: "pending"})

	tasks, count, err := repo.ListTasks(1, 10, "", nil)
	assert.NoError(t, err)
	assert.Len(t, tasks, 3)
	assert.Equal(t, int64(3), count)

	tasks, count, err = repo.ListTasks(1, 10, "Filter", nil)
	assert.NoError(t, err)
	assert.Len(t, tasks, 1)
	assert.Equal(t, int64(1), count)
	assert.Equal(t, "Task Filter", tasks[0].Name)

	userID := 1
	tasks, count, err = repo.ListTasks(1, 10, "", &userID)
	assert.NoError(t, err)
	assert.Len(t, tasks, 2)
	assert.Equal(t, int64(2), count)
}

func TestFindById(t *testing.T) {
	db, err := setupTestDB()
	assert.NoError(t, err)
	repo := NewTaskRepository(db)

	task := model.Task{Name: "Task ID", UserId: 1, Summary: "Summary", Status: "pending"}
	db.Create(&task)

	foundTask, err := repo.FindById(int(task.ID))
	assert.NoError(t, err)
	assert.Equal(t, "Task ID", foundTask.Name)

	_, err = repo.FindById(999)
	assert.Error(t, err)
	assert.EqualError(t, err, "task not found")
}

func TestUpdateTask(t *testing.T) {
	db, err := setupTestDB()
	assert.NoError(t, err)
	repo := NewTaskRepository(db)

	task := model.Task{Name: "Tarefa para Atualizar", UserId: 1, Summary: "Summary"}
	db.Create(&task)

	updateData := model.Task{Name: "Task updated", ID: int(task.ID), UserId: 1, Summary: "test summary", Status: "pending"}
	updatedTask, err := repo.UpdateTask(int(task.ID), updateData)

	assert.NoError(t, err)
	assert.Equal(t, "Task updated", updatedTask.Name)

	var checkTask model.Task
	db.First(&checkTask, int(task.ID))
	assert.Equal(t, "Task updated", checkTask.Name)
}

func TestDeleteTask(t *testing.T) {
	db, err := setupTestDB()
	assert.NoError(t, err)
	repo := NewTaskRepository(db)

	task := model.Task{Name: "Task to delete", UserId: 1, Summary: "Summary Test", Status: "pending"}
	db.Create(&task)

	err = repo.DeleteTask(int(task.ID))
	assert.NoError(t, err)

	var deletedTask model.Task
	result := db.First(&deletedTask, int(task.ID))
	assert.True(t, errors.Is(result.Error, gorm.ErrRecordNotFound))

	err = repo.DeleteTask(999)
	assert.Error(t, err)
	assert.EqualError(t, err, "task not found")
}
