package config

import (
	dao "dao/config"
	"message/littlebee"
	"message/querycommand"

	"google.golang.org/protobuf/types/known/timestamppb"
)

func CreateDriver(l *littlebee.Driver, userId int32) (*littlebee.Driver, error) {
	l.DriverId = 0
	l.CreatedAt = timestamppb.Now()
	l.CreatedBy = userId
	return dao.CreateDriver(l)
}

func RetrieveDriver(driverId int32) (*littlebee.Driver, error) {
	return dao.RetrieveDriver(driverId)
}

func UpdateDriver(l *littlebee.Driver, userId int32) error {
	l.UpdatedAt = timestamppb.Now()
	l.UpdatedBy = userId
	return dao.UpdateDriver(l)
}

func DeleteDriverWithFlag(driverId int, userId int32) error {
	l := &littlebee.Driver{
		DriverId: int32(driverId),
		IsDeleted:     true,
		DeletedAt:     timestamppb.Now(),
		DeletedBy:     userId,
	}
	return dao.DeleteDriverWithFlag(l)
}

func DeleteDriver(driverId int, userId int32) error {
	l := &littlebee.Driver{
		DriverId: int32(driverId),
		IsDeleted:     true,
		DeletedAt:     timestamppb.Now(),
		DeletedBy:     userId,
	}
	return dao.DeleteDriver(l)
}

func QueryDriver(q *querycommand.QueryCommand) (*littlebee.DriverList, error) {
	return dao.QueryDriver(q)
}

func GetAllDrivers() (*littlebee.DriverList, error) {
	return dao.GetAllDrivers()
}