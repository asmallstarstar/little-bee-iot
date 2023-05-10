package config

import (
	"dao/database"
	"message/littlebee"
	"message/querycommand"
	"util/queryclause"

	"github.com/fatih/structs"
)

func CreateAction(l *littlebee.Action) (*littlebee.Action, error) {
	result := database.DB.Omit("ActionId", "UpdatedAt", "UpdatedBy", "IsDeleted", "DeletedAt", "DeletedBy").Create(l)
	return l, result.Error
}

func RetrieveAction(actionId int32) (*littlebee.Action, error) {
	r := &littlebee.Action{}
	result := database.DB.First(&r, actionId)
	return r, result.Error
}

func GetActionByName(actionName string) (*littlebee.Action, error) {
	r := &littlebee.Action{}
	result := database.DB.Where("action_name = ?", actionName).Find(&r)

	return r, result.Error
}

func UpdateAction(l *littlebee.Action) error {
	lMap := structs.Map(l)
	lMap["created_at"] = l.CreatedAt.AsTime().UTC()
	lMap["updated_at"] = l.UpdatedAt.AsTime().UTC()
	lMap["deleted_at"] = l.DeletedAt.AsTime().UTC()
	result := database.DB.Model(l).Omit("ActionId", "CreatedAt", "CreatedBy", "IsDeleted", "DeletedAt", "DeletedBy").Updates(lMap)
	return result.Error
}

func DeleteActionWithFlag(l *littlebee.Action) error {
	result := database.DB.Model(l).Updates(littlebee.Action{IsDeleted: l.IsDeleted, DeletedAt: l.DeletedAt, DeletedBy: l.DeletedBy})
	return result.Error
}

func DeleteAction(l *littlebee.Action) error {
	result := database.DB.Delete(l)
	return result.Error
}

func QueryAction(q *querycommand.QueryCommand) (*littlebee.ActionList, error) {
	var actions []*littlebee.Action
	s := "SELECT * FROM actions "
	w, v := queryclause.QueryCommandString(q)
	s = s + w
	result := database.DB.Raw(s, v...).Scan(&actions)
	return &littlebee.ActionList{Items: actions}, result.Error
}

func GetAllActions() (*littlebee.ActionList, error) {
	var actions []*littlebee.Action
	result := database.DB.Find(&actions)
	return &littlebee.ActionList{Items: actions}, result.Error
}
