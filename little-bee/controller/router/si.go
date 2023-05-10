package router

import (
	"controller/config"

	"github.com/gin-gonic/gin"
)

const URL_ROUTER_SI = "/config/si"
const URL_ROUTER_SI_BY_ID = "/config/si/:si-id"
const URL_ROUTER_SI_QUERY = "/config/si/query"
const URL_ROUTER_SI_ALL = "/config/si/all"

func SiRouter(r *gin.RouterGroup) *gin.RouterGroup {
	r.POST(URL_ROUTER_SI, config.CreateSi)
	r.GET(URL_ROUTER_SI_BY_ID, config.RetrieveSi)
	r.PUT(URL_ROUTER_SI, config.UpdateSi)
	r.DELETE(URL_ROUTER_SI_BY_ID, config.DeleteSi)
	r.PUT(URL_ROUTER_SI_BY_ID, config.DeleteSiWithFlag)
	return r
}
