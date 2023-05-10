package router

import (
	"controller/config"

	"github.com/gin-gonic/gin"
)

const URL_ROUTER_LOGIC_OBJECT = "/config/logic-object"
const URL_ROUTER_LOGIC_OBJECT_BY_ID = "/config/logic-object/:logic-object-id"
const URL_ROUTER_LOGIC_OBJECT_QUERY = "/config/logic-object/query"
const URL_ROUTER_LOGIC_OBJECT_ALL = "/config/logic-object/all"

const URL_ROUTER_PATH_BY_LOGIC_OBJECT_ID = "/config/path/logic-object/:logic-object-id"
const URL_ROUTER_PATH_BY_SIGNAL_ID = "/config/path/signal/:signal-id"

func LogicObjectRouter(r *gin.RouterGroup) *gin.RouterGroup {
	r.POST(URL_ROUTER_LOGIC_OBJECT, config.CreateLogicObject)
	r.GET(URL_ROUTER_LOGIC_OBJECT_BY_ID, config.RetrieveLogicObject)
	r.PUT(URL_ROUTER_LOGIC_OBJECT, config.UpdateLogicObject)
	r.DELETE(URL_ROUTER_LOGIC_OBJECT_BY_ID, config.DeleteLogicObject)
	r.PUT(URL_ROUTER_LOGIC_OBJECT_BY_ID, config.DeleteLogicObjectWithFlag)
	r.POST(URL_ROUTER_LOGIC_OBJECT_QUERY, config.QueryLogicObject)
	r.GET(URL_ROUTER_LOGIC_OBJECT_ALL, config.GetAllLogicObjects)
	r.GET(URL_ROUTER_PATH_BY_LOGIC_OBJECT_ID,config.GetPathByLogicObjectId)
	r.GET(URL_ROUTER_PATH_BY_SIGNAL_ID,config.GetPathBySignalId)
	return r
}