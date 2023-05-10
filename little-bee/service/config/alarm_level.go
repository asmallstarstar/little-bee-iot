package config

import (
	dao "dao/config"
	"message/littlebee"
	"message/querycommand"

	"google.golang.org/protobuf/types/known/timestamppb"
)

func CreateAlarmLevel(l *littlebee.AlarmLevel, userId int32) (*littlebee.AlarmLevel, error) {
	l.AlarmLevelId = 0
	l.CreatedAt = timestamppb.Now()
	l.CreatedBy = userId
	return dao.CreateAlarmLevel(l)
}

func RetrieveAlarmLevel(alarmLevelId int32) (*littlebee.AlarmLevel, error) {
	return dao.RetrieveAlarmLevel(alarmLevelId)
}

func UpdateAlarmLevel(l *littlebee.AlarmLevel, userId int32) error {
	l.UpdatedAt = timestamppb.Now()
	l.UpdatedBy = userId
	return dao.UpdateAlarmLevel(l)
}

func DeleteAlarmLevelWithFlag(alarmLevelId int, userId int32) error {
	l := &littlebee.AlarmLevel{
		AlarmLevelId: int32(alarmLevelId),
		IsDeleted:     true,
		DeletedAt:     timestamppb.Now(),
		DeletedBy:     userId,
	}
	return dao.DeleteAlarmLevelWithFlag(l)
}

func DeleteAlarmLevel(alarmLevelId int, userId int32) error {
	l := &littlebee.AlarmLevel{
		AlarmLevelId: int32(alarmLevelId),
		IsDeleted:     true,
		DeletedAt:     timestamppb.Now(),
		DeletedBy:     userId,
	}
	return dao.DeleteAlarmLevel(l)
}

func QueryAlarmLevel(q *querycommand.QueryCommand) (*littlebee.AlarmLevelList, error) {
	return dao.QueryAlarmLevel(q)
}

func GetAllAlarmLevels() (*littlebee.AlarmLevelList, error) {
	return dao.GetAllAlarmLevels()
}