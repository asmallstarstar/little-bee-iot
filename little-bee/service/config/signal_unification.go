package config

import (
	dao "dao/config"
	"message/littlebee"
	"message/querycommand"

	"google.golang.org/protobuf/types/known/timestamppb"
)

func CreateSignalUnification(l *littlebee.SignalUnification, userId int32) (*littlebee.SignalUnification, error) {
	l.SignalUnificationId = 0
	l.CreatedAt = timestamppb.Now()
	l.CreatedBy = userId
	return dao.CreateSignalUnification(l)
}

func RetrieveSignalUnification(signalUnificationId int32) (*littlebee.SignalUnification, error) {
	return dao.RetrieveSignalUnification(signalUnificationId)
}

func UpdateSignalUnification(l *littlebee.SignalUnification, userId int32) error {
	l.UpdatedAt = timestamppb.Now()
	l.UpdatedBy = userId
	return dao.UpdateSignalUnification(l)
}

func DeleteSignalUnificationWithFlag(signalUnificationId int, userId int32) error {
	l := &littlebee.SignalUnification{
		SignalUnificationId: int32(signalUnificationId),
		IsDeleted:     true,
		DeletedAt:     timestamppb.Now(),
		DeletedBy:     userId,
	}
	return dao.DeleteSignalUnificationWithFlag(l)
}

func DeleteSignalUnification(signalUnificationId int, userId int32) error {
	l := &littlebee.SignalUnification{
		SignalUnificationId: int32(signalUnificationId),
		IsDeleted:     true,
		DeletedAt:     timestamppb.Now(),
		DeletedBy:     userId,
	}
	return dao.DeleteSignalUnification(l)
}

func QuerySignalUnification(q *querycommand.QueryCommand) (*littlebee.SignalUnificationList, error) {
	return dao.QuerySignalUnification(q)
}

func GetAllSignalUnifications() (*littlebee.SignalUnificationList, error) {
	return dao.GetAllSignalUnifications()
}