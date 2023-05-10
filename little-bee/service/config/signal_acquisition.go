package config

import (
	dao "dao/config"
	"message/littlebee"
	"message/querycommand"
)

func CreateSignalAcquisition(l *littlebee.SignalAcquisition, userId int32) (*littlebee.SignalAcquisition, error) {
	return dao.CreateSignalAcquisition(l)
}

func RetrieveSignalAcquisition(signalAcquisitionId int32) (*littlebee.SignalAcquisition, error) {
	return dao.RetrieveSignalAcquisition(signalAcquisitionId)
}

func UpdateSignalAcquisition(l *littlebee.SignalAcquisition, userId int32) error {
	return dao.UpdateSignalAcquisition(l)
}

func DeleteSignalAcquisition(signalAcquisitionId int, userId int32) error {
	l := &littlebee.SignalAcquisition{
		SignalId:  int32(signalAcquisitionId),
	}
	return dao.DeleteSignalAcquisition(l)
}

func QuerySignalAcquisition(q *querycommand.QueryCommand) (*littlebee.SignalAcquisitionList, error) {
	return dao.QuerySignalAcquisition(q)
}

func GetAllSignalAcquisitions() (*littlebee.SignalAcquisitionList, error) {
	return dao.GetAllSignalAcquisitions()
}
