package config

import (
	"dao/database"
	"message/littlebee"
	"message/querycommand"
	"util/queryclause"

	"github.com/fatih/structs"
)

func CreateLogicDevice(l *littlebee.LogicDevice) (*littlebee.LogicDevice, error) {
	result := database.DB.Create(l)
	return l, result.Error
}

func RetrieveLogicDevice(logicDeviceId int32) (*littlebee.LogicDevice, error) {
	r := &littlebee.LogicDevice{}
	result := database.DB.First(&r, logicDeviceId)
	return r, result.Error
}

func UpdateLogicDevice(l *littlebee.LogicDevice) error {
	lMap := structs.Map(l)
	result := database.DB.Model(l).Omit("LogicDeviceId").Updates(lMap)
	return result.Error
}

func DeleteLogicDevice(l *littlebee.LogicDevice) error {
	result := database.DB.Delete(l)
	return result.Error
}

func QueryLogicDevice(q *querycommand.QueryCommand) (*littlebee.LogicDeviceList, error) {
	var logicDevices []*littlebee.LogicDevice
	s := "SELECT * FROM logic_devices "
	w, v := queryclause.QueryCommandString(q)
	s = s + w
	result := database.DB.Raw(s, v...).Scan(&logicDevices)
	return &littlebee.LogicDeviceList{Items: logicDevices}, result.Error
}

func GetAllLogicDevices() (*littlebee.LogicDeviceList, error) {
	var logicDevices []*littlebee.LogicDevice
	result := database.DB.Find(&logicDevices)
	return &littlebee.LogicDeviceList{Items: logicDevices}, result.Error
}
