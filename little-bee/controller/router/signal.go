package router

import (
	"controller/config"

	"github.com/gin-gonic/gin"
)

const URL_ROUTER_SIGNAL = "/config/signal"
const URL_ROUTER_SIGNAL_BY_ID = "/config/signal/:signal-id"
const URL_ROUTER_SIGNAL_QUERY = "/config/signal/query"
const URL_ROUTER_SIGNAL_ALL = "/config/signal/all"

func SignalRouter(r *gin.RouterGroup) *gin.RouterGroup {
	r.POST(URL_ROUTER_SIGNAL, config.CreateSignal)
	r.GET(URL_ROUTER_SIGNAL_BY_ID, config.RetrieveSignal)
	r.PUT(URL_ROUTER_SIGNAL, config.UpdateSignal)
	r.DELETE(URL_ROUTER_SIGNAL_BY_ID, config.DeleteSignal)
	r.PUT(URL_ROUTER_SIGNAL_BY_ID, config.DeleteSignalWithFlag)
	r.POST(URL_ROUTER_SIGNAL_QUERY, config.QuerySignal)
	r.GET(URL_ROUTER_SIGNAL_ALL, config.GetAllSignals)
	return r
}