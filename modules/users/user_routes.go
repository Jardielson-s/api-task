package users

import (
	"net/http"

	"github.com/Jardielson-s/api-task/internal/authorizations"
	"github.com/Jardielson-s/api-task/modules/shared"
	userRolesModel "github.com/Jardielson-s/api-task/modules/user_roles/repository"
	UserHandlers "github.com/Jardielson-s/api-task/modules/users/handlers"
	"github.com/Jardielson-s/api-task/modules/users/repository"

	"github.com/Jardielson-s/api-task/modules/users/services"
	"gorm.io/gorm"
)

func UserRoutes(httpMux *http.ServeMux, db *gorm.DB) *http.ServeMux {
	userRepo := repository.NewUserRepository(db)
	userRolesRepo := userRolesModel.NewUserRolesRepository(db)
	userService := services.NewUserService(userRepo)
	userHandler := UserHandlers.NewUserHandler(userService, userRepo, userRolesRepo)
	httpMux.HandleFunc("/users", userHandler.CreateUserHandler)
	httpMux.Handle("/users/list", authorizations.ApplyMiddlewares([]string{shared.GetManagerRole()}, shared.GetReadPermission(), userHandler.ListUsersHandler))
	httpMux.Handle("/users/{id}", authorizations.ApplyMiddlewares([]string{shared.GetManagerRole()}, shared.GetReadPermission(), userHandler.FindUserById))
	httpMux.Handle("/users/update/{id}", authorizations.ApplyMiddlewares([]string{shared.GetManagerRole()}, shared.GetUpdatePermission(), userHandler.UpdateUserHandler))
	httpMux.Handle("/users/link", authorizations.ApplyMiddlewares([]string{shared.GetManagerRole()}, shared.GetUpdatePermission(), userHandler.LinkUserWithRoleHandler))
	httpMux.Handle("/users/delete/{id}", authorizations.ApplyMiddlewares([]string{shared.GetManagerRole()}, shared.GetDeletePermission(), userHandler.DeleteUser))

	return httpMux
}
