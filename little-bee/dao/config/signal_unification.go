package config

import (
	"dao/database"
	"message/littlebee"
	"message/querycommand"
	"util/queryclause"

	"github.com/fatih/structs"
)

func CreateSignalUnification(l *littlebee.SignalUnification) (*littlebee.SignalUnification, error) {
	result := database.DB.Omit("SignalUnificationId", "UpdatedAt", "UpdatedBy", "IsDeleted", "DeletedAt", "DeletedBy").Create(l)
	return l, result.Error
}

func RetrieveSignalUnification(signalUnificationId int32) (*littlebee.SignalUnification, error) {
	r := &littlebee.SignalUnification{}
	result := database.DB.First(&r, signalUnificationId)
	return r, result.Error
}

func UpdateSignalUnification(l *littlebee.SignalUnification) error {
	lMap := structs.Map(l)
	lMap["created_at"] = l.CreatedAt.AsTime().UTC()
	lMap["updated_at"] = l.UpdatedAt.AsTime().UTC()
	lMap["deleted_at"] = l.DeletedAt.AsTime().UTC()
	result := database.DB.Model(l).Omit("SignalUnificationId", "CreatedAt", "CreatedBy", "IsDeleted", "DeletedAt", "DeletedBy").Updates(lMap)
	return result.Error
}

func DeleteSignalUnificationWithFlag(l *littlebee.SignalUnification) error {
	result := database.DB.Model(l).Updates(littlebee.SignalUnification{IsDeleted: l.IsDeleted, DeletedAt: l.DeletedAt, DeletedBy: l.DeletedBy})
	return result.Error
}

func DeleteSignalUnification(l *littlebee.SignalUnification) error {
	result := database.DB.Delete(l)
	return result.Error
}

func QuerySignalUnification(q *querycommand.QueryCommand) (*littlebee.SignalUnificationList, error) {
	var signalUnifications []*littlebee.SignalUnification
	s := "SELECT * FROM signal_unifications "
	w, v := queryclause.QueryCommandString(q)
	s = s + w
	result := database.DB.Raw(s, v...).Scan(&signalUnifications)
	return &littlebee.SignalUnificationList{Items: signalUnifications}, result.Error
}

func GetAllSignalUnifications() (*littlebee.SignalUnificationList, error) {
	var signalUnifications []*littlebee.SignalUnification
	result := database.DB.Find(&signalUnifications)
	return &littlebee.SignalUnificationList{Items: signalUnifications}, result.Error
}
