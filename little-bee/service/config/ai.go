package config

import (
	dao "dao/config"
	"message/littlebee"

	"google.golang.org/protobuf/types/known/timestamppb"
)

func CreateAi(l *littlebee.Ai, userId int32) (*littlebee.Ai, error) {
	l.Signal.SignalId = 0
	l.Signal.CreatedAt = timestamppb.Now()
	l.Signal.CreatedBy = userId
	if l.SignalBindConfigure != nil {
		l.SignalBindConfigure.CreatedAt = timestamppb.Now()
		l.SignalBindConfigure.CreatedBy = userId
	}
	return dao.CreateAi(l)
}

func RetrieveAi(aiId int32) (*littlebee.Ai, error) {
	return dao.RetrieveAi(aiId)
}

func UpdateAi(l *littlebee.Ai, userId int32) error {
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
	return dao.UpdateAi(l)
}

func DeleteAiWithFlag(aiId int, userId int32) error {
	l := &littlebee.Ai{
		Signal: &littlebee.Signal{
			SignalId:  int32(aiId),
			IsDeleted: true,
			DeletedAt: timestamppb.Now(),
			DeletedBy: userId,
		},
		SignalAcquisition: &littlebee.SignalAcquisition{SignalId: int32(aiId)},
		SignalAi:          &littlebee.SignalAi{SignalId: int32(aiId)},
		SignalThreshold:   []*littlebee.SignalThreshold{},
	}
	return dao.DeleteAiWithFlag(l)
}

func DeleteAi(aiId int, userId int32) error {
	l := &littlebee.Ai{
		Signal: &littlebee.Signal{
			SignalId:  int32(aiId),
			IsDeleted: true,
			DeletedAt: timestamppb.Now(),
			DeletedBy: userId,
		},
		SignalAcquisition: &littlebee.SignalAcquisition{SignalId: int32(aiId)},
		SignalAi:          &littlebee.SignalAi{SignalId: int32(aiId)},
		SignalThreshold:   []*littlebee.SignalThreshold{},
	}
	return dao.DeleteAi(l)
}
