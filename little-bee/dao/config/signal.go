package config

import (
	"dao/database"
	"message/littlebee"
	"message/querycommand"
	"util/queryclause"

	"github.com/fatih/structs"
)

func CreateSignal(l *littlebee.Signal) (*littlebee.Signal, error) {
	result := database.DB.Omit("SignalId", "UpdatedAt", "UpdatedBy", "IsDeleted", "DeletedAt", "DeletedBy").Create(l)
	return l, result.Error
}

func RetrieveSignal(signalId int32) (*littlebee.Signal, error) {
	r := &littlebee.Signal{}
	result := database.DB.First(&r, signalId)
	return r, result.Error
}

func UpdateSignal(l *littlebee.Signal) error {
	lMap := structs.Map(l)
	lMap["created_at"] = l.CreatedAt.AsTime().UTC()
	lMap["updated_at"] = l.UpdatedAt.AsTime().UTC()
	lMap["deleted_at"] = l.DeletedAt.AsTime().UTC()
	result := database.DB.Model(l).Omit("SignalId", "CreatedAt", "CreatedBy", "IsDeleted", "DeletedAt", "DeletedBy").Updates(lMap)
	return result.Error
}

func DeleteSignalWithFlag(l *littlebee.Signal) error {
	result := database.DB.Model(l).Updates(littlebee.Signal{IsDeleted: l.IsDeleted, DeletedAt: l.DeletedAt, DeletedBy: l.DeletedBy})
	return result.Error
}

func DeleteSignal(l *littlebee.Signal) error {
	result := database.DB.Delete(l)
	return result.Error
}

func QuerySignal(q *querycommand.QueryCommand) (*littlebee.SignalList, error) {
	var signals []*littlebee.Signal
	s := "SELECT * FROM signals "
	w, v := queryclause.QueryCommandString(q)
	s = s + w
	result := database.DB.Raw(s, v...).Scan(&signals)
	return &littlebee.SignalList{Items: signals}, result.Error
}

func GetAllSignals() (*littlebee.SignalList, error) {
	var signals []*littlebee.Signal
	result := database.DB.Find(&signals)
	return &littlebee.SignalList{Items: signals}, result.Error
}
