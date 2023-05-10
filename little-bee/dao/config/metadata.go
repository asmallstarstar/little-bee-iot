package config

import (
	"dao/database"
	"message/littlebee"
	"message/querycommand"
	"util/queryclause"

	"github.com/fatih/structs"
)

func CreateMetadata(l *littlebee.Metadata) (*littlebee.Metadata, error) {
	result := database.DB.Omit("MetadataId", "UpdatedAt", "UpdatedBy", "IsDeleted", "DeletedAt", "DeletedBy").Table("metadatas").Create(l)
	return l, result.Error
}

func RetrieveMetadata(metadataId int32) (*littlebee.Metadata, error) {
	r := &littlebee.Metadata{}
	result := database.DB.Table("metadatas").First(&r, metadataId)
	return r, result.Error
}

func UpdateMetadata(l *littlebee.Metadata) error {
	lMap := structs.Map(l)
	lMap["created_at"] = l.CreatedAt.AsTime().UTC()
	lMap["updated_at"] = l.UpdatedAt.AsTime().UTC()
	lMap["deleted_at"] = l.DeletedAt.AsTime().UTC()
	result := database.DB.Model(l).Omit("MetadataId", "CreatedAt", "CreatedBy", "IsDeleted", "DeletedAt", "DeletedBy").Table("metadatas").Updates(lMap)
	return result.Error
}

func DeleteMetadataWithFlag(l *littlebee.Metadata) error {
	result := database.DB.Model(l).Table("metadatas").Updates(littlebee.Metadata{IsDeleted: l.IsDeleted, DeletedAt: l.DeletedAt, DeletedBy: l.DeletedBy})
	return result.Error
}

func DeleteMetadata(l *littlebee.Metadata) error {
	result := database.DB.Table("metadatas").Delete(l)
	return result.Error
}

func QueryMetadata(q *querycommand.QueryCommand) (*littlebee.MetadataList, error) {
	var metadatas []*littlebee.Metadata
	s := "SELECT * FROM metadatas "
	w, v := queryclause.QueryCommandString(q)
	s = s + w
	result := database.DB.Raw(s, v...).Scan(&metadatas)
	return &littlebee.MetadataList{Items: metadatas}, result.Error
}

func GetAllMetadatas() (*littlebee.MetadataList, error) {
	var metadatas []*littlebee.Metadata
	result := database.DB.Table("metadatas").Find(&metadatas)
	return &littlebee.MetadataList{Items: metadatas}, result.Error
}
