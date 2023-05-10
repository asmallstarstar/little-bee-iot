package config

import (
	dao "dao/config"
	"message/littlebee"
	"message/querycommand"

	"google.golang.org/protobuf/types/known/timestamppb"
)

func CreateActionGroup(l *littlebee.ActionGroup, userId int32) (*littlebee.ActionGroup, error) {
	l.CreatedAt = timestamppb.Now()
	l.CreatedBy = userId
	return dao.CreateActionGroup(l)
}

func RetrieveActionGroup(actionGroupId int32) (*littlebee.ActionGroup, error) {
	return dao.RetrieveActionGroup(actionGroupId)
}

func UpdateActionGroup(l *littlebee.ActionGroup, userId int32) error {
	l.UpdatedAt = timestamppb.Now()
	l.UpdatedBy = userId
	return dao.UpdateActionGroup(l)
}

func DeleteActionGroupWithFlag(actionGroupId int, userId int32) error {
	l := &littlebee.ActionGroup{
		ActionGroupId: int32(actionGroupId),
		IsDeleted:     true,
		DeletedAt:     timestamppb.Now(),
		DeletedBy:     userId,
	}
	return dao.DeleteActionGroupWithFlag(l)
}

func DeleteActionGroup(actionGroupId int, userId int32) error {
	l := &littlebee.ActionGroup{
		ActionGroupId: int32(actionGroupId),
		IsDeleted:     true,
		DeletedAt:     timestamppb.Now(),
		DeletedBy:     userId,
	}
	return dao.DeleteActionGroup(l)
}

func QueryActionGroup(q *querycommand.QueryCommand) (*littlebee.ActionGroupList, error) {
	return dao.QueryActionGroup(q)
}

func GetAllActionGroups() (*littlebee.ActionGroupList, error) {
	return dao.GetAllActionGroups()
}
