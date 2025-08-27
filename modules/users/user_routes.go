package users

import (
	"net/http"

	UserHandlers "github.com/Jardielson-s/api-task/modules/users/handlers"
	"github.com/Jardielson-s/api-task/modules/users/repository"
	"github.com/Jardielson-s/api-task/modules/users/services"
	"gorm.io/gorm"
)

func UserRoutes(http *http.ServeMux, db *gorm.DB) *http.ServeMux {
	// 	http.Handle("/users", authenticate.ProtectedHandler(http.HandlerFunc(userHandler.CreateUserHandler)))
	userRepo := repository.NewUserRepository(db)
	userService := services.NewUserService(userRepo)
	userHandler := UserHandlers.NewUserHandler(userService, userRepo)
	http.HandleFunc("/users/", userHandler.CreateUserHandler)
	http.HandleFunc("/users/list", userHandler.ListUsersHandler)
	http.HandleFunc("/users/{id}", userHandler.FindUserById)
	http.HandleFunc("/users/update/{id}", userHandler.UpdateUserHandler)
	http.HandleFunc("/users/delete/{id}", userHandler.DeleteUser)

	return http
}
