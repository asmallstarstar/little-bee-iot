package config

import (
	dao "dao/config"
	"message/littlebee"
	"message/querycommand"

	"google.golang.org/protobuf/types/known/timestamppb"
)

func CreateSignal(l *littlebee.Signal, userId int32) (*littlebee.Signal, error) {
	l.SignalId = 0
	l.CreatedAt = timestamppb.Now()
	l.CreatedBy = userId
	return dao.CreateSignal(l)
}

func RetrieveSignal(signalId int32) (*littlebee.Signal, error) {
	return dao.RetrieveSignal(signalId)
}

func UpdateSignal(l *littlebee.Signal, userId int32) error {
	l.UpdatedAt = timestamppb.Now()
	l.UpdatedBy = userId
	return dao.UpdateSignal(l)
}

func DeleteSignalWithFlag(signalId int, userId int32) error {
	l := &littlebee.Signal{
		SignalId: int32(signalId),
		IsDeleted:     true,
		DeletedAt:     timestamppb.Now(),
		DeletedBy:     userId,
	}
	return dao.DeleteSignalWithFlag(l)
}

func DeleteSignal(signalId int, userId int32) error {
	l := &littlebee.Signal{
		SignalId: int32(signalId),
		IsDeleted:     true,
		DeletedAt:     timestamppb.Now(),
		DeletedBy:     userId,
	}
	return dao.DeleteSignal(l)
}

func QuerySignal(q *querycommand.QueryCommand) (*littlebee.SignalList, error) {
	return dao.QuerySignal(q)
}

func GetAllSignals() (*littlebee.SignalList, error) {
	return dao.GetAllSignals()
}