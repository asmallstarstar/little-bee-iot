package config

import (
	dao "dao/config"
	"message/littlebee"
	"message/querycommand"
)

func CreateSignalSi(l *littlebee.SignalSi, userId int32) (*littlebee.SignalSi, error) {
	return dao.CreateSignalSi(l)
}

func RetrieveSignalSi(signalSiId int32) (*littlebee.SignalSi, error) {
	return dao.RetrieveSignalSi(signalSiId)
}

func UpdateSignalSi(l *littlebee.SignalSi, userId int32) error {
	return dao.UpdateSignalSi(l)
}

func DeleteSignalSi(signalSiId int, userId int32) error {
	l := &littlebee.SignalSi{
		SignalId:  int32(signalSiId),
	}
	return dao.DeleteSignalSi(l)
}

func QuerySignalSi(q *querycommand.QueryCommand) (*littlebee.SignalSiList, error) {
	return dao.QuerySignalSi(q)
}

func GetAllSignalSis() (*littlebee.SignalSiList, error) {
	return dao.GetAllSignalSis()
}
