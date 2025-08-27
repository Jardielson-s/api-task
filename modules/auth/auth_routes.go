package auth

import (
	"net/http"

	"github.com/Jardielson-s/api-task/modules/user_roles/repository"
	UserRepository "github.com/Jardielson-s/api-task/modules/users/repository"

	RolePermissionsRepository "github.com/Jardielson-s/api-task/modules/role_permissions/repository"

	"github.com/Jardielson-s/api-task/modules/auth/handlers"
	"github.com/Jardielson-s/api-task/modules/auth/services"
	"gorm.io/gorm"
)

func AuthRoutes(http *http.ServeMux, db *gorm.DB) *http.ServeMux {
	userRepo := UserRepository.NewUserRepository(db)
	rolePermissions := RolePermissionsRepository.NewPermissionRepository(db)
	userRolesRepo := repository.NewUserRolesRepository(db)
	authService := services.NewAuthService(userRepo, &userRolesRepo, &rolePermissions)
	authHandler := handlers.NewLoginHandler(authService, userRepo)
	http.HandleFunc("/auth/login", authHandler.LoginHandler)

	return http
}
