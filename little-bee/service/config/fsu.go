package config

import (
	dao "dao/config"
	"message/littlebee"
	"message/querycommand"

	"google.golang.org/protobuf/types/known/timestamppb"
)

func CreateFsu(l *littlebee.Fsu, userId int32) (*littlebee.Fsu, error) {
	l.FsuId = 0
	l.CreatedAt = timestamppb.Now()
	l.CreatedBy = userId
	return dao.CreateFsu(l)
}

func RetrieveFsu(fsuId int32) (*littlebee.Fsu, error) {
	return dao.RetrieveFsu(fsuId)
}

func UpdateFsu(l *littlebee.Fsu, userId int32) error {
	l.UpdatedAt = timestamppb.Now()
	l.UpdatedBy = userId
	return dao.UpdateFsu(l)
}

func DeleteFsuWithFlag(fsuId int, userId int32) error {
	l := &littlebee.Fsu{
		FsuId: int32(fsuId),
		IsDeleted:     true,
		DeletedAt:     timestamppb.Now(),
		DeletedBy:     userId,
	}
	return dao.DeleteFsuWithFlag(l)
}

func DeleteFsu(fsuId int, userId int32) error {
	l := &littlebee.Fsu{
		FsuId: int32(fsuId),
		IsDeleted:     true,
		DeletedAt:     timestamppb.Now(),
		DeletedBy:     userId,
	}
	return dao.DeleteFsu(l)
}

func QueryFsu(q *querycommand.QueryCommand) (*littlebee.FsuList, error) {
	return dao.QueryFsu(q)
}

func GetAllFsus() (*littlebee.FsuList, error) {
	return dao.GetAllFsus()
}