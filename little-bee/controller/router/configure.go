package router

import (
	"controller/config"

	"github.com/gin-gonic/gin"
)

const URL_ROUTER_CONFIGURE = "/config/configure"
const URL_ROUTER_CONFIGURE_BY_ID = "/config/configure/:configure-id"
const URL_ROUTER_CONFIGURE_QUERY = "/config/configure/query"
const URL_ROUTER_CONFIGURE_ALL = "/config/configure/all"

func ConfigureRouter(r *gin.RouterGroup) *gin.RouterGroup {
	r.POST(URL_ROUTER_CONFIGURE, config.CreateConfigure)
	r.GET(URL_ROUTER_CONFIGURE_BY_ID, config.RetrieveConfigure)
	r.PUT(URL_ROUTER_CONFIGURE, config.UpdateConfigure)
	r.DELETE(URL_ROUTER_CONFIGURE_BY_ID, config.DeleteConfigure)
	r.PUT(URL_ROUTER_CONFIGURE_BY_ID, config.DeleteConfigureWithFlag)
	r.POST(URL_ROUTER_CONFIGURE_QUERY, config.QueryConfigure)
	r.GET(URL_ROUTER_CONFIGURE_ALL, config.GetAllConfigures)
	return r
}