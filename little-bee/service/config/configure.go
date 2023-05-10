package config

import (
	dao "dao/config"
	"message/littlebee"
	"message/querycommand"

	"google.golang.org/protobuf/types/known/timestamppb"
)

func CreateConfigure(l *littlebee.Configure, userId int32) (*littlebee.Configure, error) {
	l.ConfigureId = 0
	l.CreatedAt = timestamppb.Now()
	l.CreatedBy = userId
	return dao.CreateConfigure(l)
}

func RetrieveConfigure(configureId int32) (*littlebee.Configure, error) {
	return dao.RetrieveConfigure(configureId)
}

func UpdateConfigure(l *littlebee.Configure, userId int32) error {
	l.UpdatedAt = timestamppb.Now()
	l.UpdatedBy = userId
	return dao.UpdateConfigure(l)
}

func DeleteConfigureWithFlag(configureId int, userId int32) error {
	l := &littlebee.Configure{
		ConfigureId: int32(configureId),
		IsDeleted:     true,
		DeletedAt:     timestamppb.Now(),
		DeletedBy:     userId,
	}
	return dao.DeleteConfigureWithFlag(l)
}

func DeleteConfigure(configureId int, userId int32) error {
	l := &littlebee.Configure{
		ConfigureId: int32(configureId),
		IsDeleted:     true,
		DeletedAt:     timestamppb.Now(),
		DeletedBy:     userId,
	}
	return dao.DeleteConfigure(l)
}

func QueryConfigure(q *querycommand.QueryCommand) (*littlebee.ConfigureList, error) {
	return dao.QueryConfigure(q)
}

func GetAllConfigures() (*littlebee.ConfigureList, error) {
	return dao.GetAllConfigures()
}