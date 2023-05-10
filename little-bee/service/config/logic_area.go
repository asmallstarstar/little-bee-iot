package config

import (
	dao "dao/config"
	"message/littlebee"
	"message/querycommand"
)

func CreateLogicArea(l *littlebee.LogicArea, userId int32) (*littlebee.LogicArea, error) {
	return dao.CreateLogicArea(l)
}

func RetrieveLogicArea(logicAreaId int32) (*littlebee.LogicArea, error) {
	return dao.RetrieveLogicArea(logicAreaId)
}

func UpdateLogicArea(l *littlebee.LogicArea, userId int32) error {
	return dao.UpdateLogicArea(l)
}

func DeleteLogicArea(logicAreaId int, userId int32) error {
	l := &littlebee.LogicArea{
		LogicObjectId: int32(logicAreaId),
	}
	return dao.DeleteLogicArea(l)
}

func QueryLogicArea(q *querycommand.QueryCommand) (*littlebee.LogicAreaList, error) {
	return dao.QueryLogicArea(q)
}

func GetAllLogicAreas() (*littlebee.LogicAreaList, error) {
	return dao.GetAllLogicAreas()
}
