package router

import (
	"controller/config"

	"github.com/gin-gonic/gin"
)

const URL_ROUTER_SIGNAL_DI = "/config/signal-di"
const URL_ROUTER_SIGNAL_DI_BY_ID = "/config/signal-di/:signal-di-id"
const URL_ROUTER_SIGNAL_DI_QUERY = "/config/signal-di/query"
const URL_ROUTER_SIGNAL_DI_ALL = "/config/signal-di/all"

func SignalDiRouter(r *gin.RouterGroup) *gin.RouterGroup {
	r.POST(URL_ROUTER_SIGNAL_DI, config.CreateSignalDi)
	r.GET(URL_ROUTER_SIGNAL_DI_BY_ID, config.RetrieveSignalDi)
	r.PUT(URL_ROUTER_SIGNAL_DI, config.UpdateSignalDi)
	r.DELETE(URL_ROUTER_SIGNAL_DI_BY_ID, config.DeleteSignalDi)
	r.POST(URL_ROUTER_SIGNAL_DI_QUERY, config.QuerySignalDi)
	r.GET(URL_ROUTER_SIGNAL_DI_ALL, config.GetAllSignalDis)
	return r
}