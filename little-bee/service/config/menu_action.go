package config

import (
	dao "dao/config"
	"message"
	"message/littlebee"
	"message/querycommand"

	"google.golang.org/protobuf/types/known/timestamppb"
)

func BatchCreateMenuAction(l *message.ActionWithMenu, userId int32) error {
	return dao.BatchCreateMenuAction(l, userId)
}

func CreateMenuAction(l *littlebee.MenuAction, userId int32) (*littlebee.MenuAction, error) {
	l.MenuActionId = 0
	l.CreatedAt = timestamppb.Now()
	l.CreatedBy = userId
	return dao.CreateMenuAction(l)
}

func RetrieveMenuAction(menuActionId int32) (*littlebee.MenuAction, error) {
	return dao.RetrieveMenuAction(menuActionId)
}

func UpdateMenuAction(l *littlebee.MenuAction, userId int32) error {
	l.UpdatedAt = timestamppb.Now()
	l.UpdatedBy = userId
	return dao.UpdateMenuAction(l)
}

func DeleteMenuActionWithFlag(menuActionId int, userId int32) error {
	l := &littlebee.MenuAction{
		MenuActionId: int32(menuActionId),
		IsDeleted:    true,
		DeletedAt:    timestamppb.Now(),
		DeletedBy:    userId,
	}
	return dao.DeleteMenuActionWithFlag(l)
}

func DeleteMenuAction(menuActionId int, userId int32) error {
	l := &littlebee.MenuAction{
		MenuActionId: int32(menuActionId),
		IsDeleted:    true,
		DeletedAt:    timestamppb.Now(),
		DeletedBy:    userId,
	}
	return dao.DeleteMenuAction(l)
}

func QueryMenuAction(q *querycommand.QueryCommand) (*littlebee.MenuActionList, error) {
	return dao.QueryMenuAction(q)
}

func GetAllMenuActions() (*littlebee.MenuActionList, error) {
	return dao.GetAllMenuActions()
}
