package router

import (
	"controller/config"

	"github.com/gin-gonic/gin"
)

const URL_ROUTER_VSI = "/config/vsi"
const URL_ROUTER_VSI_BY_ID = "/config/vsi/:vsi-id"
const URL_ROUTER_VSI_QUERY = "/config/vsi/query"
const URL_ROUTER_VSI_ALL = "/config/vsi/all"

func VsiRouter(r *gin.RouterGroup) *gin.RouterGroup {
	r.POST(URL_ROUTER_VSI, config.CreateVsi)
	r.GET(URL_ROUTER_VSI_BY_ID, config.RetrieveVsi)
	r.PUT(URL_ROUTER_VSI, config.UpdateVsi)
	r.DELETE(URL_ROUTER_VSI_BY_ID, config.DeleteVsi)
	r.PUT(URL_ROUTER_VSI_BY_ID, config.DeleteVsiWithFlag)
	return r
}
