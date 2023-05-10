package config

import (
	"dao/database"
	"message/littlebee"
	"message/querycommand"
	"util/queryclause"

	"github.com/fatih/structs"
)

func CreateSignalAi(l *littlebee.SignalAi) (*littlebee.SignalAi, error) {
	result := database.DB.Omit("SignalAiId", "UpdatedAt", "UpdatedBy", "IsDeleted", "DeletedAt", "DeletedBy").Create(l)
	return l, result.Error
}

func RetrieveSignalAi(signalAiId int32) (*littlebee.SignalAi, error) {
	r := &littlebee.SignalAi{}
	result := database.DB.First(&r, signalAiId)
	return r, result.Error
}

func UpdateSignalAi(l *littlebee.SignalAi) error {
	lMap := structs.Map(l)
	result := database.DB.Model(l).Omit("SignalAiId").Updates(lMap)
	return result.Error
}

func DeleteSignalAi(l *littlebee.SignalAi) error {
	result := database.DB.Delete(l)
	return result.Error
}

func QuerySignalAi(q *querycommand.QueryCommand) (*littlebee.SignalAiList, error) {
	var signalAis []*littlebee.SignalAi
	s := "SELECT * FROM signal_ais "
	w, v := queryclause.QueryCommandString(q)
	s = s + w
	result := database.DB.Raw(s, v...).Scan(&signalAis)
	return &littlebee.SignalAiList{Items: signalAis}, result.Error
}

func GetAllSignalAis() (*littlebee.SignalAiList, error) {
	var signalAis []*littlebee.SignalAi
	result := database.DB.Find(&signalAis)
	return &littlebee.SignalAiList{Items: signalAis}, result.Error
}
