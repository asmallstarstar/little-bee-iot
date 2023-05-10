package router

import (
	"controller/config"

	"github.com/gin-gonic/gin"
)

const URL_ROUTER_DI = "/config/di"
const URL_ROUTER_DI_BY_ID = "/config/di/:di-id"
const URL_ROUTER_DI_QUERY = "/config/di/query"
const URL_ROUTER_DI_ALL = "/config/di/all"

func DiRouter(r *gin.RouterGroup) *gin.RouterGroup {
	r.POST(URL_ROUTER_DI, config.CreateDi)
	r.GET(URL_ROUTER_DI_BY_ID, config.RetrieveDi)
	r.PUT(URL_ROUTER_DI, config.UpdateDi)
	r.DELETE(URL_ROUTER_DI_BY_ID, config.DeleteDi)
	r.PUT(URL_ROUTER_DI_BY_ID, config.DeleteDiWithFlag)
	return r
}
