package router

import (
	"controller/config"

	"github.com/gin-gonic/gin"
)

const URL_ROUTER_SIGNAL_VIRTUAL = "/config/signal-virtual"
const URL_ROUTER_SIGNAL_VIRTUAL_BY_ID = "/config/signal-virtual/:signal-virtual-id"
const URL_ROUTER_SIGNAL_VIRTUAL_QUERY = "/config/signal-virtual/query"
const URL_ROUTER_SIGNAL_VIRTUAL_ALL = "/config/signal-virtual/all"

func SignalVirtualRouter(r *gin.RouterGroup) *gin.RouterGroup {
	r.POST(URL_ROUTER_SIGNAL_VIRTUAL, config.CreateSignalVirtual)
	r.GET(URL_ROUTER_SIGNAL_VIRTUAL_BY_ID, config.RetrieveSignalVirtual)
	r.PUT(URL_ROUTER_SIGNAL_VIRTUAL, config.UpdateSignalVirtual)
	r.DELETE(URL_ROUTER_SIGNAL_VIRTUAL_BY_ID, config.DeleteSignalVirtual)
	r.POST(URL_ROUTER_SIGNAL_VIRTUAL_QUERY, config.QuerySignalVirtual)
	r.GET(URL_ROUTER_SIGNAL_VIRTUAL_ALL, config.GetAllSignalVirtuals)
	return r
}