package config

import (
	dao "dao/config"
	"message/littlebee"

	"google.golang.org/protobuf/types/known/timestamppb"
)

func CreateVsi(l *littlebee.VSi, userId int32) (*littlebee.VSi, error) {
	l.Signal.SignalId = 0
	l.Signal.CreatedAt = timestamppb.Now()
	l.Signal.CreatedBy = userId
	return dao.CreateVsi(l)
}

func RetrieveVsi(vsiId int32) (*littlebee.VSi, error) {
	return dao.RetrieveVsi(vsiId)
}

func UpdateVsi(l *littlebee.VSi, userId int32) error {
	l.Signal.UpdatedAt = timestamppb.Now()
	l.Signal.UpdatedBy = userId
	return dao.UpdateVsi(l)
}

func DeleteVsiWithFlag(vsiId int, userId int32) error {
	l := &littlebee.VSi{
		Signal: &littlebee.Signal{
			SignalId:  int32(vsiId),
			IsDeleted: true,
			DeletedAt: timestamppb.Now(),
			DeletedBy: userId,
		},
		SignalVirtual: &littlebee.SignalVirtual{SignalId: int32(vsiId)},
		SignalSi:      &littlebee.SignalSi{SignalId: int32(vsiId)},
	}
	return dao.DeleteVsiWithFlag(l)
}

func DeleteVsi(vsiId int, userId int32) error {
	l := &littlebee.VSi{
		Signal: &littlebee.Signal{
			SignalId:  int32(vsiId),
			IsDeleted: true,
			DeletedAt: timestamppb.Now(),
			DeletedBy: userId,
		},
		SignalVirtual: &littlebee.SignalVirtual{SignalId: int32(vsiId)},
		SignalSi:      &littlebee.SignalSi{SignalId: int32(vsiId)},
	}
	return dao.DeleteVsi(l)
}
