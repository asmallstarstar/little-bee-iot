package config

import (
	dao "dao/config"
	"message/littlebee"
	"message/querycommand"

	"google.golang.org/protobuf/types/known/timestamppb"
)

func CreateFsuType(l *littlebee.FsuType, userId int32) (*littlebee.FsuType, error) {
	l.FsuTypeId = 0
	l.CreatedAt = timestamppb.Now()
	l.CreatedBy = userId
	return dao.CreateFsuType(l)
}

func RetrieveFsuType(fsuTypeId int32) (*littlebee.FsuType, error) {
	return dao.RetrieveFsuType(fsuTypeId)
}

func UpdateFsuType(l *littlebee.FsuType, userId int32) error {
	l.UpdatedAt = timestamppb.Now()
	l.UpdatedBy = userId
	return dao.UpdateFsuType(l)
}

func DeleteFsuTypeWithFlag(fsuTypeId int, userId int32) error {
	l := &littlebee.FsuType{
		FsuTypeId: int32(fsuTypeId),
		IsDeleted:     true,
		DeletedAt:     timestamppb.Now(),
		DeletedBy:     userId,
	}
	return dao.DeleteFsuTypeWithFlag(l)
}

func DeleteFsuType(fsuTypeId int, userId int32) error {
	l := &littlebee.FsuType{
		FsuTypeId: int32(fsuTypeId),
		IsDeleted:     true,
		DeletedAt:     timestamppb.Now(),
		DeletedBy:     userId,
	}
	return dao.DeleteFsuType(l)
}

func QueryFsuType(q *querycommand.QueryCommand) (*littlebee.FsuTypeList, error) {
	return dao.QueryFsuType(q)
}

func GetAllFsuTypes() (*littlebee.FsuTypeList, error) {
	return dao.GetAllFsuTypes()
}