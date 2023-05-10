package router

import (
	"controller/fsustatus"

	"github.com/gin-gonic/gin"
)

const URL_ROUTER_FSUSTATUS = "/fsustatus"

func FsuStatusRouter(r *gin.RouterGroup) *gin.RouterGroup {
	r.POST(URL_ROUTER_FSUSTATUS, fsustatus.RequestFsuStatus)
	return r
}
