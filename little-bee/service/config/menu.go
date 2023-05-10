package config

import (
	dao "dao/config"
	"message/littlebee"
	"message/querycommand"

	"google.golang.org/protobuf/types/known/timestamppb"
)

func CreateMenu(l *littlebee.Menu, userId int32) (*littlebee.Menu, error) {
	l.CreatedAt = timestamppb.Now()
	l.CreatedBy = userId
	return dao.CreateMenu(l)
}

func RetrieveMenu(menuId int32) (*littlebee.Menu, error) {
	return dao.RetrieveMenu(menuId)
}

func UpdateMenu(l *littlebee.Menu, userId int32) error {
	l.UpdatedAt = timestamppb.Now()
	l.UpdatedBy = userId
	return dao.UpdateMenu(l)
}

func DeleteMenuWithFlag(menuId int, userId int32) error {
	l := &littlebee.Menu{
		MenuId:    int32(menuId),
		IsDeleted: true,
		DeletedAt: timestamppb.Now(),
		DeletedBy: userId,
	}
	return dao.DeleteMenuWithFlag(l)
}

func DeleteMenu(menuId int, userId int32) error {
	l := &littlebee.Menu{
		MenuId:    int32(menuId),
		IsDeleted: true,
		DeletedAt: timestamppb.Now(),
		DeletedBy: userId,
	}
	return dao.DeleteMenu(l)
}

func QueryMenu(q *querycommand.QueryCommand) (*littlebee.MenuList, error) {
	return dao.QueryMenu(q)
}

func GetAllMenus() (*littlebee.MenuList, error) {
	return dao.GetAllMenus()
}
