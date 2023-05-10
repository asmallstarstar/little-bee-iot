package router

import (
	"controller/config"

	"github.com/gin-gonic/gin"
)

const URL_ROUTER_MENU_ACTION = "/config/menu-action"
const URL_ROUTER_MENU_ACTION_BATCH = "/config/menu-action/batch"
const URL_ROUTER_MENU_ACTION_BY_ID = "/config/menu-action/:menu-action-id"
const URL_ROUTER_MENU_ACTION_QUERY = "/config/menu-action/query"
const URL_ROUTER_MENU_ACTION_ALL = "/config/menu-action/all"

func MenuActionRouter(r *gin.RouterGroup) *gin.RouterGroup {
	r.POST(URL_ROUTER_MENU_ACTION, config.CreateMenuAction)
	r.POST(URL_ROUTER_MENU_ACTION_BATCH, config.BatchCreateMenuAction)
	r.GET(URL_ROUTER_MENU_ACTION_BY_ID, config.RetrieveMenuAction)
	r.PUT(URL_ROUTER_MENU_ACTION, config.UpdateMenuAction)
	r.DELETE(URL_ROUTER_MENU_ACTION_BY_ID, config.DeleteMenuAction)
	r.PUT(URL_ROUTER_MENU_ACTION_BY_ID, config.DeleteMenuActionWithFlag)
	r.POST(URL_ROUTER_MENU_ACTION_QUERY, config.QueryMenuAction)
	r.GET(URL_ROUTER_MENU_ACTION_ALL, config.GetAllMenuActions)
	return r
}
