package router

import (
	"controller/config"

	"github.com/gin-gonic/gin"
)

const URL_ROUTER_VDI = "/config/vdi"
const URL_ROUTER_VDI_BY_ID = "/config/vdi/:vdi-id"
const URL_ROUTER_VDI_QUERY = "/config/vdi/query"
const URL_ROUTER_VDI_ALL = "/config/vdi/all"

func VdiRouter(r *gin.RouterGroup) *gin.RouterGroup {
	r.POST(URL_ROUTER_VDI, config.CreateVdi)
	r.GET(URL_ROUTER_VDI_BY_ID, config.RetrieveVdi)
	r.PUT(URL_ROUTER_VDI, config.UpdateVdi)
	r.DELETE(URL_ROUTER_VDI_BY_ID, config.DeleteVdi)
	r.PUT(URL_ROUTER_VDI_BY_ID, config.DeleteVdiWithFlag)
	return r
}
