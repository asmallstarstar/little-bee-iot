package config

import (
	"dao/database"
	"message/littlebee"
	"message/querycommand"
	"util/queryclause"

	"github.com/fatih/structs"
)

func CreateSignalVirtual(l *littlebee.SignalVirtual) (*littlebee.SignalVirtual, error) {
	result := database.DB.Omit("SignalVirtualId", "UpdatedAt", "UpdatedBy", "IsDeleted", "DeletedAt", "DeletedBy").Create(l)
	return l, result.Error
}

func RetrieveSignalVirtual(signalVirtualId int32) (*littlebee.SignalVirtual, error) {
	r := &littlebee.SignalVirtual{}
	result := database.DB.First(&r, signalVirtualId)
	return r, result.Error
}

func UpdateSignalVirtual(l *littlebee.SignalVirtual) error {
	lMap := structs.Map(l)
	result := database.DB.Model(l).Omit("SignalAiId").Updates(lMap)
	return result.Error
}

func DeleteSignalVirtual(l *littlebee.SignalVirtual) error {
	result := database.DB.Delete(l)
	return result.Error
}

func QuerySignalVirtual(q *querycommand.QueryCommand) (*littlebee.SignalVirtualList, error) {
	var signalVirtuals []*littlebee.SignalVirtual
	s := "SELECT * FROM signal_virtuals "
	w, v := queryclause.QueryCommandString(q)
	s = s + w
	result := database.DB.Raw(s, v...).Scan(&signalVirtuals)
	return &littlebee.SignalVirtualList{Items: signalVirtuals}, result.Error
}

func GetAllSignalVirtuals() (*littlebee.SignalVirtualList, error) {
	var signalVirtuals []*littlebee.SignalVirtual
	result := database.DB.Find(&signalVirtuals)
	return &littlebee.SignalVirtualList{Items: signalVirtuals}, result.Error
}
