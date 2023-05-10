package config

import (
	dao "dao/config"
	"message/littlebee"
	"message/querycommand"

	"google.golang.org/protobuf/types/known/timestamppb"
)

func CreateDeviceType(l *littlebee.DeviceType, userId int32) (*littlebee.DeviceType, error) {
	l.DeviceTypeId = 0
	l.CreatedAt = timestamppb.Now()
	l.CreatedBy = userId
	return dao.CreateDeviceType(l)
}

func RetrieveDeviceType(deviceTypeId int32) (*littlebee.DeviceType, error) {
	return dao.RetrieveDeviceType(deviceTypeId)
}

func UpdateDeviceType(l *littlebee.DeviceType, userId int32) error {
	l.UpdatedAt = timestamppb.Now()
	l.UpdatedBy = userId
	return dao.UpdateDeviceType(l)
}

func DeleteDeviceTypeWithFlag(deviceTypeId int, userId int32) error {
	l := &littlebee.DeviceType{
		DeviceTypeId: int32(deviceTypeId),
		IsDeleted:     true,
		DeletedAt:     timestamppb.Now(),
		DeletedBy:     userId,
	}
	return dao.DeleteDeviceTypeWithFlag(l)
}

func DeleteDeviceType(deviceTypeId int, userId int32) error {
	l := &littlebee.DeviceType{
		DeviceTypeId: int32(deviceTypeId),
		IsDeleted:     true,
		DeletedAt:     timestamppb.Now(),
		DeletedBy:     userId,
	}
	return dao.DeleteDeviceType(l)
}

func QueryDeviceType(q *querycommand.QueryCommand) (*littlebee.DeviceTypeList, error) {
	return dao.QueryDeviceType(q)
}

func GetAllDeviceTypes() (*littlebee.DeviceTypeList, error) {
	return dao.GetAllDeviceTypes()
}