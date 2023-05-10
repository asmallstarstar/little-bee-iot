package config

import (
	"dao/database"
	"message/littlebee"
	"message/querycommand"
	"util/queryclause"

	"github.com/fatih/structs"
)

func CreateAlarmLevel(l *littlebee.AlarmLevel) (*littlebee.AlarmLevel, error) {
	result := database.DB.Omit("AlarmLevelId", "UpdatedAt", "UpdatedBy", "IsDeleted", "DeletedAt", "DeletedBy").Create(l)
	return l, result.Error
}

func RetrieveAlarmLevel(alarmLevelId int32) (*littlebee.AlarmLevel, error) {
	r := &littlebee.AlarmLevel{}
	result := database.DB.First(&r, alarmLevelId)
	return r, result.Error
}

func UpdateAlarmLevel(l *littlebee.AlarmLevel) error {
	lMap := structs.Map(l)
	lMap["created_at"] = l.CreatedAt.AsTime().UTC()
	lMap["updated_at"] = l.UpdatedAt.AsTime().UTC()
	lMap["deleted_at"] = l.DeletedAt.AsTime().UTC()
	result := database.DB.Model(l).Omit("AlarmLevelId", "CreatedAt", "CreatedBy", "IsDeleted", "DeletedAt", "DeletedBy").Updates(lMap)
	return result.Error
}

func DeleteAlarmLevelWithFlag(l *littlebee.AlarmLevel) error {
	result := database.DB.Model(l).Updates(littlebee.AlarmLevel{IsDeleted: l.IsDeleted, DeletedAt: l.DeletedAt, DeletedBy: l.DeletedBy})
	return result.Error
}

func DeleteAlarmLevel(l *littlebee.AlarmLevel) error {
	result := database.DB.Delete(l)
	return result.Error
}

func QueryAlarmLevel(q *querycommand.QueryCommand) (*littlebee.AlarmLevelList, error) {
	var alarmLevels []*littlebee.AlarmLevel
	s := "SELECT * FROM alarm_levels "
	w, v := queryclause.QueryCommandString(q)
	s = s + w
	result := database.DB.Raw(s, v...).Scan(&alarmLevels)
	return &littlebee.AlarmLevelList{Items: alarmLevels}, result.Error
}

func GetAllAlarmLevels() (*littlebee.AlarmLevelList, error) {
	var alarmLevels []*littlebee.AlarmLevel
	result := database.DB.Find(&alarmLevels)
	return &littlebee.AlarmLevelList{Items: alarmLevels}, result.Error
}
