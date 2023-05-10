package router

import (
	"controller/config"

	"github.com/gin-gonic/gin"
)

const URL_ROUTER_AI = "/config/ai"
const URL_ROUTER_AI_BY_ID = "/config/ai/:ai-id"

func AiRouter(r *gin.RouterGroup) *gin.RouterGroup {
	r.POST(URL_ROUTER_AI, config.CreateAi)
	r.GET(URL_ROUTER_AI_BY_ID, config.RetrieveAi)
	r.PUT(URL_ROUTER_AI, config.UpdateAi)
	r.DELETE(URL_ROUTER_AI_BY_ID, config.DeleteAi)
	r.PUT(URL_ROUTER_AI_BY_ID, config.DeleteAiWithFlag)
	return r
}
