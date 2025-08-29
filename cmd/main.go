package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/Jardielson-s/api-task/cmd/migrations"
	"github.com/Jardielson-s/api-task/cmd/seeders"
	"github.com/Jardielson-s/api-task/configs"
	"github.com/Jardielson-s/api-task/infra"
	"github.com/Jardielson-s/api-task/infra/aws/sqs"
	"github.com/Jardielson-s/api-task/modules/auth"
	"github.com/Jardielson-s/api-task/modules/tasks"
	"github.com/Jardielson-s/api-task/modules/users"
	httpSwagger "github.com/swaggo/http-swagger"

	_ "github.com/Jardielson-s/api-task/docs"
	_ "github.com/go-sql-driver/mysql"
)

// @title			TASKS-API REST API
// @version		1.0
// @description	This is a simple Go REST API using MySQL and Swagger for documentation.
// @host			localhost:8080
// @SecurityDefinitions.apiKey Bearer
// @in header
// @name Authorization
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
	auth.AuthRoutes(mux, db)
	tasks.TaskRoutes(mux, db)

	mux.Handle("/swagger/", httpSwagger.WrapHandler)
	client, _ := sqs.CreateSQSClient()
	if os.Getenv("ACTIVE_NOTIFICATION") == "true" {
		go sqs.ProcessMessages(client, db)
	}
	log.Println(fmt.Sprint("Server has started in: ", os.Getenv("PORT")))
	http.ListenAndServe(fmt.Sprint(":", os.Getenv("PORT")), mux)
}
