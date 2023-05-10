package config

import (
	dao "dao/config"
	"message/littlebee"
	"message/querycommand"
)

func CreateSignalThreshhold(l *littlebee.SignalThreshold, userId int32) (*littlebee.SignalThreshold, error) {
	return dao.CreateSignalThreshhold(l)
}

func RetrieveSignalThreshhold(signalThreshholdId int32) (*littlebee.SignalThreshold, error) {
	return dao.RetrieveSignalThreshhold(signalThreshholdId)
}

func UpdateSignalThreshhold(l *littlebee.SignalThreshold, userId int32) error {
	return dao.UpdateSignalThreshhold(l)
}

func DeleteSignalThreshhold(signalThresholdId int, userId int32) error {
	l := &littlebee.SignalThreshold{
		ThresholdId: int32(signalThresholdId),
	}
	return dao.DeleteSignalThreshhold(l)
}

func QuerySignalThreshhold(q *querycommand.QueryCommand) (*littlebee.SignalThresholdList, error) {
	return dao.QuerySignalThreshhold(q)
}

func GetAllSignalThreshholds() (*littlebee.SignalThresholdList, error) {
	return dao.GetAllSignalThreshholds()
}
