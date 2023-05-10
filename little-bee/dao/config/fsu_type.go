package config

import (
	"dao/database"
	"message/littlebee"
	"message/querycommand"
	"util/queryclause"

	"github.com/fatih/structs"
)

func CreateFsuType(l *littlebee.FsuType) (*littlebee.FsuType, error) {
	result := database.DB.Omit("FsuTypeId", "UpdatedAt", "UpdatedBy", "IsDeleted", "DeletedAt", "DeletedBy").Create(l)
	return l, result.Error
}

func RetrieveFsuType(fsuTypeId int32) (*littlebee.FsuType, error) {
	r := &littlebee.FsuType{}
	result := database.DB.First(&r, fsuTypeId)
	return r, result.Error
}

func UpdateFsuType(l *littlebee.FsuType) error {
	lMap := structs.Map(l)
	lMap["created_at"] = l.CreatedAt.AsTime().UTC()
	lMap["updated_at"] = l.UpdatedAt.AsTime().UTC()
	lMap["deleted_at"] = l.DeletedAt.AsTime().UTC()
	result := database.DB.Model(l).Omit("FsuTypeId", "CreatedAt", "CreatedBy", "IsDeleted", "DeletedAt", "DeletedBy").Updates(lMap)
	return result.Error
}

func DeleteFsuTypeWithFlag(l *littlebee.FsuType) error {
	result := database.DB.Model(l).Updates(littlebee.FsuType{IsDeleted: l.IsDeleted, DeletedAt: l.DeletedAt, DeletedBy: l.DeletedBy})
	return result.Error
}

func DeleteFsuType(l *littlebee.FsuType) error {
	result := database.DB.Delete(l)
	return result.Error
}

func QueryFsuType(q *querycommand.QueryCommand) (*littlebee.FsuTypeList, error) {
	var fsuTypes []*littlebee.FsuType
	s := "SELECT * FROM fsu_types "
	w, v := queryclause.QueryCommandString(q)
	s = s + w
	result := database.DB.Raw(s, v...).Scan(&fsuTypes)
	return &littlebee.FsuTypeList{Items: fsuTypes}, result.Error
}

func GetAllFsuTypes() (*littlebee.FsuTypeList, error) {
	var fsuTypes []*littlebee.FsuType
	result := database.DB.Find(&fsuTypes)
	return &littlebee.FsuTypeList{Items: fsuTypes}, result.Error
}
