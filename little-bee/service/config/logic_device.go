package config

import (
	dao "dao/config"
	"message/littlebee"
	"message/querycommand"
)

func CreateLogicDevice(l *littlebee.LogicDevice, userId int32) (*littlebee.LogicDevice, error) {
	return dao.CreateLogicDevice(l)
}

func RetrieveLogicDevice(logicDeviceId int32) (*littlebee.LogicDevice, error) {
	return dao.RetrieveLogicDevice(logicDeviceId)
}

func UpdateLogicDevice(l *littlebee.LogicDevice, userId int32) error {
	return dao.UpdateLogicDevice(l)
}

func DeleteLogicDevice(logicDeviceId int, userId int32) error {
	l := &littlebee.LogicDevice{
		LogicObjectId: int32(logicDeviceId),
	}
	return dao.DeleteLogicDevice(l)
}

func QueryLogicDevice(q *querycommand.QueryCommand) (*littlebee.LogicDeviceList, error) {
	return dao.QueryLogicDevice(q)
}

func GetAllLogicDevices() (*littlebee.LogicDeviceList, error) {
	return dao.GetAllLogicDevices()
}
