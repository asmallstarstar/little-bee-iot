package router

import (
	"controller/config"

	"github.com/gin-gonic/gin"
)

const URL_ROUTER_SIGNAL_AI = "/config/signal-ai"
const URL_ROUTER_SIGNAL_AI_BY_ID = "/config/signal-ai/:signal-ai-id"
const URL_ROUTER_SIGNAL_AI_QUERY = "/config/signal-ai/query"
const URL_ROUTER_SIGNAL_AI_ALL = "/config/signal-ai/all"

func SignalAiRouter(r *gin.RouterGroup) *gin.RouterGroup {
	r.POST(URL_ROUTER_SIGNAL_AI, config.CreateSignalAi)
	r.GET(URL_ROUTER_SIGNAL_AI_BY_ID, config.RetrieveSignalAi)
	r.PUT(URL_ROUTER_SIGNAL_AI, config.UpdateSignalAi)
	r.DELETE(URL_ROUTER_SIGNAL_AI_BY_ID, config.DeleteSignalAi)
	r.POST(URL_ROUTER_SIGNAL_AI_QUERY, config.QuerySignalAi)
	r.GET(URL_ROUTER_SIGNAL_AI_ALL, config.GetAllSignalAis)
	return r
}