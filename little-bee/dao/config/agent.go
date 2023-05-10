package config

import (
	"dao/database"
	"message/littlebee"
	"message/querycommand"
	"util/queryclause"

	"github.com/fatih/structs"
)

func CreateAgent(l *littlebee.Agent) (*littlebee.Agent, error) {
	result := database.DB.Omit("AgentId", "UpdatedAt", "UpdatedBy", "IsDeleted", "DeletedAt", "DeletedBy").Create(l)
	return l, result.Error
}

func RetrieveAgent(agentId int32) (*littlebee.Agent, error) {
	r := &littlebee.Agent{}
	result := database.DB.First(&r, agentId)
	return r, result.Error
}

func UpdateAgent(l *littlebee.Agent) error {
	lMap := structs.Map(l)
	lMap["created_at"] = l.CreatedAt.AsTime().UTC()
	lMap["updated_at"] = l.UpdatedAt.AsTime().UTC()
	lMap["deleted_at"] = l.DeletedAt.AsTime().UTC()
	result := database.DB.Model(l).Omit("AgentId", "CreatedAt", "CreatedBy", "IsDeleted", "DeletedAt", "DeletedBy").Updates(lMap)
	return result.Error
}

func DeleteAgentWithFlag(l *littlebee.Agent) error {
	result := database.DB.Model(l).Updates(littlebee.Agent{IsDeleted: l.IsDeleted, DeletedAt: l.DeletedAt, DeletedBy: l.DeletedBy})
	return result.Error
}

func DeleteAgent(l *littlebee.Agent) error {
	result := database.DB.Delete(l)
	return result.Error
}

func QueryAgent(q *querycommand.QueryCommand) (*littlebee.AgentList, error) {
	var agents []*littlebee.Agent
	s := "SELECT * FROM agents "
	w, v := queryclause.QueryCommandString(q)
	s = s + w
	result := database.DB.Raw(s, v...).Scan(&agents)
	return &littlebee.AgentList{Items: agents}, result.Error
}

func GetAllAgents() (*littlebee.AgentList, error) {
	var agents []*littlebee.Agent
	result := database.DB.Find(&agents)
	return &littlebee.AgentList{Items: agents}, result.Error
}
