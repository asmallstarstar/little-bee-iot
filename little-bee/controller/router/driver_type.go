package router

import (
	"controller/config"

	"github.com/gin-gonic/gin"
)

const URL_ROUTER_DRIVER_TYPE = "/config/driver-type"
const URL_ROUTER_DRIVER_TYPE_BY_ID = "/config/driver-type/:driver-type-id"
const URL_ROUTER_DRIVER_TYPE_QUERY = "/config/driver-type/query"
const URL_ROUTER_DRIVER_TYPE_ALL = "/config/driver-type/all"

func DriverTypeRouter(r *gin.RouterGroup) *gin.RouterGroup {
	r.POST(URL_ROUTER_DRIVER_TYPE, config.CreateDriverType)
	r.GET(URL_ROUTER_DRIVER_TYPE_BY_ID, config.RetrieveDriverType)
	r.PUT(URL_ROUTER_DRIVER_TYPE, config.UpdateDriverType)
	r.DELETE(URL_ROUTER_DRIVER_TYPE_BY_ID, config.DeleteDriverType)
	r.PUT(URL_ROUTER_DRIVER_TYPE_BY_ID, config.DeleteDriverTypeWithFlag)
	r.POST(URL_ROUTER_DRIVER_TYPE_QUERY, config.QueryDriverType)
	r.GET(URL_ROUTER_DRIVER_TYPE_ALL, config.GetAllDriverTypes)
	return r
}