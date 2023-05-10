package config

import (
	"dao/database"
	"message/littlebee"
	"message/querycommand"
	"util/queryclause"
)

func CreateOperationLog(l *littlebee.OperationLog) (*littlebee.OperationLog, error) {
	result := database.DB.Omit("LogId", "OperationTime").Create(l)
	return l, result.Error
}

func QueryOperationLog(q *querycommand.QueryCommand) (*littlebee.OperationLogList, error) {
	var operationLogs []*littlebee.OperationLog
	s := "SELECT * FROM operation_logs "
	w, v := queryclause.QueryCommandString(q)
	s = s + w
	result := database.DB.Raw(s, v...).Scan(&operationLogs)
	return &littlebee.OperationLogList{Items: operationLogs}, result.Error
}
