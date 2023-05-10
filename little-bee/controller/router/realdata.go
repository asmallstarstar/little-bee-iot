package router

import (
	"controller/realdata"

	"github.com/gin-gonic/gin"
)

const URL_ROUTER_REALDATA = "/realdata"
const URL_ROUTER_CONTROL = "/control"

func RealdataRouter(r *gin.RouterGroup) *gin.RouterGroup {
	r.POST(URL_ROUTER_REALDATA, realdata.RequestRealdata)
	r.POST(URL_ROUTER_CONTROL, realdata.ControlCommand)
	return r
}
