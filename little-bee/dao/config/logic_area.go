package config

import (
	"dao/database"
	"message/littlebee"
	"message/querycommand"
	"util/queryclause"

	"github.com/fatih/structs"
)

func CreateLogicArea(l *littlebee.LogicArea) (*littlebee.LogicArea, error) {
	result := database.DB.Create(l)
	return l, result.Error
}

func RetrieveLogicArea(logicObjectId int32) (*littlebee.LogicArea, error) {
	r := &littlebee.LogicArea{}
	result := database.DB.First(&r, logicObjectId)
	return r, result.Error
}

func UpdateLogicArea(l *littlebee.LogicArea) error {
	lMap := structs.Map(l)
	result := database.DB.Model(l).Omit("LogicAreaId").Updates(lMap)
	return result.Error
}

func DeleteLogicArea(l *littlebee.LogicArea) error {
	result := database.DB.Delete(l)
	return result.Error
}

func QueryLogicArea(q *querycommand.QueryCommand) (*littlebee.LogicAreaList, error) {
	var logicAreas []*littlebee.LogicArea
	s := "SELECT * FROM logic_areas "
	w, v := queryclause.QueryCommandString(q)
	s = s + w
	result := database.DB.Raw(s, v...).Scan(&logicAreas)
	return &littlebee.LogicAreaList{Items: logicAreas}, result.Error
}

func GetAllLogicAreas() (*littlebee.LogicAreaList, error) {
	var logicAreas []*littlebee.LogicArea
	result := database.DB.Find(&logicAreas)
	return &littlebee.LogicAreaList{Items: logicAreas}, result.Error
}
