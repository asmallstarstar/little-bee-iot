package router

import (
	"controller/config"

	"github.com/gin-gonic/gin"
)

const URL_ROUTER_MENU = "/config/menu"
const URL_ROUTER_MENU_BY_ID = "/config/menu/:menu-id"
const URL_ROUTER_MENU_QUERY = "/config/menu/query"
const URL_ROUTER_MENU_ALL = "/config/menu/all"

func MenuRouter(r *gin.RouterGroup) *gin.RouterGroup {
	r.POST(URL_ROUTER_MENU, config.CreateMenu)
	r.GET(URL_ROUTER_MENU_BY_ID, config.RetrieveMenu)
	r.PUT(URL_ROUTER_MENU, config.UpdateMenu)
	r.DELETE(URL_ROUTER_MENU_BY_ID, config.DeleteMenu)
	r.PUT(URL_ROUTER_MENU_BY_ID, config.DeleteMenuWithFlag)
	r.POST(URL_ROUTER_MENU_QUERY, config.QueryMenu)
	r.GET(URL_ROUTER_MENU_ALL, config.GetAllMenus)
	return r
}