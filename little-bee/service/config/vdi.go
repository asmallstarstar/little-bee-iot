package config

import (
	dao "dao/config"
	"message/littlebee"

	"google.golang.org/protobuf/types/known/timestamppb"
)

func CreateVdi(l *littlebee.VDi, userId int32) (*littlebee.VDi, error) {
	l.Signal.SignalId = 0
	l.Signal.CreatedAt = timestamppb.Now()
	l.Signal.CreatedBy = userId
	return dao.CreateVdi(l)
}

func RetrieveVdi(vdiId int32) (*littlebee.VDi, error) {
	return dao.RetrieveVdi(vdiId)
}

func UpdateVdi(l *littlebee.VDi, userId int32) error {
	l.Signal.UpdatedAt = timestamppb.Now()
	l.Signal.UpdatedBy = userId
	return dao.UpdateVdi(l)
}

func DeleteVdiWithFlag(vdiId int, userId int32) error {
	l := &littlebee.VDi{
		Signal: &littlebee.Signal{
			SignalId:  int32(vdiId),
			IsDeleted: true,
			DeletedAt: timestamppb.Now(),
			DeletedBy: userId,
		},
		SignalVirtual:   &littlebee.SignalVirtual{SignalId: int32(vdiId)},
		SignalDi:        &littlebee.SignalDi{SignalId: int32(vdiId)},
		SignalThreshold: []*littlebee.SignalThreshold{},
	}
	return dao.DeleteVdiWithFlag(l)
}

func DeleteVdi(vdiId int, userId int32) error {
	l := &littlebee.VDi{
		Signal: &littlebee.Signal{
			SignalId:  int32(vdiId),
			IsDeleted: true,
			DeletedAt: timestamppb.Now(),
			DeletedBy: userId,
		},
		SignalVirtual:   &littlebee.SignalVirtual{SignalId: int32(vdiId)},
		SignalDi:        &littlebee.SignalDi{SignalId: int32(vdiId)},
		SignalThreshold: []*littlebee.SignalThreshold{},
	}
	return dao.DeleteVdi(l)
}
