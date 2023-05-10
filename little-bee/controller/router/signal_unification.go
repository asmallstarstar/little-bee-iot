package router

import (
	"controller/config"

	"github.com/gin-gonic/gin"
)

const URL_ROUTER_SIGNAL_UNIFICATION = "/config/signal-unification"
const URL_ROUTER_SIGNAL_UNIFICATION_BY_ID = "/config/signal-unification/:signal-unification-id"
const URL_ROUTER_SIGNAL_UNIFICATION_QUERY = "/config/signal-unification/query"
const URL_ROUTER_SIGNAL_UNIFICATION_ALL = "/config/signal-unification/all"

func SignalUnificationRouter(r *gin.RouterGroup) *gin.RouterGroup {
	r.POST(URL_ROUTER_SIGNAL_UNIFICATION, config.CreateSignalUnification)
	r.GET(URL_ROUTER_SIGNAL_UNIFICATION_BY_ID, config.RetrieveSignalUnification)
	r.PUT(URL_ROUTER_SIGNAL_UNIFICATION, config.UpdateSignalUnification)
	r.DELETE(URL_ROUTER_SIGNAL_UNIFICATION_BY_ID, config.DeleteSignalUnification)
	r.PUT(URL_ROUTER_SIGNAL_UNIFICATION_BY_ID, config.DeleteSignalUnificationWithFlag)
	r.POST(URL_ROUTER_SIGNAL_UNIFICATION_QUERY, config.QuerySignalUnification)
	r.GET(URL_ROUTER_SIGNAL_UNIFICATION_ALL, config.GetAllSignalUnifications)
	return r
}