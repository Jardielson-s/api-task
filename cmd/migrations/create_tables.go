package migrations

import (
	"fmt"

	permissionModel "github.com/Jardielson-s/api-task/modules/permissions/model"
	rolePermissionsModel "github.com/Jardielson-s/api-task/modules/role_permissions/model"
	roleModel "github.com/Jardielson-s/api-task/modules/roles/model"
	taskModel "github.com/Jardielson-s/api-task/modules/tasks/model"
	userRoles "github.com/Jardielson-s/api-task/modules/user_roles/model"
	userModel "github.com/Jardielson-s/api-task/modules/users/model"

	"gorm.io/gorm"
)

type IndexInfo struct {
	ColumnName string `gorm:"column:Column_name"`
	KeyName    string `gorm:"column:Key_name"`
	NonUnique  int    `gorm:"column:Non_unique"` // 1 se o índice não for único
}

func createIndex(db *gorm.DB, tableName string, name string, keys string) {
	var indexes []IndexInfo
	db.Raw(fmt.Sprintf(`SHOW INDEX FROM %s where key_name ='%s';`, tableName, name)).Scan(&indexes)
	if len(indexes) == 0 {
		db.Exec(fmt.Sprintf("CREATE UNIQUE INDEX %s ON %s(%s)", name, tableName, keys))
	}
}

func RunMigrates(db *gorm.DB) {
	db.AutoMigrate(&userModel.User{})
	db.AutoMigrate(&roleModel.Role{})
	db.AutoMigrate(&permissionModel.Permission{})
	db.AutoMigrate(&rolePermissionsModel.RolePermissions{})
	db.AutoMigrate(&userRoles.UserRoles{})
	db.AutoMigrate(&taskModel.Task{})
	createIndex(db, "role_permissions", "idx_role_permission", "role_id, permission_id")
	createIndex(db, "user_roles", "idx_users_roles", "role_id, user_id")
	createIndex(db, "tasks", "idx_tasks", "id, user_id")
}
