package config

import (
	dao "dao/config"
	"message/littlebee"

	"google.golang.org/protobuf/types/known/timestamppb"
)

func CreateVai(l *littlebee.VAi, userId int32) (*littlebee.VAi, error) {
	l.Signal.SignalId = 0
	l.Signal.CreatedAt = timestamppb.Now()
	l.Signal.CreatedBy = userId
	return dao.CreateVai(l)
}

func RetrieveVai(vaiId int32) (*littlebee.VAi, error) {
	return dao.RetrieveVai(vaiId)
}

func UpdateVai(l *littlebee.VAi, userId int32) error {
	l.Signal.UpdatedAt = timestamppb.Now()
	l.Signal.UpdatedBy = userId
	return dao.UpdateVai(l)
}

func DeleteVaiWithFlag(vaiId int, userId int32) error {
	l := &littlebee.VAi{
		Signal: &littlebee.Signal{
			SignalId:  int32(vaiId),
			IsDeleted: true,
			DeletedAt: timestamppb.Now(),
			DeletedBy: userId,
		},
		SignalVirtual:   &littlebee.SignalVirtual{SignalId: int32(vaiId)},
		SignalAi:        &littlebee.SignalAi{SignalId: int32(vaiId)},
		SignalThreshold: []*littlebee.SignalThreshold{},
	}
	return dao.DeleteVaiWithFlag(l)
}

func DeleteVai(vaiId int, userId int32) error {
	l := &littlebee.VAi{
		Signal: &littlebee.Signal{
			SignalId:  int32(vaiId),
			IsDeleted: true,
			DeletedAt: timestamppb.Now(),
			DeletedBy: userId,
		},
		SignalVirtual:   &littlebee.SignalVirtual{SignalId: int32(vaiId)},
		SignalAi:        &littlebee.SignalAi{SignalId: int32(vaiId)},
		SignalThreshold: []*littlebee.SignalThreshold{},
	}
	return dao.DeleteVai(l)
}
