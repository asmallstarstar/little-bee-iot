package router

import (
	"controller/config"

	"github.com/gin-gonic/gin"
)

const URL_ROUTER_ACTION = "/config/action"
const URL_ROUTER_ACTION_BY_ID = "/config/action/:action-id"
const URL_ROUTER_ACTION_QUERY = "/config/action/query"
const URL_ROUTER_ACTION_ALL = "/config/action/all"

func ActionRouter(r *gin.RouterGroup) *gin.RouterGroup {
	r.POST(URL_ROUTER_ACTION, config.CreateAction)
	r.GET(URL_ROUTER_ACTION_BY_ID, config.RetrieveAction)
	r.PUT(URL_ROUTER_ACTION, config.UpdateAction)
	r.DELETE(URL_ROUTER_ACTION_BY_ID, config.DeleteAction)
	r.PUT(URL_ROUTER_ACTION_BY_ID, config.DeleteActionWithFlag)
	r.POST(URL_ROUTER_ACTION_QUERY, config.QueryAction)
	r.GET(URL_ROUTER_ACTION_ALL, config.GetAllActions)
	return r
}