package router

import (
	"controller/config"

	"github.com/gin-gonic/gin"
)

const URL_ROUTER_AREA_TYPE = "/config/area-type"
const URL_ROUTER_AREA_TYPE_BY_ID = "/config/area-type/:area-type-id"
const URL_ROUTER_AREA_TYPE_QUERY = "/config/area-type/query"
const URL_ROUTER_AREA_TYPE_ALL = "/config/area-type/all"

func AreaTypeRouter(r *gin.RouterGroup) *gin.RouterGroup {
	r.POST(URL_ROUTER_AREA_TYPE, config.CreateAreaType)
	r.GET(URL_ROUTER_AREA_TYPE_BY_ID, config.RetrieveAreaType)
	r.PUT(URL_ROUTER_AREA_TYPE, config.UpdateAreaType)
	r.DELETE(URL_ROUTER_AREA_TYPE_BY_ID, config.DeleteAreaType)
	r.PUT(URL_ROUTER_AREA_TYPE_BY_ID, config.DeleteAreaTypeWithFlag)
	r.POST(URL_ROUTER_AREA_TYPE_QUERY, config.QueryAreaType)
	r.GET(URL_ROUTER_AREA_TYPE_ALL, config.GetAllAreaTypes)
	return r
}