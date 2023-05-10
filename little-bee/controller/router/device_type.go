package router

import (
	"controller/config"

	"github.com/gin-gonic/gin"
)

const URL_ROUTER_DEVICE_TYPE = "/config/device-type"
const URL_ROUTER_DEVICE_TYPE_BY_ID = "/config/device-type/:device-type-id"
const URL_ROUTER_DEVICE_TYPE_QUERY = "/config/device-type/query"
const URL_ROUTER_DEVICE_TYPE_ALL = "/config/device-type/all"

func DeviceTypeRouter(r *gin.RouterGroup) *gin.RouterGroup {
	r.POST(URL_ROUTER_DEVICE_TYPE, config.CreateDeviceType)
	r.GET(URL_ROUTER_DEVICE_TYPE_BY_ID, config.RetrieveDeviceType)
	r.PUT(URL_ROUTER_DEVICE_TYPE, config.UpdateDeviceType)
	r.DELETE(URL_ROUTER_DEVICE_TYPE_BY_ID, config.DeleteDeviceType)
	r.PUT(URL_ROUTER_DEVICE_TYPE_BY_ID, config.DeleteDeviceTypeWithFlag)
	r.POST(URL_ROUTER_DEVICE_TYPE_QUERY, config.QueryDeviceType)
	r.GET(URL_ROUTER_DEVICE_TYPE_ALL, config.GetAllDeviceTypes)
	return r
}