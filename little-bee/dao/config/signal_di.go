package config

import (
	"dao/database"
	"message/littlebee"
	"message/querycommand"
	"util/queryclause"

	"github.com/fatih/structs"
)

func CreateSignalDi(l *littlebee.SignalDi) (*littlebee.SignalDi, error) {
	result := database.DB.Omit("SignalDiId", "UpdatedAt", "UpdatedBy", "IsDeleted", "DeletedAt", "DeletedBy").Create(l)
	return l, result.Error
}

func RetrieveSignalDi(signalDiId int32) (*littlebee.SignalDi, error) {
	r := &littlebee.SignalDi{}
	result := database.DB.First(&r, signalDiId)
	return r, result.Error
}

func UpdateSignalDi(l *littlebee.SignalDi) error {
	lMap := structs.Map(l)
	result := database.DB.Model(l).Omit("SignalAiId").Updates(lMap)
	return result.Error
}

func DeleteSignalDi(l *littlebee.SignalDi) error {
	result := database.DB.Delete(l)
	return result.Error
}

func QuerySignalDi(q *querycommand.QueryCommand) (*littlebee.SignalDiList, error) {
	var signalDis []*littlebee.SignalDi
	s := "SELECT * FROM signal_dis "
	w, v := queryclause.QueryCommandString(q)
	s = s + w
	result := database.DB.Raw(s, v...).Scan(&signalDis)
	return &littlebee.SignalDiList{Items: signalDis}, result.Error
}

func GetAllSignalDis() (*littlebee.SignalDiList, error) {
	var signalDis []*littlebee.SignalDi
	result := database.DB.Find(&signalDis)
	return &littlebee.SignalDiList{Items: signalDis}, result.Error
}
