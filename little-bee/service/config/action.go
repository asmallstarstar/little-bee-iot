package config

import (
	dao "dao/config"
	"message/littlebee"
	"message/querycommand"

	"google.golang.org/protobuf/types/known/timestamppb"
)

func CreateAction(l *littlebee.Action, userId int32) (*littlebee.Action, error) {
	l.ActionId = 0
	l.CreatedAt = timestamppb.Now()
	l.CreatedBy = userId
	return dao.CreateAction(l)
}

func RetrieveAction(actionId int32) (*littlebee.Action, error) {
	return dao.RetrieveAction(actionId)
}

func GetActionByName(actionName string) (*littlebee.Action, error) {
	return dao.GetActionByName(actionName)
}

func UpdateAction(l *littlebee.Action, userId int32) error {
	l.UpdatedAt = timestamppb.Now()
	l.UpdatedBy = userId
	return dao.UpdateAction(l)
}

func DeleteActionWithFlag(actionId int, userId int32) error {
	l := &littlebee.Action{
		ActionId:  int32(actionId),
		IsDeleted: true,
		DeletedAt: timestamppb.Now(),
		DeletedBy: userId,
	}
	return dao.DeleteActionWithFlag(l)
}

func DeleteAction(actionId int, userId int32) error {
	l := &littlebee.Action{
		ActionId:  int32(actionId),
		IsDeleted: true,
		DeletedAt: timestamppb.Now(),
		DeletedBy: userId,
	}
	return dao.DeleteAction(l)
}

func QueryAction(q *querycommand.QueryCommand) (*littlebee.ActionList, error) {
	return dao.QueryAction(q)
}

func GetAllActions() (*littlebee.ActionList, error) {
	return dao.GetAllActions()
}
