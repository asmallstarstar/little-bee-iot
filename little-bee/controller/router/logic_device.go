package router

import (
	"controller/config"

	"github.com/gin-gonic/gin"
)

const URL_ROUTER_LOGIC_DEVICE = "/config/logic-device"
const URL_ROUTER_LOGIC_DEVICE_BY_ID = "/config/logic-device/:logic-device-id"
const URL_ROUTER_LOGIC_DEVICE_QUERY = "/config/logic-device/query"
const URL_ROUTER_LOGIC_DEVICE_ALL = "/config/logic-device/all"

func LogicDeviceRouter(r *gin.RouterGroup) *gin.RouterGroup {
	r.POST(URL_ROUTER_LOGIC_DEVICE, config.CreateLogicDevice)
	r.GET(URL_ROUTER_LOGIC_DEVICE_BY_ID, config.RetrieveLogicDevice)
	r.PUT(URL_ROUTER_LOGIC_DEVICE, config.UpdateLogicDevice)
	r.DELETE(URL_ROUTER_LOGIC_DEVICE_BY_ID, config.DeleteLogicDevice)
	r.POST(URL_ROUTER_LOGIC_DEVICE_QUERY, config.QueryLogicDevice)
	r.GET(URL_ROUTER_LOGIC_DEVICE_ALL, config.GetAllLogicDevices)
	return r
}