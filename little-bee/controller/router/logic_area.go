package router

import (
	"controller/config"

	"github.com/gin-gonic/gin"
)

const URL_ROUTER_LOGIC_AREA = "/config/logic-area"
const URL_ROUTER_LOGIC_AREA_BY_ID = "/config/logic-area/:logic-area-id"
const URL_ROUTER_LOGIC_AREA_QUERY = "/config/logic-area/query"
const URL_ROUTER_LOGIC_AREA_ALL = "/config/logic-area/all"

func LogicAreaRouter(r *gin.RouterGroup) *gin.RouterGroup {
	r.POST(URL_ROUTER_LOGIC_AREA, config.CreateLogicArea)
	r.GET(URL_ROUTER_LOGIC_AREA_BY_ID, config.RetrieveLogicArea)
	r.PUT(URL_ROUTER_LOGIC_AREA, config.UpdateLogicArea)
	r.DELETE(URL_ROUTER_LOGIC_AREA_BY_ID, config.DeleteLogicArea)
	r.POST(URL_ROUTER_LOGIC_AREA_QUERY, config.QueryLogicArea)
	r.GET(URL_ROUTER_LOGIC_AREA_ALL, config.GetAllLogicAreas)
	return r
}