package config

import (
	dao "dao/config"
	"message/littlebee"

	"google.golang.org/protobuf/types/known/timestamppb"
)

func CreateDi(l *littlebee.Di, userId int32) (*littlebee.Di, error) {
	l.Signal.SignalId = 0
	l.Signal.CreatedAt = timestamppb.Now()
	l.Signal.CreatedBy = userId
	if l.SignalBindConfigure != nil {
		l.SignalBindConfigure.CreatedAt = timestamppb.Now()
		l.SignalBindConfigure.CreatedBy = userId
	}
	return dao.CreateDi(l)
}

func RetrieveDi(diId int32) (*littlebee.Di, error) {
	return dao.RetrieveDi(diId)
}

func UpdateDi(l *littlebee.Di, userId int32) error {
	l.Signal.UpdatedAt = timestamppb.Now()
	l.Signal.UpdatedBy = userId
	if l.SignalBindConfigure != nil {
		if l.SignalBindConfigure.ConfigureId == 0 {
			l.SignalBindConfigure.CreatedAt = timestamppb.Now()
			l.SignalBindConfigure.CreatedBy = userId
		} else {
			l.SignalBindConfigure.UpdatedAt = timestamppb.Now()
			l.SignalBindConfigure.UpdatedBy = userId
		}
	}
	return dao.UpdateDi(l)
}

func DeleteDiWithFlag(diId int, userId int32) error {
	l := &littlebee.Di{
		Signal: &littlebee.Signal{
			SignalId:  int32(diId),
			IsDeleted: true,
			DeletedAt: timestamppb.Now(),
			DeletedBy: userId,
		},
		SignalAcquisition: &littlebee.SignalAcquisition{SignalId: int32(diId)},
		SignalDi:          &littlebee.SignalDi{SignalId: int32(diId)},
		SignalThreshold:   []*littlebee.SignalThreshold{},
	}
	return dao.DeleteDiWithFlag(l)
}

func DeleteDi(diId int, userId int32) error {
	l := &littlebee.Di{
		Signal: &littlebee.Signal{
			SignalId:  int32(diId),
			IsDeleted: true,
			DeletedAt: timestamppb.Now(),
			DeletedBy: userId,
		},
		SignalAcquisition: &littlebee.SignalAcquisition{SignalId: int32(diId)},
		SignalDi:          &littlebee.SignalDi{SignalId: int32(diId)},
		SignalThreshold:   []*littlebee.SignalThreshold{},
	}
	return dao.DeleteDi(l)
}
