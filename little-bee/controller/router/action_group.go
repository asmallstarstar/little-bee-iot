package router

import (
	"controller/config"

	"github.com/gin-gonic/gin"
)

const URL_ROUTER_ACTION_GROUP = "/config/action-group"
const URL_ROUTER_ACTION_GROUP_BY_ID = "/config/action-group/:action-group-id"
const URL_ROUTER_ACTION_GROUP_QUERY = "/config/action-group/query"
const URL_ROUTER_ACTION_GROUP_ALL = "/config/action-group/all"

func ActionGroupRouter(r *gin.RouterGroup) *gin.RouterGroup {
	r.POST(URL_ROUTER_ACTION_GROUP, config.CreateActionGroup)
	r.GET(URL_ROUTER_ACTION_GROUP_BY_ID, config.RetrieveActionGroup)
	r.PUT(URL_ROUTER_ACTION_GROUP, config.UpdateActionGroup)
	r.DELETE(URL_ROUTER_ACTION_GROUP_BY_ID, config.DeleteActionGroup)
	r.PUT(URL_ROUTER_ACTION_GROUP_BY_ID, config.DeleteActionGroupWithFlag)
	r.POST(URL_ROUTER_ACTION_GROUP_QUERY, config.QueryActionGroup)
	r.GET(URL_ROUTER_ACTION_GROUP_ALL, config.GetAllActionGroups)
	return r
}