package config

import (
	dao "dao/config"
	"message/littlebee"
	"message/querycommand"

	"google.golang.org/protobuf/types/known/timestamppb"
)

func CreateDriverType(l *littlebee.DriverType, userId int32) (*littlebee.DriverType, error) {
	l.DriverTypeId = 0
	l.CreatedAt = timestamppb.Now()
	l.CreatedBy = userId
	return dao.CreateDriverType(l)
}

func RetrieveDriverType(driverTypeId int32) (*littlebee.DriverType, error) {
	return dao.RetrieveDriverType(driverTypeId)
}

func UpdateDriverType(l *littlebee.DriverType, userId int32) error {
	l.UpdatedAt = timestamppb.Now()
	l.UpdatedBy = userId
	return dao.UpdateDriverType(l)
}

func DeleteDriverTypeWithFlag(driverTypeId int, userId int32) error {
	l := &littlebee.DriverType{
		DriverTypeId: int32(driverTypeId),
		IsDeleted:     true,
		DeletedAt:     timestamppb.Now(),
		DeletedBy:     userId,
	}
	return dao.DeleteDriverTypeWithFlag(l)
}

func DeleteDriverType(driverTypeId int, userId int32) error {
	l := &littlebee.DriverType{
		DriverTypeId: int32(driverTypeId),
		IsDeleted:     true,
		DeletedAt:     timestamppb.Now(),
		DeletedBy:     userId,
	}
	return dao.DeleteDriverType(l)
}

func QueryDriverType(q *querycommand.QueryCommand) (*littlebee.DriverTypeList, error) {
	return dao.QueryDriverType(q)
}

func GetAllDriverTypes() (*littlebee.DriverTypeList, error) {
	return dao.GetAllDriverTypes()
}