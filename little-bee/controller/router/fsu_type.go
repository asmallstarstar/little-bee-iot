package router

import (
	"controller/config"

	"github.com/gin-gonic/gin"
)

const URL_ROUTER_FSU_TYPE = "/config/fsu-type"
const URL_ROUTER_FSU_TYPE_BY_ID = "/config/fsu-type/:fsu-type-id"
const URL_ROUTER_FSU_TYPE_QUERY = "/config/fsu-type/query"
const URL_ROUTER_FSU_TYPE_ALL = "/config/fsu-type/all"

func FsuTypeRouter(r *gin.RouterGroup) *gin.RouterGroup {
	r.POST(URL_ROUTER_FSU_TYPE, config.CreateFsuType)
	r.GET(URL_ROUTER_FSU_TYPE_BY_ID, config.RetrieveFsuType)
	r.PUT(URL_ROUTER_FSU_TYPE, config.UpdateFsuType)
	r.DELETE(URL_ROUTER_FSU_TYPE_BY_ID, config.DeleteFsuType)
	r.PUT(URL_ROUTER_FSU_TYPE_BY_ID, config.DeleteFsuTypeWithFlag)
	r.POST(URL_ROUTER_FSU_TYPE_QUERY, config.QueryFsuType)
	r.GET(URL_ROUTER_FSU_TYPE_ALL, config.GetAllFsuTypes)
	return r
}