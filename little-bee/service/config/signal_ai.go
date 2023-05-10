package config

import (
	dao "dao/config"
	"message/littlebee"
	"message/querycommand"
)

func CreateSignalAi(l *littlebee.SignalAi, userId int32) (*littlebee.SignalAi, error) {
	return dao.CreateSignalAi(l)
}

func RetrieveSignalAi(signalAiId int32) (*littlebee.SignalAi, error) {
	return dao.RetrieveSignalAi(signalAiId)
}

func UpdateSignalAi(l *littlebee.SignalAi, userId int32) error {
	return dao.UpdateSignalAi(l)
}

func DeleteSignalAi(signalAiId int, userId int32) error {
	l := &littlebee.SignalAi{
		SignalId:  int32(signalAiId),
	}
	return dao.DeleteSignalAi(l)
}

func QuerySignalAi(q *querycommand.QueryCommand) (*littlebee.SignalAiList, error) {
	return dao.QuerySignalAi(q)
}

func GetAllSignalAis() (*littlebee.SignalAiList, error) {
	return dao.GetAllSignalAis()
}
