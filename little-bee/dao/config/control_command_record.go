package config

import (
	"dao/database"
	"message/littlebee"
	"message/querycommand"
	"util/queryclause"

	"github.com/fatih/structs"
)

func CreateControlCommandRecord(l *littlebee.ControlCommandRecord) (*littlebee.ControlCommandRecord, error) {
	result := database.DB.Omit("ControlCommandRecordId").Create(l)
	return l, result.Error
}

func RetrieveControlCommandRecord(controlCommandRecordId int32) (*littlebee.ControlCommandRecord, error) {
	r := &littlebee.ControlCommandRecord{}
	result := database.DB.First(&r, controlCommandRecordId)
	return r, result.Error
}

func UpdateControlCommandRecord(l *littlebee.ControlCommandRecord) error {
	lMap := structs.Map(l)
	result := database.DB.Model(l).Omit("ControlCommandRecordId", "CreatedAt", "CreatedBy", "IsDeleted", "DeletedAt", "DeletedBy").Updates(lMap)
	return result.Error
}

func UpdateControlCommandResult(l *littlebee.ControlCommandRecord) error {
	result := database.DB.Model(l).Where("control_command_serial_number = ?", l.ControlCommandSerialNumber).Update("control_command_result", l.ControlCommandResult)
	return result.Error
}

func DeleteControlCommandRecordWithFlag(l *littlebee.ControlCommandRecord) error {
	result := database.DB.Model(l).Updates(littlebee.ControlCommandRecord{})
	return result.Error
}

func DeleteControlCommandRecord(l *littlebee.ControlCommandRecord) error {
	result := database.DB.Delete(l)
	return result.Error
}

func QueryControlCommandRecord(q *querycommand.QueryCommand) (*littlebee.ControlCommandRecordList, error) {
	var controlCommandRecords []*littlebee.ControlCommandRecord
	s := "SELECT * FROM control_command_records "
	w, v := queryclause.QueryCommandString(q)
	s = s + w
	result := database.DB.Raw(s, v...).Scan(&controlCommandRecords)
	return &littlebee.ControlCommandRecordList{Items: controlCommandRecords}, result.Error
}

func GetAllControlCommandRecords() (*littlebee.ControlCommandRecordList, error) {
	var controlCommandRecords []*littlebee.ControlCommandRecord
	result := database.DB.Find(&controlCommandRecords)
	return &littlebee.ControlCommandRecordList{Items: controlCommandRecords}, result.Error
}
