package config

import (
	"dao/database"
	"message/littlebee"
	"message/querycommand"
	"util/queryclause"

	"github.com/fatih/structs"
)

func CreateRole(l *littlebee.Role) (*littlebee.Role, error) {
	result := database.DB.Omit("RoleId", "UpdatedAt", "UpdatedBy", "IsDeleted", "DeletedAt", "DeletedBy").Create(l)
	return l, result.Error
}

func RetrieveRole(roleId int32) (*littlebee.Role, error) {
	r := &littlebee.Role{}
	result := database.DB.First(&r, roleId)
	return r, result.Error
}

func UpdateRole(l *littlebee.Role) error {
	lMap := structs.Map(l)
	lMap["created_at"] = l.CreatedAt.AsTime().UTC()
	lMap["updated_at"] = l.UpdatedAt.AsTime().UTC()
	lMap["deleted_at"] = l.DeletedAt.AsTime().UTC()
	result := database.DB.Model(l).Omit("RoleId", "CreatedAt", "CreatedBy", "IsDeleted", "DeletedAt", "DeletedBy").Updates(lMap)
	return result.Error
}

func DeleteRoleWithFlag(l *littlebee.Role) error {
	result := database.DB.Model(l).Updates(littlebee.Role{IsDeleted: l.IsDeleted, DeletedAt: l.DeletedAt, DeletedBy: l.DeletedBy})
	return result.Error
}

func DeleteRole(l *littlebee.Role) error {
	result := database.DB.Delete(l)
	return result.Error
}

func QueryRole(q *querycommand.QueryCommand) (*littlebee.RoleList, error) {
	var roles []*littlebee.Role
	s := "SELECT * FROM roles "
	w, v := queryclause.QueryCommandString(q)
	s = s + w
	result := database.DB.Raw(s, v...).Scan(&roles)
	return &littlebee.RoleList{Items: roles}, result.Error
}

func GetAllRoles() (*littlebee.RoleList, error) {
	var roles []*littlebee.Role
	result := database.DB.Find(&roles)
	return &littlebee.RoleList{Items: roles}, result.Error
}
