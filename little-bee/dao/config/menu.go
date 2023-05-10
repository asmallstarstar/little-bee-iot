package config

import (
	"dao/database"
	"message/littlebee"
	"message/querycommand"
	"util/queryclause"

	"github.com/fatih/structs"
)

func CreateMenu(l *littlebee.Menu) (*littlebee.Menu, error) {
	result := database.DB.Omit("IsShow", "UpdatedAt", "UpdatedBy", "IsDeleted", "DeletedAt", "DeletedBy").Create(l)
	return l, result.Error
}

func RetrieveMenu(menuId int32) (*littlebee.Menu, error) {
	r := &littlebee.Menu{}
	result := database.DB.First(&r, menuId)
	return r, result.Error
}

func UpdateMenu(l *littlebee.Menu) error {
	lMap := structs.Map(l)
	lMap["created_at"] = l.CreatedAt.AsTime().UTC()
	lMap["updated_at"] = l.UpdatedAt.AsTime().UTC()
	lMap["deleted_at"] = l.DeletedAt.AsTime().UTC()
	result := database.DB.Model(l).Omit("IsShow", "MenuId", "CreatedAt", "CreatedBy", "IsDeleted", "DeletedAt", "DeletedBy").Updates(lMap)
	return result.Error
}

func DeleteMenuWithFlag(l *littlebee.Menu) error {
	result := database.DB.Model(l).Updates(littlebee.Menu{IsDeleted: l.IsDeleted, DeletedAt: l.DeletedAt, DeletedBy: l.DeletedBy})
	return result.Error
}

func DeleteMenu(l *littlebee.Menu) error {
	result := database.DB.Delete(l)
	return result.Error
}

func QueryMenu(q *querycommand.QueryCommand) (*littlebee.MenuList, error) {
	var menus []*littlebee.Menu
	s := "SELECT * FROM menus "
	w, v := queryclause.QueryCommandString(q)
	s = s + w
	result := database.DB.Raw(s, v...).Scan(&menus)
	return &littlebee.MenuList{Items: menus}, result.Error
}

func GetAllMenus() (*littlebee.MenuList, error) {
	var menus []*littlebee.Menu
	result := database.DB.Find(&menus)
	return &littlebee.MenuList{Items: menus}, result.Error
}
