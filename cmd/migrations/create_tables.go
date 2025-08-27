package migrations

import (
	"fmt"

	permissionModel "github.com/Jardielson-s/api-task/modules/permissions/model"
	rolePermissionsModel "github.com/Jardielson-s/api-task/modules/role_permissions/model"
	roleModel "github.com/Jardielson-s/api-task/modules/roles/model"
	userRoles "github.com/Jardielson-s/api-task/modules/user_roles/model"
	userModel "github.com/Jardielson-s/api-task/modules/users/model"

	"gorm.io/gorm"
)

type IndexInfo struct {
	// Essa estrutura pode ter mais ou menos campos dependendo das colunas do resultado
	ColumnName string `gorm:"column:Column_name"`
	KeyName    string `gorm:"column:Key_name"`
	NonUnique  int    `gorm:"column:Non_unique"` // 1 se o índice não for único
}

func RunMigrates(db *gorm.DB) {
	db.AutoMigrate(&userModel.User{})
	db.AutoMigrate(&roleModel.Role{})
	db.AutoMigrate(&permissionModel.Permission{})
	db.AutoMigrate(&rolePermissionsModel.RolePermissions{})
	db.AutoMigrate(&userRoles.UserRoles{})

	var indexes []IndexInfo
	db.Raw(`SHOW INDEX FROM role_permissions where key_name ='idx_role_permission';`).Scan(&indexes)
	if len(indexes) == 0 {
		db.Exec("CREATE UNIQUE INDEX idx_role_permission ON role_permissions(role_id, permission_id)")
	} else {
		fmt.Println("Índice único composto já existe!")
	}
}
