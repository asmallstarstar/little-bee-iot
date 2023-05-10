package config

import (
	"dao/database"
	"message/littlebee"
	"message/querycommand"
	"util/queryclause"

	"github.com/fatih/structs"
)

func CreateAlarm(l *littlebee.Alarm) (*littlebee.Alarm, error) {
	result := database.DB.Create(l)
	return l, result.Error
}

func RetrieveAlarm(alarmId int32) (*littlebee.Alarm, error) {
	r := &littlebee.Alarm{}
	result := database.DB.First(&r, alarmId)
	return r, result.Error
}

func EndAlarmBySerialNumber(l *littlebee.Alarm) error {
	result := database.DB.Model(l).Where("alarm_serial_number = ?", l.AlarmSerialNumber).Select("AlarmEndTime", "AlarmDesc").Updates(l)
	return result.Error
}

func EndAlarmByAlarmTime(l *littlebee.Alarm) error {
	result := database.DB.Model(l).Where("signal_id = ? AND alarm_begin_time = ?", l.AlarmBeginTime, l.AlarmSerialNumber).Select("AlarmEndTime", "AlarmDesc").Updates(l)
	return result.Error
}

func UpdateAlarm(l *littlebee.Alarm) error {
	lMap := structs.Map(l)
	result := database.DB.Model(l).Omit("AlarmId", "CreatedAt", "CreatedBy", "IsDeleted", "DeletedAt", "DeletedBy").Updates(lMap)
	return result.Error
}

func AckAlarm(l *littlebee.Alarm) error {
	lMap := structs.Map(l)
	lMap["alarm_begin_time"] = l.AlarmBeginTime.AsTime().UTC()
	lMap["ack_time"] = l.AckTime.AsTime().UTC()
	result := database.DB.Model(l).Where("signal_id = ? AND alarm_begin_time = ? AND ack_state = 0", l.SignalId, lMap["alarm_begin_time"]).Select("AckState", "AckTime", "AckUserId", "AckUserName").Updates(lMap)
	return result.Error
}

func DeleteAlarmWithFlag(l *littlebee.Alarm) error {
	result := database.DB.Model(l).Updates(littlebee.Alarm{})
	return result.Error
}

func DeleteAlarm(l *littlebee.Alarm) error {
	result := database.DB.Delete(l)
	return result.Error
}

func QueryAlarm(q *querycommand.QueryCommand) (*littlebee.AlarmList, error) {
	var alarms []*littlebee.Alarm
	s := "SELECT * FROM alarms "
	w, v := queryclause.QueryCommandString(q)
	s = s + w
	result := database.DB.Raw(s, v...).Scan(&alarms)
	return &littlebee.AlarmList{Items: alarms}, result.Error
}

func GetAllAlarms() (*littlebee.AlarmList, error) {
	var alarms []*littlebee.Alarm
	result := database.DB.Find(&alarms)
	return &littlebee.AlarmList{Items: alarms}, result.Error
}

func EndAlarm(l *littlebee.Alarm) error {
	lMap := structs.Map(l)
	result := database.DB.Model(l).Omit("AlarmId", "CreatedAt", "CreatedBy", "IsDeleted", "DeletedAt", "DeletedBy").Updates(lMap)
	return result.Error
}
