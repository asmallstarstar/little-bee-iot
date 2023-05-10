package router

import (
	"controller/config"

	"github.com/gin-gonic/gin"
)

const URL_ROUTER_DRIVER = "/config/driver"
const URL_ROUTER_DRIVER_BY_ID = "/config/driver/:driver-id"
const URL_ROUTER_DRIVER_QUERY = "/config/driver/query"
const URL_ROUTER_DRIVER_ALL = "/config/driver/all"

func DriverRouter(r *gin.RouterGroup) *gin.RouterGroup {
	r.POST(URL_ROUTER_DRIVER, config.CreateDriver)
	r.GET(URL_ROUTER_DRIVER_BY_ID, config.RetrieveDriver)
	r.PUT(URL_ROUTER_DRIVER, config.UpdateDriver)
	r.DELETE(URL_ROUTER_DRIVER_BY_ID, config.DeleteDriver)
	r.PUT(URL_ROUTER_DRIVER_BY_ID, config.DeleteDriverWithFlag)
	r.POST(URL_ROUTER_DRIVER_QUERY, config.QueryDriver)
	r.GET(URL_ROUTER_DRIVER_ALL, config.GetAllDrivers)
	return r
}