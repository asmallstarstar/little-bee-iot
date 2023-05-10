package config

import (
	dao "dao/config"
	"message/littlebee"
	"message/querycommand"

	"google.golang.org/protobuf/types/known/timestamppb"
)

func CreateAgent(l *littlebee.Agent, userId int32) (*littlebee.Agent, error) {
	l.AgentId = 0
	l.CreatedAt = timestamppb.Now()
	l.CreatedBy = userId
	return dao.CreateAgent(l)
}

func RetrieveAgent(agentId int32) (*littlebee.Agent, error) {
	return dao.RetrieveAgent(agentId)
}

func UpdateAgent(l *littlebee.Agent, userId int32) error {
	l.UpdatedAt = timestamppb.Now()
	l.UpdatedBy = userId
	return dao.UpdateAgent(l)
}

func DeleteAgentWithFlag(agentId int, userId int32) error {
	l := &littlebee.Agent{
		AgentId: int32(agentId),
		IsDeleted:     true,
		DeletedAt:     timestamppb.Now(),
		DeletedBy:     userId,
	}
	return dao.DeleteAgentWithFlag(l)
}

func DeleteAgent(agentId int, userId int32) error {
	l := &littlebee.Agent{
		AgentId: int32(agentId),
		IsDeleted:     true,
		DeletedAt:     timestamppb.Now(),
		DeletedBy:     userId,
	}
	return dao.DeleteAgent(l)
}

func QueryAgent(q *querycommand.QueryCommand) (*littlebee.AgentList, error) {
	return dao.QueryAgent(q)
}

func GetAllAgents() (*littlebee.AgentList, error) {
	return dao.GetAllAgents()
}