package users

// import (
// 	"net/http"

// 	"github.com/Jardielson-s/api-task/internal/authenticate"
// 	"github.com/Jardielson-s/api-task/modules/users/handlers"
// 	"github.com/Jardielson-s/api-task/modules/users/repository"
// 	"github.com/Jardielson-s/api-task/modules/users/services"
// 	"gorm.io/gorm"
// )

// func UserModule(db *gorm.DB) http {
// 	hp := http
// 	userRepo := repository.NewUserRepository(db)
// 	userService := services.NewUserService(userRepo)
// 	userHandler := handlers.NewUserHandler(userService)
// 	// hp.Handle("/swagger/", httpSwagger.WrapHandler)
// 	//userHandler.CreateUserHandler
// 	//http.HandlerFunc()
// 	http.Handle("/users", authenticate.ProtectedHandler(http.HandlerFunc(userHandler.CreateUserHandler)))
// 	return http

// }
