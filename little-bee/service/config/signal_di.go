package config

import (
	dao "dao/config"
	"message/littlebee"
	"message/querycommand"
)

func CreateSignalDi(l *littlebee.SignalDi, userId int32) (*littlebee.SignalDi, error) {
	return dao.CreateSignalDi(l)
}

func RetrieveSignalDi(signalDiId int32) (*littlebee.SignalDi, error) {
	return dao.RetrieveSignalDi(signalDiId)
}

func UpdateSignalDi(l *littlebee.SignalDi, userId int32) error {
	return dao.UpdateSignalDi(l)
}

func DeleteSignalDi(signalDiId int, userId int32) error {
	l := &littlebee.SignalDi{
		SignalId:  int32(signalDiId),
	}
	return dao.DeleteSignalDi(l)
}

func QuerySignalDi(q *querycommand.QueryCommand) (*littlebee.SignalDiList, error) {
	return dao.QuerySignalDi(q)
}

func GetAllSignalDis() (*littlebee.SignalDiList, error) {
	return dao.GetAllSignalDis()
}
