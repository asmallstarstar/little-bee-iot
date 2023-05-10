package config

import (
	"dao/database"
	"message/littlebee"
	"message/querycommand"
	"util/queryclause"

	"github.com/fatih/structs"
)

func CreateDepartment(l *littlebee.Department) (*littlebee.Department, error) {
	result := database.DB.Omit("DepartmentId", "UpdatedAt", "UpdatedBy", "IsDeleted", "DeletedAt", "DeletedBy").Create(l)
	return l, result.Error
}

func RetrieveDepartment(departmentId int32) (*littlebee.Department, error) {
	r := &littlebee.Department{}
	result := database.DB.First(&r, departmentId)
	return r, result.Error
}

func UpdateDepartment(l *littlebee.Department) error {
	lMap := structs.Map(l)
	lMap["created_at"] = l.CreatedAt.AsTime().UTC()
	lMap["updated_at"] = l.UpdatedAt.AsTime().UTC()
	lMap["deleted_at"] = l.DeletedAt.AsTime().UTC()
	result := database.DB.Model(l).Omit("DepartmentId", "CreatedAt", "CreatedBy", "IsDeleted", "DeletedAt", "DeletedBy").Updates(lMap)
	return result.Error
}

func DeleteDepartmentWithFlag(l *littlebee.Department) error {
	result := database.DB.Model(l).Updates(littlebee.Department{IsDeleted: l.IsDeleted, DeletedAt: l.DeletedAt, DeletedBy: l.DeletedBy})
	return result.Error
}

func DeleteDepartment(l *littlebee.Department) error {
	result := database.DB.Delete(l)
	return result.Error
}

func QueryDepartment(q *querycommand.QueryCommand) (*littlebee.DepartmentList, error) {
	var departments []*littlebee.Department
	s := "SELECT * FROM departments "
	w, v := queryclause.QueryCommandString(q)
	s = s + w
	result := database.DB.Raw(s, v...).Scan(&departments)
	return &littlebee.DepartmentList{Items: departments}, result.Error
}

func GetAllDepartments() (*littlebee.DepartmentList, error) {
	var departments []*littlebee.Department
	result := database.DB.Find(&departments)
	return &littlebee.DepartmentList{Items: departments}, result.Error
}
