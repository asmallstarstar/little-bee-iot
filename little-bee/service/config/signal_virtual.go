package config

import (
	dao "dao/config"
	"message/littlebee"
	"message/querycommand"
)

func CreateSignalVirtual(l *littlebee.SignalVirtual, userId int32) (*littlebee.SignalVirtual, error) {
	return dao.CreateSignalVirtual(l)
}

func RetrieveSignalVirtual(signalVirtualId int32) (*littlebee.SignalVirtual, error) {
	return dao.RetrieveSignalVirtual(signalVirtualId)
}

func UpdateSignalVirtual(l *littlebee.SignalVirtual, userId int32) error {
	return dao.UpdateSignalVirtual(l)
}

func DeleteSignalVirtual(signalVirtualId int, userId int32) error {
	l := &littlebee.SignalVirtual{
		SignalId: int32(signalVirtualId),
	}
	return dao.DeleteSignalVirtual(l)
}

func QuerySignalVirtual(q *querycommand.QueryCommand) (*littlebee.SignalVirtualList, error) {
	return dao.QuerySignalVirtual(q)
}

func GetAllSignalVirtuals() (*littlebee.SignalVirtualList, error) {
	return dao.GetAllSignalVirtuals()
}
