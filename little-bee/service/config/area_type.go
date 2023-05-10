package config

import (
	dao "dao/config"
	"message/littlebee"
	"message/querycommand"

	"google.golang.org/protobuf/types/known/timestamppb"
)

func CreateAreaType(l *littlebee.AreaType, userId int32) (*littlebee.AreaType, error) {
	l.AreaTypeId = 0
	l.CreatedAt = timestamppb.Now()
	l.CreatedBy = userId
	return dao.CreateAreaType(l)
}

func RetrieveAreaType(areaTypeId int32) (*littlebee.AreaType, error) {
	return dao.RetrieveAreaType(areaTypeId)
}

func UpdateAreaType(l *littlebee.AreaType, userId int32) error {
	l.UpdatedAt = timestamppb.Now()
	l.UpdatedBy = userId
	return dao.UpdateAreaType(l)
}

func DeleteAreaTypeWithFlag(areaTypeId int, userId int32) error {
	l := &littlebee.AreaType{
		AreaTypeId: int32(areaTypeId),
		IsDeleted:     true,
		DeletedAt:     timestamppb.Now(),
		DeletedBy:     userId,
	}
	return dao.DeleteAreaTypeWithFlag(l)
}

func DeleteAreaType(areaTypeId int, userId int32) error {
	l := &littlebee.AreaType{
		AreaTypeId: int32(areaTypeId),
		IsDeleted:     true,
		DeletedAt:     timestamppb.Now(),
		DeletedBy:     userId,
	}
	return dao.DeleteAreaType(l)
}

func QueryAreaType(q *querycommand.QueryCommand) (*littlebee.AreaTypeList, error) {
	return dao.QueryAreaType(q)
}

func GetAllAreaTypes() (*littlebee.AreaTypeList, error) {
	return dao.GetAllAreaTypes()
}