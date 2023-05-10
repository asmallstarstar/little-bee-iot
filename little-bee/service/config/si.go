package config

import (
	dao "dao/config"
	"message/littlebee"

	"google.golang.org/protobuf/types/known/timestamppb"
)

func CreateSi(l *littlebee.Si, userId int32) (*littlebee.Si, error) {
	l.Signal.SignalId = 0
	l.Signal.CreatedAt = timestamppb.Now()
	l.Signal.CreatedBy = userId
	if l.SignalBindConfigure != nil {
		l.SignalBindConfigure.CreatedAt = timestamppb.Now()
		l.SignalBindConfigure.CreatedBy = userId
	}
	return dao.CreateSi(l)
}

func RetrieveSi(siId int32) (*littlebee.Si, error) {
	return dao.RetrieveSi(siId)
}

func UpdateSi(l *littlebee.Si, userId int32) error {
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
	return dao.UpdateSi(l)
}

func DeleteSiWithFlag(siId int, userId int32) error {
	l := &littlebee.Si{
		Signal: &littlebee.Signal{
			SignalId:  int32(siId),
			IsDeleted: true,
			DeletedAt: timestamppb.Now(),
			DeletedBy: userId,
		},
		SignalAcquisition: &littlebee.SignalAcquisition{SignalId: int32(siId)},
		SignalSi:          &littlebee.SignalSi{SignalId: int32(siId)},
	}
	return dao.DeleteSiWithFlag(l)
}

func DeleteSi(siId int, userId int32) error {
	l := &littlebee.Si{
		Signal: &littlebee.Signal{
			SignalId:  int32(siId),
			IsDeleted: true,
			DeletedAt: timestamppb.Now(),
			DeletedBy: userId,
		},
		SignalAcquisition: &littlebee.SignalAcquisition{SignalId: int32(siId)},
		SignalSi:          &littlebee.SignalSi{SignalId: int32(siId)},
	}
	return dao.DeleteSi(l)
}
