package router

import (
	"controller/config"

	"github.com/gin-gonic/gin"
)

const URL_ROUTER_SIGNAL_SI = "/config/signal-si"
const URL_ROUTER_SIGNAL_SI_BY_ID = "/config/signal-si/:signal-si-id"
const URL_ROUTER_SIGNAL_SI_QUERY = "/config/signal-si/query"
const URL_ROUTER_SIGNAL_SI_ALL = "/config/signal-si/all"

func SignalSiRouter(r *gin.RouterGroup) *gin.RouterGroup {
	r.POST(URL_ROUTER_SIGNAL_SI, config.CreateSignalSi)
	r.GET(URL_ROUTER_SIGNAL_SI_BY_ID, config.RetrieveSignalSi)
	r.PUT(URL_ROUTER_SIGNAL_SI, config.UpdateSignalSi)
	r.DELETE(URL_ROUTER_SIGNAL_SI_BY_ID, config.DeleteSignalSi)
	r.POST(URL_ROUTER_SIGNAL_SI_QUERY, config.QuerySignalSi)
	r.GET(URL_ROUTER_SIGNAL_SI_ALL, config.GetAllSignalSis)
	return r
}