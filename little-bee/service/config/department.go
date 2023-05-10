package config

import (
	dao "dao/config"
	"message/littlebee"
	"message/querycommand"

	"google.golang.org/protobuf/types/known/timestamppb"
)

func CreateDepartment(l *littlebee.Department, userId int32) (*littlebee.Department, error) {
	l.DepartmentId = 0
	l.CreatedAt = timestamppb.Now()
	l.CreatedBy = userId
	return dao.CreateDepartment(l)
}

func RetrieveDepartment(departmentId int32) (*littlebee.Department, error) {
	return dao.RetrieveDepartment(departmentId)
}

func UpdateDepartment(l *littlebee.Department, userId int32) error {
	l.UpdatedAt = timestamppb.Now()
	l.UpdatedBy = userId
	return dao.UpdateDepartment(l)
}

func DeleteDepartmentWithFlag(departmentId int, userId int32) error {
	l := &littlebee.Department{
		DepartmentId: int32(departmentId),
		IsDeleted:     true,
		DeletedAt:     timestamppb.Now(),
		DeletedBy:     userId,
	}
	return dao.DeleteDepartmentWithFlag(l)
}

func DeleteDepartment(departmentId int, userId int32) error {
	l := &littlebee.Department{
		DepartmentId: int32(departmentId),
		IsDeleted:     true,
		DeletedAt:     timestamppb.Now(),
		DeletedBy:     userId,
	}
	return dao.DeleteDepartment(l)
}

func QueryDepartment(q *querycommand.QueryCommand) (*littlebee.DepartmentList, error) {
	return dao.QueryDepartment(q)
}

func GetAllDepartments() (*littlebee.DepartmentList, error) {
	return dao.GetAllDepartments()
}