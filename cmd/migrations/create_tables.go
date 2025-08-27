package migrations

import (
	"github.com/Jardielson-s/api-task/infra"
	"github.com/Jardielson-s/api-task/modules/roles/model"
	entity "github.com/Jardielson-s/api-task/modules/users/entities"
)

func MigrationsStart01() {
	db, _ := infra.InitInfraDb()
	db.AutoMigrate(&entity.User{}, &model.Role{})

	var users = []entity.User{
		{Username: "manager", Email: "manager@company.com", Password: "password123"},
		{Username: "tech1", Email: "tech1@company.com", Password: "password123"},
	}
	db.CreateInBatches(users, len(users))
	var roles = []model.Role{
		{Name: "manager"},
		{Name: "tech1"},
	}
	db.CreateInBatches(roles, len(roles))

}
