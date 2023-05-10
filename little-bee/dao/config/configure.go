package config

import (
	"dao/database"
	"message/littlebee"
	"message/querycommand"
	"util/queryclause"

	"github.com/fatih/structs"
)

func CreateConfigure(l *littlebee.Configure) (*littlebee.Configure, error) {
	result := database.DB.Omit("ConfigureId", "UpdatedAt", "UpdatedBy", "IsDeleted", "DeletedAt", "DeletedBy").Create(l)
	return l, result.Error
}

func RetrieveConfigure(configureId int32) (*littlebee.Configure, error) {
	r := &littlebee.Configure{}
	result := database.DB.First(&r, configureId)
	return r, result.Error
}

func UpdateConfigure(l *littlebee.Configure) error {
	lMap := structs.Map(l)
	lMap["created_at"] = l.CreatedAt.AsTime().UTC()
	lMap["updated_at"] = l.UpdatedAt.AsTime().UTC()
	lMap["deleted_at"] = l.DeletedAt.AsTime().UTC()
	result := database.DB.Model(l).Omit("ConfigureId", "CreatedAt", "CreatedBy", "IsDeleted", "DeletedAt", "DeletedBy").Updates(lMap)
	return result.Error
}

func DeleteConfigureWithFlag(l *littlebee.Configure) error {
	result := database.DB.Model(l).Updates(littlebee.Configure{IsDeleted: l.IsDeleted, DeletedAt: l.DeletedAt, DeletedBy: l.DeletedBy})
	return result.Error
}

func DeleteConfigure(l *littlebee.Configure) error {
	result := database.DB.Delete(l)
	return result.Error
}

func QueryConfigure(q *querycommand.QueryCommand) (*littlebee.ConfigureList, error) {
	var configures []*littlebee.Configure
	s := "SELECT * FROM configures "
	w, v := queryclause.QueryCommandString(q)
	s = s + w
	result := database.DB.Raw(s, v...).Scan(&configures)
	return &littlebee.ConfigureList{Items: configures}, result.Error
}

func GetAllConfigures() (*littlebee.ConfigureList, error) {
	var configures []*littlebee.Configure
	result := database.DB.Find(&configures)
	return &littlebee.ConfigureList{Items: configures}, result.Error
}
