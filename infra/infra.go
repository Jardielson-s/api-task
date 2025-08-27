package infra

import (
	"github.com/Jardielson-s/api-task/infra/database"
	"gorm.io/gorm"
)

func InitInfraDb() (*gorm.DB, error) {
	db, err := database.ConnectMysql()
	return db, err
}
