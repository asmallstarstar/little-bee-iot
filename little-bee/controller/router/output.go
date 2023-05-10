package router

import (
	"controller/config"

	"github.com/gin-gonic/gin"
)

const URL_ROUTER_OUTPUT = "/config/output"
const URL_ROUTER_OUTPUT_BY_ID = "/config/output/:output-id"
const URL_ROUTER_OUTPUT_QUERY = "/config/output/query"
const URL_ROUTER_OUTPUT_ALL = "/config/output/all"

func OutputRouter(r *gin.RouterGroup) *gin.RouterGroup {
	r.POST(URL_ROUTER_OUTPUT, config.CreateOutput)
	r.GET(URL_ROUTER_OUTPUT_BY_ID, config.RetrieveOutput)
	r.PUT(URL_ROUTER_OUTPUT, config.UpdateOutput)
	r.DELETE(URL_ROUTER_OUTPUT_BY_ID, config.DeleteOutput)
	r.PUT(URL_ROUTER_OUTPUT_BY_ID, config.DeleteOutputWithFlag)
	return r
}
