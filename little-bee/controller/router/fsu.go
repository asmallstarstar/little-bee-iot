package router

import (
	"controller/config"

	"github.com/gin-gonic/gin"
)

const URL_ROUTER_FSU = "/config/fsu"
const URL_ROUTER_FSU_BY_ID = "/config/fsu/:fsu-id"
const URL_ROUTER_FSU_QUERY = "/config/fsu/query"
const URL_ROUTER_FSU_ALL = "/config/fsu/all"

func FsuRouter(r *gin.RouterGroup) *gin.RouterGroup {
	r.POST(URL_ROUTER_FSU, config.CreateFsu)
	r.GET(URL_ROUTER_FSU_BY_ID, config.RetrieveFsu)
	r.PUT(URL_ROUTER_FSU, config.UpdateFsu)
	r.DELETE(URL_ROUTER_FSU_BY_ID, config.DeleteFsu)
	r.PUT(URL_ROUTER_FSU_BY_ID, config.DeleteFsuWithFlag)
	r.POST(URL_ROUTER_FSU_QUERY, config.QueryFsu)
	r.GET(URL_ROUTER_FSU_ALL, config.GetAllFsus)
	return r
}