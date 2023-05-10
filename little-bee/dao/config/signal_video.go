package config

import (
	"dao/database"
	"message/littlebee"
	"message/querycommand"
	"util/queryclause"

	"github.com/fatih/structs"
)

func CreateSignalVideo(l *littlebee.SignalVideo) (*littlebee.SignalVideo, error) {
	result := database.DB.Omit("SignalVideoId").Create(l)
	return l, result.Error
}

func RetrieveSignalVideo(signalVideoId int32) (*littlebee.SignalVideo, error) {
	r := &littlebee.SignalVideo{}
	result := database.DB.First(&r, signalVideoId)
	return r, result.Error
}

func UpdateSignalVideo(l *littlebee.SignalVideo) error {
	lMap := structs.Map(l)
	result := database.DB.Model(l).Omit("SignalVideoId").Updates(lMap)
	return result.Error
}

func DeleteSignalVideo(l *littlebee.SignalVideo) error {
	result := database.DB.Delete(l)
	return result.Error
}

func QuerySignalVideo(q *querycommand.QueryCommand) (*littlebee.SignalVideoList, error) {
	var signalVideos []*littlebee.SignalVideo
	s := "SELECT * FROM signal_videos "
	w, v := queryclause.QueryCommandString(q)
	s = s + w
	result := database.DB.Raw(s, v...).Scan(&signalVideos)
	return &littlebee.SignalVideoList{Items: signalVideos}, result.Error
}

func GetAllSignalVideos() (*littlebee.SignalVideoList, error) {
	var signalVideos []*littlebee.SignalVideo
	result := database.DB.Find(&signalVideos)
	return &littlebee.SignalVideoList{Items: signalVideos}, result.Error
}
