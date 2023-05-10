package config

import (
	"dao/database"
	"message/littlebee"
	"message/querycommand"
	"util/queryclause"

	"github.com/fatih/structs"
)

func CreateDeviceType(l *littlebee.DeviceType) (*littlebee.DeviceType, error) {
	result := database.DB.Omit("DeviceTypeId", "UpdatedAt", "UpdatedBy", "IsDeleted", "DeletedAt", "DeletedBy").Create(l)
	return l, result.Error
}

func RetrieveDeviceType(deviceTypeId int32) (*littlebee.DeviceType, error) {
	r := &littlebee.DeviceType{}
	result := database.DB.First(&r, deviceTypeId)
	return r, result.Error
}

func UpdateDeviceType(l *littlebee.DeviceType) error {
	lMap := structs.Map(l)
	lMap["created_at"] = l.CreatedAt.AsTime().UTC()
	lMap["updated_at"] = l.UpdatedAt.AsTime().UTC()
	lMap["deleted_at"] = l.DeletedAt.AsTime().UTC()
	result := database.DB.Model(l).Omit("DeviceTypeId", "CreatedAt", "CreatedBy", "IsDeleted", "DeletedAt", "DeletedBy").Updates(lMap)
	return result.Error
}

func DeleteDeviceTypeWithFlag(l *littlebee.DeviceType) error {
	result := database.DB.Model(l).Updates(littlebee.DeviceType{IsDeleted: l.IsDeleted, DeletedAt: l.DeletedAt, DeletedBy: l.DeletedBy})
	return result.Error
}

func DeleteDeviceType(l *littlebee.DeviceType) error {
	result := database.DB.Delete(l)
	return result.Error
}

func QueryDeviceType(q *querycommand.QueryCommand) (*littlebee.DeviceTypeList, error) {
	var deviceTypes []*littlebee.DeviceType
	s := "SELECT * FROM device_types "
	w, v := queryclause.QueryCommandString(q)
	s = s + w
	result := database.DB.Raw(s, v...).Scan(&deviceTypes)
	return &littlebee.DeviceTypeList{Items: deviceTypes}, result.Error
}

func GetAllDeviceTypes() (*littlebee.DeviceTypeList, error) {
	var deviceTypes []*littlebee.DeviceType
	result := database.DB.Find(&deviceTypes)
	return &littlebee.DeviceTypeList{Items: deviceTypes}, result.Error
}
