package config

import (
	dao "dao/config"
	"message/littlebee"
	"message/querycommand"
)

func CreateControlCommandRecord(l *littlebee.ControlCommandRecord, userId int32) (*littlebee.ControlCommandRecord, error) {
	return dao.CreateControlCommandRecord(l)
}

func RetrieveControlCommandRecord(controlCommandRecordId int32) (*littlebee.ControlCommandRecord, error) {
	return dao.RetrieveControlCommandRecord(controlCommandRecordId)
}

func UpdateControlCommandRecord(l *littlebee.ControlCommandRecord, userId int32) error {
	return dao.UpdateControlCommandRecord(l)
}

func UpdateControlCommandResult(l *littlebee.ControlCommandRecord) error {
	return dao.UpdateControlCommandResult(l)
}

func QueryControlCommandRecord(q *querycommand.QueryCommand) (*littlebee.ControlCommandRecordList, error) {
	return dao.QueryControlCommandRecord(q)
}

func GetAllControlCommandRecords() (*littlebee.ControlCommandRecordList, error) {
	return dao.GetAllControlCommandRecords()
}
