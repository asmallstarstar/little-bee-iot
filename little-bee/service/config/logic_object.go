package config

import (
	dao "dao/config"
	"fmt"
	"message/littlebee"
	"message/querycommand"

	"google.golang.org/protobuf/types/known/timestamppb"
)

func CreateLogicObject(l *littlebee.LogicObjectExtend, userId int32) (*littlebee.LogicObjectExtend, error) {
	l.LogicObject.LogicObjectId = 0
	l.LogicObject.CreatedAt = timestamppb.Now()
	l.LogicObject.CreatedBy = userId
	if l.LogicConfigure != nil {
		l.LogicConfigure.CreatedAt = timestamppb.Now()
		l.LogicConfigure.CreatedBy = userId
	}
	return dao.CreateLogicObject(l)
}

func RetrieveLogicObject(logicObjectId int32) (*littlebee.LogicObject, error) {
	return dao.RetrieveLogicObject(logicObjectId)
}

func UpdateLogicObject(l *littlebee.LogicObjectExtend, userId int32) error {
	l.LogicObject.UpdatedAt = timestamppb.Now()
	l.LogicObject.UpdatedBy = userId
	if l.LogicConfigure != nil {
		if l.LogicConfigure.ConfigureId == 0 {
			l.LogicConfigure.CreatedAt = timestamppb.Now()
			l.LogicConfigure.CreatedBy = userId
		} else {
			l.LogicConfigure.UpdatedAt = timestamppb.Now()
			l.LogicConfigure.UpdatedBy = userId
		}
	}
	return dao.UpdateLogicObject(l)
}

func DeleteLogicObjectWithFlag(logicObjectId int, userId int32) error {
	l := &littlebee.LogicObject{
		LogicObjectId: int32(logicObjectId),
		IsDeleted:     true,
		DeletedAt:     timestamppb.Now(),
		DeletedBy:     userId,
	}
	return dao.DeleteLogicObjectWithFlag(l)
}

func DeleteLogicObject(logicObjectId int, userId int32) error {
	l := &littlebee.LogicObject{
		LogicObjectId: int32(logicObjectId),
		IsDeleted:     true,
		DeletedAt:     timestamppb.Now(),
		DeletedBy:     userId,
	}
	return dao.DeleteLogicObject(l)
}

func QueryLogicObject(q *querycommand.QueryCommand) (*littlebee.LogicObjectList, error) {
	return dao.QueryLogicObject(q)
}

func GetAllLogicObjects() (*littlebee.LogicObjectList, error) {
	return dao.GetAllLogicObjects()
}

func GetParentIdListByLogicObjectId(logicObjectId int) string {
	if logicObjectId == 0 {
		return "0"
	}
	id := logicObjectId
	idList := ""
	var e error = nil
	var l *littlebee.LogicObject = nil
	for e == nil && id != 0 {
		l, e = RetrieveLogicObject(int32(id))
		id = int(l.ParentLogicObjectId)
		if idList==""{
			idList = fmt.Sprintf("%d", l.LogicObjectId)
		}else{
			idList = fmt.Sprintf("%d.%s", l.LogicObjectId, idList)
		}
	}
	idList = "0." + idList
	return idList
}

func GetParentIdListBySignalId(signalId int) string {
	s, e := RetrieveSignal(int32(signalId))
	if e == nil {
		return GetParentIdListByLogicObjectId(int(s.ParentLogicId))
	}
	return ""
}
