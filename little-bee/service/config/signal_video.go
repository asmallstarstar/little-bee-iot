package config

import (
	dao "dao/config"
	"message/littlebee"
	"message/querycommand"
)

func CreateSignalVideo(l *littlebee.SignalVideo, userId int32) (*littlebee.SignalVideo, error) {
	l.SignalId = 0
	return dao.CreateSignalVideo(l)
}

func RetrieveSignalVideo(signalVideoId int32) (*littlebee.SignalVideo, error) {
	return dao.RetrieveSignalVideo(signalVideoId)
}

func UpdateSignalVideo(l *littlebee.SignalVideo, userId int32) error {
	return dao.UpdateSignalVideo(l)
}

func DeleteSignalVideo(signalVideoId int, userId int32) error {
	l := &littlebee.SignalVideo{
		SignalId: int32(signalVideoId),
	}
	return dao.DeleteSignalVideo(l)
}

func QuerySignalVideo(q *querycommand.QueryCommand) (*littlebee.SignalVideoList, error) {
	return dao.QuerySignalVideo(q)
}

func GetAllSignalVideos() (*littlebee.SignalVideoList, error) {
	return dao.GetAllSignalVideos()
}
