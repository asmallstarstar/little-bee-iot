package config

import (
	dao "dao/config"
	"message/littlebee"
	"message/querycommand"

	"google.golang.org/protobuf/types/known/timestamppb"
)

func CreateRole(l *littlebee.Role, userId int32) (*littlebee.Role, error) {
	l.RoleId = 0
	l.CreatedAt = timestamppb.Now()
	l.CreatedBy = userId
	return dao.CreateRole(l)
}

func RetrieveRole(roleId int32) (*littlebee.Role, error) {
	return dao.RetrieveRole(roleId)
}

func UpdateRole(l *littlebee.Role, userId int32) error {
	l.UpdatedAt = timestamppb.Now()
	l.UpdatedBy = userId
	return dao.UpdateRole(l)
}

func DeleteRoleWithFlag(roleId int, userId int32) error {
	l := &littlebee.Role{
		RoleId: int32(roleId),
		IsDeleted:     true,
		DeletedAt:     timestamppb.Now(),
		DeletedBy:     userId,
	}
	return dao.DeleteRoleWithFlag(l)
}

func DeleteRole(roleId int, userId int32) error {
	l := &littlebee.Role{
		RoleId: int32(roleId),
		IsDeleted:     true,
		DeletedAt:     timestamppb.Now(),
		DeletedBy:     userId,
	}
	return dao.DeleteRole(l)
}

func QueryRole(q *querycommand.QueryCommand) (*littlebee.RoleList, error) {
	return dao.QueryRole(q)
}

func GetAllRoles() (*littlebee.RoleList, error) {
	return dao.GetAllRoles()
}