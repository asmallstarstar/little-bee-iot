package config

import (
	"dao/database"
	"message/littlebee"
	"message/querycommand"
	"util/queryclause"

	"github.com/fatih/structs"
)

func CreateSignalAcquisition(l *littlebee.SignalAcquisition) (*littlebee.SignalAcquisition, error) {
	result := database.DB.Omit("SignalAcquisitionId", "UpdatedAt", "UpdatedBy", "IsDeleted", "DeletedAt", "DeletedBy").Create(l)
	return l, result.Error
}

func RetrieveSignalAcquisition(signalAcquisitionId int32) (*littlebee.SignalAcquisition, error) {
	r := &littlebee.SignalAcquisition{}
	result := database.DB.First(&r, signalAcquisitionId)
	return r, result.Error
}

func UpdateSignalAcquisition(l *littlebee.SignalAcquisition) error {
	lMap := structs.Map(l)
	result := database.DB.Model(l).Omit("SignalAcquisitionId").Updates(lMap)
	return result.Error
}

func DeleteSignalAcquisition(l *littlebee.SignalAcquisition) error {
	result := database.DB.Delete(l)
	return result.Error
}

func QuerySignalAcquisition(q *querycommand.QueryCommand) (*littlebee.SignalAcquisitionList, error) {
	var signalAcquisitions []*littlebee.SignalAcquisition
	s := "SELECT * FROM signal_acquisitions "
	w, v := queryclause.QueryCommandString(q)
	s = s + w
	result := database.DB.Raw(s, v...).Scan(&signalAcquisitions)
	return &littlebee.SignalAcquisitionList{Items: signalAcquisitions}, result.Error
}

func GetAllSignalAcquisitions() (*littlebee.SignalAcquisitionList, error) {
	var signalAcquisitions []*littlebee.SignalAcquisition
	result := database.DB.Find(&signalAcquisitions)
	return &littlebee.SignalAcquisitionList{Items: signalAcquisitions}, result.Error
}
