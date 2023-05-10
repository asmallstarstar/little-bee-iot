package router

import (
	"controller/config"

	"github.com/gin-gonic/gin"
)

const URL_ROUTER_AGENT = "/config/agent"
const URL_ROUTER_AGENT_BY_ID = "/config/agent/:agent-id"
const URL_ROUTER_AGENT_QUERY = "/config/agent/query"
const URL_ROUTER_AGENT_ALL = "/config/agent/all"

func AgentRouter(r *gin.RouterGroup) *gin.RouterGroup {
	r.POST(URL_ROUTER_AGENT, config.CreateAgent)
	r.GET(URL_ROUTER_AGENT_BY_ID, config.RetrieveAgent)
	r.PUT(URL_ROUTER_AGENT, config.UpdateAgent)
	r.DELETE(URL_ROUTER_AGENT_BY_ID, config.DeleteAgent)
	r.PUT(URL_ROUTER_AGENT_BY_ID, config.DeleteAgentWithFlag)
	r.POST(URL_ROUTER_AGENT_QUERY, config.QueryAgent)
	r.GET(URL_ROUTER_AGENT_ALL, config.GetAllAgents)
	return r
}