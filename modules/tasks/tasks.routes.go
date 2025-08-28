package tasks

import (
	"net/http"

	"github.com/Jardielson-s/api-task/internal/authorizations"
	"github.com/Jardielson-s/api-task/modules/shared"
	"github.com/Jardielson-s/api-task/modules/tasks/handlers"
	"github.com/Jardielson-s/api-task/modules/tasks/repository"
	"github.com/Jardielson-s/api-task/modules/tasks/services"
	"gorm.io/gorm"
)

func TaskRoutes(httpMux *http.ServeMux, db *gorm.DB) *http.ServeMux {
	taskRepo := repository.NewTaskRepository(db)
	taskService := services.NewTaskService(taskRepo)
	taskHandler := handlers.NewTaskHandler(taskService, taskRepo)
	httpMux.Handle("/tasks", authorizations.ApplyMiddlewares([]string{shared.GetTechnicianRole()}, shared.GetCreatePermission(), taskHandler.CreateTaskHandler))
	httpMux.Handle("/tasks/list", authorizations.ApplyMiddlewares([]string{shared.GetTechnicianRole(), shared.GetManagerRole()}, shared.GetReadPermission(), taskHandler.ListTasksHandler))
	httpMux.Handle("/tasks/{id}", authorizations.ApplyMiddlewares([]string{shared.GetTechnicianRole(), shared.GetManagerRole()}, shared.GetReadPermission(), taskHandler.FindTaskById))
	httpMux.Handle("/tasks/update/{id}", authorizations.ApplyMiddlewares([]string{shared.GetManagerRole(), shared.GetTechnicianRole()}, shared.GetUpdatePermission(), taskHandler.UpdateTaskHandler))
	httpMux.Handle("/tasks/delete/{id}", authorizations.ApplyMiddlewares([]string{shared.GetManagerRole()}, shared.GetDeletePermission(), taskHandler.DeleteTask))

	// httpMux.Handle("/users/{id}", authorizations.ApplyMiddlewares([]string{shared.GetManagerRole()}, shared.GetReadPermission(), userHandler.FindUserById))
	// httpMux.Handle("/users/update/{id}", authorizations.ApplyMiddlewares([]string{shared.GetManagerRole()}, shared.GetWritePermission(), userHandler.UpdateUserHandler))
	// httpMux.Handle("/users/delete/{id}", authorizations.ApplyMiddlewares([]string{shared.GetManagerRole()}, shared.GetWritePermission(), userHandler.DeleteUser))

	return httpMux
}
