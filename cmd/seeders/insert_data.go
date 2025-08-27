package seeders

import (
	"github.com/Jardielson-s/api-task/internal/authenticate"
	permissionModel "github.com/Jardielson-s/api-task/modules/permissions/model"
	rolePermissionsModel "github.com/Jardielson-s/api-task/modules/role_permissions/model"
	"github.com/Jardielson-s/api-task/modules/roles/model"
	roleModel "github.com/Jardielson-s/api-task/modules/roles/model"
	userModel "github.com/Jardielson-s/api-task/modules/users/model"

	userRolesModel "github.com/Jardielson-s/api-task/modules/user_roles/model"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func InsertData01(db *gorm.DB) {

	hash, _ := authenticate.CreateHash("password123")
	var users = []userModel.User{
		{Username: "manager", Email: "manager@company.com", Password: hash},
		{Username: "tech1", Email: "tech1@company.com", Password: hash},
	}
	db.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "email"}},
		DoUpdates: clause.AssignmentColumns([]string{"email"}),
	}).CreateInBatches(users, len(users))

	var roles = []model.Role{
		{Name: "Manager"},
		{Name: "Technician"},
	}
	db.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "name"}},
		DoUpdates: clause.AssignmentColumns([]string{"name"}),
	}).CreateInBatches(roles, len(roles))

	var permissions = []permissionModel.Permission{
		{Name: "create"},
		{Name: "update"},
		{Name: "read"},
		{Name: "delete"},
		{Name: "received notifications"},
	}
	db.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "name"}},
		DoUpdates: clause.AssignmentColumns([]string{"name"}),
	}).CreateInBatches(permissions, len(permissions))
}

func InsertData02(db *gorm.DB) {
	var rolesManager []roleModel.Role
	db.Where("name = ?", "Manager").Model(&roleModel.Role{}).Find(&rolesManager)
	var rolesTechnician []roleModel.Role
	db.Where("name = ?", "Technician").Model(&roleModel.Role{}).Find(&rolesTechnician)
	var permissions []permissionModel.Permission
	db.Where(
		"name IN ?",
		[]interface{}{"create", "update", "read", "received notifications", "delete"},
	).Model(&permissionModel.Permission{}).Find(&permissions)

	var rolePermissions []rolePermissionsModel.RolePermissions
	for _, roles := range rolesManager {
		for _, permission := range permissions {
			rolePermissions = append(
				rolePermissions,
				rolePermissionsModel.RolePermissions{
					RoleId:       roles.ID,
					PermissionId: permission.ID,
				},
			)
		}
	}
	excludeNames := map[string]bool{
		"delete":                 true,
		"received notifications": true,
	}
	for _, roles := range rolesTechnician {
		for _, permission := range permissions {
			if _, found := excludeNames[permission.Name]; !found {
				rolePermissions = append(
					rolePermissions,
					rolePermissionsModel.RolePermissions{
						RoleId:       roles.ID,
						PermissionId: permission.ID,
					},
				)
			}
		}
	}
	db.Clauses(clause.OnConflict{
		DoUpdates: clause.AssignmentColumns([]string{"role_id", "permission_id"}),
	}).CreateInBatches(rolePermissions, len(rolePermissions))
}

func InsertData03(db *gorm.DB) {
	var users []userModel.User
	db.Where("email = ?", "manager@company.com").Model(&userModel.User{}).Find(&users)

	var rolesManager []roleModel.Role
	db.Where("name = ?", "Manager").Model(&roleModel.Role{}).Find(&rolesManager)

	var userRoles []userRolesModel.UserRoles
	for _, user := range users {
		for _, role := range rolesManager {
			userRoles = append(
				userRoles,
				userRolesModel.UserRoles{
					RoleId: role.ID,
					UserId: user.ID,
				},
			)
		}
	}
	db.Clauses(clause.OnConflict{
		DoUpdates: clause.AssignmentColumns([]string{"role_id", "user_id"}),
	}).CreateInBatches(userRoles, len(userRoles))
}

func RunSeeders(db *gorm.DB) {
	InsertData01(db)
	InsertData02(db)
	InsertData03(db)
}
