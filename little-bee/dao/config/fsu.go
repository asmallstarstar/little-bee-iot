package config

import (
	"dao/database"
	"message/littlebee"
	"message/querycommand"
	"util/queryclause"

	"github.com/fatih/structs"
)

func CreateFsu(l *littlebee.Fsu) (*littlebee.Fsu, error) {
	result := database.DB.Omit("FsuId", "UpdatedAt", "UpdatedBy", "IsDeleted", "DeletedAt", "DeletedBy").Create(l)
	return l, result.Error
}

func RetrieveFsu(fsuId int32) (*littlebee.Fsu, error) {
	r := &littlebee.Fsu{}
	result := database.DB.First(&r, fsuId)
	return r, result.Error
}

func UpdateFsu(l *littlebee.Fsu) error {
	lMap := structs.Map(l)
	lMap["created_at"] = l.CreatedAt.AsTime().UTC()
	lMap["updated_at"] = l.UpdatedAt.AsTime().UTC()
	lMap["deleted_at"] = l.DeletedAt.AsTime().UTC()
	result := database.DB.Model(l).Omit("FsuId", "CreatedAt", "CreatedBy", "IsDeleted", "DeletedAt", "DeletedBy").Updates(lMap)
	return result.Error
}

func DeleteFsuWithFlag(l *littlebee.Fsu) error {
	result := database.DB.Model(l).Updates(littlebee.Fsu{IsDeleted: l.IsDeleted, DeletedAt: l.DeletedAt, DeletedBy: l.DeletedBy})
	return result.Error
}

func DeleteFsu(l *littlebee.Fsu) error {
	result := database.DB.Delete(l)
	return result.Error
}

func QueryFsu(q *querycommand.QueryCommand) (*littlebee.FsuList, error) {
	var fsus []*littlebee.Fsu
	s := "SELECT * FROM fsus "
	w, v := queryclause.QueryCommandString(q)
	s = s + w
	result := database.DB.Raw(s, v...).Scan(&fsus)
	return &littlebee.FsuList{Items: fsus}, result.Error
}

func GetAllFsus() (*littlebee.FsuList, error) {
	var fsus []*littlebee.Fsu
	result := database.DB.Find(&fsus)
	return &littlebee.FsuList{Items: fsus}, result.Error
}
