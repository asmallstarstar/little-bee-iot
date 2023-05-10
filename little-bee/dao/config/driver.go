package config

import (
	"dao/database"
	"message/littlebee"
	"message/querycommand"
	"util/queryclause"

	"github.com/fatih/structs"
)

func CreateDriver(l *littlebee.Driver) (*littlebee.Driver, error) {
	result := database.DB.Omit("DriverId", "UpdatedAt", "UpdatedBy", "IsDeleted", "DeletedAt", "DeletedBy").Create(l)
	return l, result.Error
}

func RetrieveDriver(driverId int32) (*littlebee.Driver, error) {
	r := &littlebee.Driver{}
	result := database.DB.First(&r, driverId)
	return r, result.Error
}

func UpdateDriver(l *littlebee.Driver) error {
	lMap := structs.Map(l)
	lMap["created_at"] = l.CreatedAt.AsTime().UTC()
	lMap["updated_at"] = l.UpdatedAt.AsTime().UTC()
	lMap["deleted_at"] = l.DeletedAt.AsTime().UTC()
	result := database.DB.Model(l).Omit("DriverId", "CreatedAt", "CreatedBy", "IsDeleted", "DeletedAt", "DeletedBy").Updates(lMap)
	return result.Error
}

func DeleteDriverWithFlag(l *littlebee.Driver) error {
	result := database.DB.Model(l).Updates(littlebee.Driver{IsDeleted: l.IsDeleted, DeletedAt: l.DeletedAt, DeletedBy: l.DeletedBy})
	return result.Error
}

func DeleteDriver(l *littlebee.Driver) error {
	result := database.DB.Delete(l)
	return result.Error
}

func QueryDriver(q *querycommand.QueryCommand) (*littlebee.DriverList, error) {
	var drivers []*littlebee.Driver
	s := "SELECT * FROM drivers "
	w, v := queryclause.QueryCommandString(q)
	s = s + w
	result := database.DB.Raw(s, v...).Scan(&drivers)
	return &littlebee.DriverList{Items: drivers}, result.Error
}

func GetAllDrivers() (*littlebee.DriverList, error) {
	var drivers []*littlebee.Driver
	result := database.DB.Find(&drivers)
	return &littlebee.DriverList{Items: drivers}, result.Error
}
