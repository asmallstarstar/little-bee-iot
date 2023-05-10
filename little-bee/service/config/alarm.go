package config

import (
	dao "dao/config"
	"message/littlebee"
	"message/querycommand"
	"service/consumer"

	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func CreateAlarm(l *littlebee.Alarm, userId int32) (*littlebee.Alarm, error) {
	return dao.CreateAlarm(l)
}

func RetrieveAlarm(alarmId int32) (*littlebee.Alarm, error) {
	return dao.RetrieveAlarm(alarmId)
}

func EndAlarmBySerialNumber(l *littlebee.Alarm) error {
	return dao.EndAlarmBySerialNumber(l)
}

func EndAlarmByAlarmTime(l *littlebee.Alarm) error {
	return dao.EndAlarmByAlarmTime(l)
}

func UpdateAlarm(l *littlebee.Alarm, userId int32) error {
	return dao.UpdateAlarm(l)
}

func AckAlarm(l *littlebee.Alarm, userId int32) error {
	l.AckUserId = userId
	l.AckState = 1
	l.AckTime = timestamppb.Now()
	m, _ := proto.Marshal(l)
	go consumer.PublishAckAlarm(m)
	return dao.AckAlarm(l)
}

func QueryAlarm(q *querycommand.QueryCommand) (*littlebee.AlarmList, error) {
	return dao.QueryAlarm(q)
}

func GetAllAlarms() (*littlebee.AlarmList, error) {
	return dao.GetAllAlarms()
}

func EndAlarm(l *littlebee.Alarm) error {
	return dao.EndAlarm(l)
}
