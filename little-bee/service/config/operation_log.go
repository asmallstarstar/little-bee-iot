package config

import (
	dao "dao/config"
	"message/littlebee"
	"message/querycommand"
	"service/user"

	"google.golang.org/protobuf/types/known/timestamppb"
)

func CreateOperationLog(l *littlebee.OperationLog) (*littlebee.OperationLog, error) {
	u, _ := user.RetrieveUser(l.UserId)
	if u != nil {
		l.UserNick = u.UserNick
	}
	l.OperationTime = timestamppb.Now()
	return dao.CreateOperationLog(l)
}

func QueryOperationLog(q *querycommand.QueryCommand) (*littlebee.OperationLogList, error) {
	return dao.QueryOperationLog(q)
}
