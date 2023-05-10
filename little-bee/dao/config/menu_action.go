package config

import (
	"dao/database"
	"message"
	"message/littlebee"
	"message/querycommand"
	"util/queryclause"

	"github.com/fatih/structs"
	"google.golang.org/protobuf/types/known/timestamppb"
	"gorm.io/gorm"
)

func BatchCreateMenuAction(l *message.ActionWithMenu, userId int32) error {
	return database.DB.Transaction(func(tx *gorm.DB) error {
		r := database.DB.Where("menu_name = ?", l.MenuName).Delete(&littlebee.MenuAction{})
		if r.Error != nil {
			return r.Error
		}

		for i := 0; i < len(l.ActionNameItems); i++ {
			ma := &littlebee.MenuAction{
				MenuName:   l.MenuName,
				ActionName: *l.ActionNameItems[i],
				CreatedBy:  userId,
				CreatedAt:  timestamppb.Now(),
			}
			result := database.DB.Omit("MenuActionId", "UpdatedAt", "UpdatedBy", "IsDeleted", "DeletedAt", "DeletedBy").Create(ma)
			if result.Error != nil {
				return result.Error
			}
		}

		return nil
	})
}

func CreateMenuAction(l *littlebee.MenuAction) (*littlebee.MenuAction, error) {
	result := database.DB.Omit("MenuActionId", "UpdatedAt", "UpdatedBy", "IsDeleted", "DeletedAt", "DeletedBy").Create(l)
	return l, result.Error
}

func RetrieveMenuAction(menuActionId int32) (*littlebee.MenuAction, error) {
	r := &littlebee.MenuAction{}
	result := database.DB.First(&r, menuActionId)
	return r, result.Error
}

func UpdateMenuAction(l *littlebee.MenuAction) error {
	lMap := structs.Map(l)
	lMap["created_at"] = l.CreatedAt.AsTime().UTC()
	lMap["updated_at"] = l.UpdatedAt.AsTime().UTC()
	lMap["deleted_at"] = l.DeletedAt.AsTime().UTC()
	result := database.DB.Model(l).Omit("MenuActionId", "CreatedAt", "CreatedBy", "IsDeleted", "DeletedAt", "DeletedBy").Updates(lMap)
	return result.Error
}

func DeleteMenuActionWithFlag(l *littlebee.MenuAction) error {
	result := database.DB.Model(l).Updates(littlebee.MenuAction{IsDeleted: l.IsDeleted, DeletedAt: l.DeletedAt, DeletedBy: l.DeletedBy})
	return result.Error
}

func DeleteMenuAction(l *littlebee.MenuAction) error {
	result := database.DB.Delete(l)
	return result.Error
}

func QueryMenuAction(q *querycommand.QueryCommand) (*littlebee.MenuActionList, error) {
	var menuActions []*littlebee.MenuAction
	s := "SELECT * FROM menu_actions "
	w, v := queryclause.QueryCommandString(q)
	s = s + w
	result := database.DB.Raw(s, v...).Scan(&menuActions)
	return &littlebee.MenuActionList{Items: menuActions}, result.Error
}

func GetAllMenuActions() (*littlebee.MenuActionList, error) {
	var menuActions []*littlebee.MenuAction
	result := database.DB.Find(&menuActions)
	return &littlebee.MenuActionList{Items: menuActions}, result.Error
}
