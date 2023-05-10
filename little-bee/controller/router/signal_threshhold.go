package router

import (
	"controller/config"

	"github.com/gin-gonic/gin"
)

const URL_ROUTER_SIGNAL_THRESHHOLD = "/config/signal-threshhold"
const URL_ROUTER_SIGNAL_THRESHHOLD_BY_ID = "/config/signal-threshhold/:signal-threshhold-id"
const URL_ROUTER_SIGNAL_THRESHHOLD_QUERY = "/config/signal-threshhold/query"
const URL_ROUTER_SIGNAL_THRESHHOLD_ALL = "/config/signal-threshhold/all"

func SignalThreshholdRouter(r *gin.RouterGroup) *gin.RouterGroup {
	r.POST(URL_ROUTER_SIGNAL_THRESHHOLD, config.CreateSignalThreshhold)
	r.GET(URL_ROUTER_SIGNAL_THRESHHOLD_BY_ID, config.RetrieveSignalThreshhold)
	r.PUT(URL_ROUTER_SIGNAL_THRESHHOLD, config.UpdateSignalThreshhold)
	r.DELETE(URL_ROUTER_SIGNAL_THRESHHOLD_BY_ID, config.DeleteSignalThreshhold)
	r.POST(URL_ROUTER_SIGNAL_THRESHHOLD_QUERY, config.QuerySignalThreshhold)
	r.GET(URL_ROUTER_SIGNAL_THRESHHOLD_ALL, config.GetAllSignalThreshholds)
	return r
}