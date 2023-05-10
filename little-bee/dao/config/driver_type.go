package config

import (
	"dao/database"
	"message/littlebee"
	"message/querycommand"
	"util/queryclause"

	"github.com/fatih/structs"
)

func CreateDriverType(l *littlebee.DriverType) (*littlebee.DriverType, error) {
	result := database.DB.Omit("DriverTypeId", "UpdatedAt", "UpdatedBy", "IsDeleted", "DeletedAt", "DeletedBy").Create(l)
	return l, result.Error
}

func RetrieveDriverType(driverTypeId int32) (*littlebee.DriverType, error) {
	r := &littlebee.DriverType{}
	result := database.DB.First(&r, driverTypeId)
	return r, result.Error
}

func UpdateDriverType(l *littlebee.DriverType) error {
	lMap := structs.Map(l)
	lMap["created_at"] = l.CreatedAt.AsTime().UTC()
	lMap["updated_at"] = l.UpdatedAt.AsTime().UTC()
	lMap["deleted_at"] = l.DeletedAt.AsTime().UTC()
	result := database.DB.Model(l).Omit("DriverTypeId", "CreatedAt", "CreatedBy", "IsDeleted", "DeletedAt", "DeletedBy").Updates(lMap)
	return result.Error
}

func DeleteDriverTypeWithFlag(l *littlebee.DriverType) error {
	result := database.DB.Model(l).Updates(littlebee.DriverType{IsDeleted: l.IsDeleted, DeletedAt: l.DeletedAt, DeletedBy: l.DeletedBy})
	return result.Error
}

func DeleteDriverType(l *littlebee.DriverType) error {
	result := database.DB.Delete(l)
	return result.Error
}

func QueryDriverType(q *querycommand.QueryCommand) (*littlebee.DriverTypeList, error) {
	var driverTypes []*littlebee.DriverType
	s := "SELECT * FROM driver_types "
	w, v := queryclause.QueryCommandString(q)
	s = s + w
	result := database.DB.Raw(s, v...).Scan(&driverTypes)
	return &littlebee.DriverTypeList{Items: driverTypes}, result.Error
}

func GetAllDriverTypes() (*littlebee.DriverTypeList, error) {
	var driverTypes []*littlebee.DriverType
	result := database.DB.Find(&driverTypes)
	return &littlebee.DriverTypeList{Items: driverTypes}, result.Error
}
