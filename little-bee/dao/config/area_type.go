package config

import (
	"dao/database"
	"message/littlebee"
	"message/querycommand"
	"util/queryclause"

	"github.com/fatih/structs"
)

func CreateAreaType(l *littlebee.AreaType) (*littlebee.AreaType, error) {
	result := database.DB.Omit("AreaTypeId", "UpdatedAt", "UpdatedBy", "IsDeleted", "DeletedAt", "DeletedBy").Create(l)
	return l, result.Error
}

func RetrieveAreaType(areaTypeId int32) (*littlebee.AreaType, error) {
	r := &littlebee.AreaType{}
	result := database.DB.First(&r, areaTypeId)
	return r, result.Error
}

func UpdateAreaType(l *littlebee.AreaType) error {
	lMap := structs.Map(l)
	lMap["created_at"] = l.CreatedAt.AsTime().UTC()
	lMap["updated_at"] = l.UpdatedAt.AsTime().UTC()
	lMap["deleted_at"] = l.DeletedAt.AsTime().UTC()
	result := database.DB.Model(l).Omit("AreaTypeId", "CreatedAt", "CreatedBy", "IsDeleted", "DeletedAt", "DeletedBy").Updates(lMap)
	return result.Error
}

func DeleteAreaTypeWithFlag(l *littlebee.AreaType) error {
	result := database.DB.Model(l).Updates(littlebee.AreaType{IsDeleted: l.IsDeleted, DeletedAt: l.DeletedAt, DeletedBy: l.DeletedBy})
	return result.Error
}

func DeleteAreaType(l *littlebee.AreaType) error {
	result := database.DB.Delete(l)
	return result.Error
}

func QueryAreaType(q *querycommand.QueryCommand) (*littlebee.AreaTypeList, error) {
	var areaTypes []*littlebee.AreaType
	s := "SELECT * FROM area_types "
	w, v := queryclause.QueryCommandString(q)
	s = s + w
	result := database.DB.Raw(s, v...).Scan(&areaTypes)
	return &littlebee.AreaTypeList{Items: areaTypes}, result.Error
}

func GetAllAreaTypes() (*littlebee.AreaTypeList, error) {
	var areaTypes []*littlebee.AreaType
	result := database.DB.Find(&areaTypes)
	return &littlebee.AreaTypeList{Items: areaTypes}, result.Error
}
