package main

import (
	"log"
	"net/http"

	"github.com/Jardielson-s/api-task/cmd/configs"
	"github.com/Jardielson-s/api-task/cmd/migrations"
	"github.com/Jardielson-s/api-task/cmd/seeders"
	"github.com/Jardielson-s/api-task/infra"
	"github.com/Jardielson-s/api-task/modules/users"
	httpSwagger "github.com/swaggo/http-swagger"

	_ "github.com/Jardielson-s/api-task/docs"
	_ "github.com/go-sql-driver/mysql"
	/*adicionar essa linha */)

// func init() {
// 	viper.SetConfigFile(`config.json`)
// 	err := viper.ReadInConfig()
// 	if err != nil {
// 		panic(err)
// 	}
// }

// @title			TASKS-API REST API
// @version		1.0
// @description	This is a simple Go REST API using MySQL and Swagger for documentation.
// @host			localhost:8080
// @BasePath		/
func main() {

	configs.Envs()
	db, _ := infra.InitInfraDb()
	// run migrates
	migrations.RunMigrates(db)
	// seed the tables
	seeders.RunSeeders(db)

	mux := http.NewServeMux()
	users.UserRoutes(mux, db)

	// authService := authService.NewAuthService(userRepo)
	mux.Handle("/swagger/", httpSwagger.WrapHandler)
	// auth := AuthHandlers.NewLoginHandler(authService, userRepo)
	// http.HandleFunc("/auth/login", auth.LoginHandler)
	//userHandler.CreateUserHandler
	//http.HandlerFunc()
	// http.Handle("/users", authenticate.ProtectedHandler(http.HandlerFunc(userHandler.CreateUserHandler)))

	log.Println("Server has started in: 8080")
	http.ListenAndServe(":8080", mux)
}
