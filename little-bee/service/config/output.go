package config

import (
	dao "dao/config"
	"message/littlebee"

	"google.golang.org/protobuf/types/known/timestamppb"
)

func CreateOutput(l *littlebee.Output, userId int32) (*littlebee.Output, error) {
	l.Signal.SignalId = 0
	l.Signal.CreatedAt = timestamppb.Now()
	l.Signal.CreatedBy = userId
	if l.SignalBindConfigure != nil {
		l.SignalBindConfigure.CreatedAt = timestamppb.Now()
		l.SignalBindConfigure.CreatedBy = userId
	}
	return dao.CreateOutput(l)
}

func RetrieveOutput(outputId int32) (*littlebee.Output, error) {
	return dao.RetrieveOutput(outputId)
}

func UpdateOutput(l *littlebee.Output, userId int32) error {
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
	return dao.UpdateOutput(l)
}

func DeleteOutputWithFlag(outputId int, userId int32) error {
	l := &littlebee.Output{
		Signal: &littlebee.Signal{
			SignalId:  int32(outputId),
			IsDeleted: true,
			DeletedAt: timestamppb.Now(),
			DeletedBy: userId,
		},
		SignalAcquisition: &littlebee.SignalAcquisition{SignalId: int32(outputId)},
	}
	return dao.DeleteOutputWithFlag(l)
}

func DeleteOutput(outputId int, userId int32) error {
	l := &littlebee.Output{
		Signal: &littlebee.Signal{
			SignalId:  int32(outputId),
			IsDeleted: true,
			DeletedAt: timestamppb.Now(),
			DeletedBy: userId,
		},
		SignalAcquisition: &littlebee.SignalAcquisition{SignalId: int32(outputId)},
	}
	return dao.DeleteOutput(l)
}
