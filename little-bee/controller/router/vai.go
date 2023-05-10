package router

import (
	"controller/config"

	"github.com/gin-gonic/gin"
)

const URL_ROUTER_VAI = "/config/vai"
const URL_ROUTER_VAI_BY_ID = "/config/vai/:vai-id"
const URL_ROUTER_VAI_QUERY = "/config/vai/query"
const URL_ROUTER_VAI_ALL = "/config/vai/all"

func VaiRouter(r *gin.RouterGroup) *gin.RouterGroup {
	r.POST(URL_ROUTER_VAI, config.CreateVai)
	r.GET(URL_ROUTER_VAI_BY_ID, config.RetrieveVai)
	r.PUT(URL_ROUTER_VAI, config.UpdateVai)
	r.DELETE(URL_ROUTER_VAI_BY_ID, config.DeleteVai)
	r.PUT(URL_ROUTER_VAI_BY_ID, config.DeleteVaiWithFlag)
	return r
}
