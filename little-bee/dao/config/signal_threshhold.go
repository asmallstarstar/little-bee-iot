package config

import (
	"dao/database"
	"message/littlebee"
	"message/querycommand"
	"util/queryclause"

	"github.com/fatih/structs"
)

func CreateSignalThreshhold(l *littlebee.SignalThreshold) (*littlebee.SignalThreshold, error) {
	result := database.DB.Omit("SignalThreshholdId").Create(l)
	return l, result.Error
}

func RetrieveSignalThreshhold(signalThreshholdId int32) (*littlebee.SignalThreshold, error) {
	r := &littlebee.SignalThreshold{}
	result := database.DB.First(&r, signalThreshholdId)
	return r, result.Error
}

func UpdateSignalThreshhold(l *littlebee.SignalThreshold) error {
	lMap := structs.Map(l)
	result := database.DB.Model(l).Omit("SignalThreshholdId").Updates(lMap)
	return result.Error
}

func DeleteSignalThreshhold(l *littlebee.SignalThreshold) error {
	result := database.DB.Delete(l)
	return result.Error
}

func QuerySignalThreshhold(q *querycommand.QueryCommand) (*littlebee.SignalThresholdList, error) {
	var signalThreshholds []*littlebee.SignalThreshold
	s := "SELECT * FROM signal_thresholds "
	w, v := queryclause.QueryCommandString(q)
	s = s + w
	result := database.DB.Raw(s, v...).Scan(&signalThreshholds)
	return &littlebee.SignalThresholdList{Items: signalThreshholds}, result.Error
}

func GetAllSignalThreshholds() (*littlebee.SignalThresholdList, error) {
	var signalThreshholds []*littlebee.SignalThreshold
	result := database.DB.Find(&signalThreshholds)
	return &littlebee.SignalThresholdList{Items: signalThreshholds}, result.Error
}
