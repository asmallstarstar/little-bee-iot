package config

import (
	dao "dao/config"
	"message/littlebee"
	"message/querycommand"

	"google.golang.org/protobuf/types/known/timestamppb"
)

func CreateMetadata(l *littlebee.Metadata, userId int32) (*littlebee.Metadata, error) {
	l.MetadataId = 0
	l.CreatedAt = timestamppb.Now()
	l.CreatedBy = userId
	return dao.CreateMetadata(l)
}

func RetrieveMetadata(metadataId int32) (*littlebee.Metadata, error) {
	return dao.RetrieveMetadata(metadataId)
}

func UpdateMetadata(l *littlebee.Metadata, userId int32) error {
	l.UpdatedAt = timestamppb.Now()
	l.UpdatedBy = userId
	return dao.UpdateMetadata(l)
}

func DeleteMetadataWithFlag(metadataId int, userId int32) error {
	l := &littlebee.Metadata{
		MetadataId: int32(metadataId),
		IsDeleted:     true,
		DeletedAt:     timestamppb.Now(),
		DeletedBy:     userId,
	}
	return dao.DeleteMetadataWithFlag(l)
}

func DeleteMetadata(metadataId int, userId int32) error {
	l := &littlebee.Metadata{
		MetadataId: int32(metadataId),
		IsDeleted:     true,
		DeletedAt:     timestamppb.Now(),
		DeletedBy:     userId,
	}
	return dao.DeleteMetadata(l)
}

func QueryMetadata(q *querycommand.QueryCommand) (*littlebee.MetadataList, error) {
	return dao.QueryMetadata(q)
}

func GetAllMetadatas() (*littlebee.MetadataList, error) {
	return dao.GetAllMetadatas()
}