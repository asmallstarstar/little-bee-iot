package config

import (
	"dao/database"
	"message/littlebee"
	"message/querycommand"
	"util/queryclause"

	"github.com/fatih/structs"
)

func CreateActionGroup(l *littlebee.ActionGroup) (*littlebee.ActionGroup, error) {
	result := database.DB.Omit("UpdatedAt", "UpdatedBy", "IsDeleted", "DeletedAt", "DeletedBy").Create(l)
	return l, result.Error
}

func RetrieveActionGroup(actionGroupId int32) (*littlebee.ActionGroup, error) {
	r := &littlebee.ActionGroup{}
	result := database.DB.First(&r, actionGroupId)
	return r, result.Error
}

func UpdateActionGroup(l *littlebee.ActionGroup) error {
	lMap := structs.Map(l)
	lMap["created_at"] = l.CreatedAt.AsTime().UTC()
	lMap["updated_at"] = l.UpdatedAt.AsTime().UTC()
	lMap["deleted_at"] = l.DeletedAt.AsTime().UTC()
	result := database.DB.Model(l).Omit("ActionGroupId", "CreatedAt", "CreatedBy", "IsDeleted", "DeletedAt", "DeletedBy").Updates(lMap)
	return result.Error
}

func DeleteActionGroupWithFlag(l *littlebee.ActionGroup) error {
	result := database.DB.Model(l).Updates(littlebee.ActionGroup{IsDeleted: l.IsDeleted, DeletedAt: l.DeletedAt, DeletedBy: l.DeletedBy})
	return result.Error
}

func DeleteActionGroup(l *littlebee.ActionGroup) error {
	result := database.DB.Delete(l)
	return result.Error
}

func QueryActionGroup(q *querycommand.QueryCommand) (*littlebee.ActionGroupList, error) {
	var actionGroups []*littlebee.ActionGroup
	s := "SELECT * FROM action_groups "
	w, v := queryclause.QueryCommandString(q)
	s = s + w
	result := database.DB.Raw(s, v...).Scan(&actionGroups)
	return &littlebee.ActionGroupList{Items: actionGroups}, result.Error
}

func GetAllActionGroups() (*littlebee.ActionGroupList, error) {
	var actionGroups []*littlebee.ActionGroup
	result := database.DB.Find(&actionGroups)
	return &littlebee.ActionGroupList{Items: actionGroups}, result.Error
}
