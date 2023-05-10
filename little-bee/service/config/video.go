package config

import (
	dao "dao/config"
	"message/littlebee"

	"google.golang.org/protobuf/types/known/timestamppb"
)

func CreateVideo(l *littlebee.Video, userId int32) (*littlebee.Video, error) {
	l.Signal.SignalId = 0
	l.Signal.CreatedAt = timestamppb.Now()
	l.Signal.CreatedBy = userId
	if l.VideoBindConfigure != nil {
		l.VideoBindConfigure.CreatedAt = timestamppb.Now()
		l.VideoBindConfigure.CreatedBy = userId
	}
	return dao.CreateVideo(l)
}

func RetrieveVideo(videoId int32) (*littlebee.Video, error) {
	return dao.RetrieveVideo(videoId)
}

func UpdateVideo(l *littlebee.Video, userId int32) error {
	l.Signal.UpdatedAt = timestamppb.Now()
	l.Signal.UpdatedBy = userId
	if l.VideoBindConfigure != nil {
		if l.VideoBindConfigure.ConfigureId == 0 {
			l.VideoBindConfigure.CreatedAt = timestamppb.Now()
			l.VideoBindConfigure.CreatedBy = userId
		} else {
			l.VideoBindConfigure.UpdatedAt = timestamppb.Now()
			l.VideoBindConfigure.UpdatedBy = userId
		}
	}
	return dao.UpdateVideo(l)
}

func DeleteVideoWithFlag(videoId int, userId int32) error {
	l := &littlebee.Video{
		Signal: &littlebee.Signal{
			SignalId:  int32(videoId),
			IsDeleted: true,
			DeletedAt: timestamppb.Now(),
			DeletedBy: userId,
		},
		SignalVideo: &littlebee.SignalVideo{SignalId: int32(videoId)},
	}
	return dao.DeleteVideoWithFlag(l)
}

func DeleteVideo(videoId int, userId int32) error {
	l := &littlebee.Video{
		Signal: &littlebee.Signal{
			SignalId:  int32(videoId),
			IsDeleted: true,
			DeletedAt: timestamppb.Now(),
			DeletedBy: userId,
		},
		SignalVideo: &littlebee.SignalVideo{SignalId: int32(videoId)},
	}
	return dao.DeleteVideo(l)
}
