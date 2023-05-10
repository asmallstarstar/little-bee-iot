package config

import (
	"dao/database"
	"message/littlebee"
	"message/querycommand"
	"util/queryclause"

	"github.com/fatih/structs"
)

func CreateSignalSi(l *littlebee.SignalSi) (*littlebee.SignalSi, error) {
	result := database.DB.Omit("SignalSiId", "UpdatedAt", "UpdatedBy", "IsDeleted", "DeletedAt", "DeletedBy").Create(l)
	return l, result.Error
}

func RetrieveSignalSi(signalSiId int32) (*littlebee.SignalSi, error) {
	r := &littlebee.SignalSi{}
	result := database.DB.First(&r, signalSiId)
	return r, result.Error
}

func UpdateSignalSi(l *littlebee.SignalSi) error {
	lMap := structs.Map(l)
	result := database.DB.Model(l).Omit("SignalAiId").Updates(lMap)
	return result.Error
}

func DeleteSignalSi(l *littlebee.SignalSi) error {
	result := database.DB.Delete(l)
	return result.Error
}

func QuerySignalSi(q *querycommand.QueryCommand) (*littlebee.SignalSiList, error) {
	var signalSis []*littlebee.SignalSi
	s := "SELECT * FROM signal_sis "
	w, v := queryclause.QueryCommandString(q)
	s = s + w
	result := database.DB.Raw(s, v...).Scan(&signalSis)
	return &littlebee.SignalSiList{Items: signalSis}, result.Error
}

func GetAllSignalSis() (*littlebee.SignalSiList, error) {
	var signalSis []*littlebee.SignalSi
	result := database.DB.Find(&signalSis)
	return &littlebee.SignalSiList{Items: signalSis}, result.Error
}
