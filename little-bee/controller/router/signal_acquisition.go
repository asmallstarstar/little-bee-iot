package router

import (
	"controller/config"

	"github.com/gin-gonic/gin"
)

const URL_ROUTER_SIGNAL_ACQUISITION = "/config/signal-acquisition"
const URL_ROUTER_SIGNAL_ACQUISITION_BY_ID = "/config/signal-acquisition/:signal-acquisition-id"
const URL_ROUTER_SIGNAL_ACQUISITION_QUERY = "/config/signal-acquisition/query"
const URL_ROUTER_SIGNAL_ACQUISITION_ALL = "/config/signal-acquisition/all"

func SignalAcquisitionRouter(r *gin.RouterGroup) *gin.RouterGroup {
	r.POST(URL_ROUTER_SIGNAL_ACQUISITION, config.CreateSignalAcquisition)
	r.GET(URL_ROUTER_SIGNAL_ACQUISITION_BY_ID, config.RetrieveSignalAcquisition)
	r.PUT(URL_ROUTER_SIGNAL_ACQUISITION, config.UpdateSignalAcquisition)
	r.DELETE(URL_ROUTER_SIGNAL_ACQUISITION_BY_ID, config.DeleteSignalAcquisition)
	r.POST(URL_ROUTER_SIGNAL_ACQUISITION_QUERY, config.QuerySignalAcquisition)
	r.GET(URL_ROUTER_SIGNAL_ACQUISITION_ALL, config.GetAllSignalAcquisitions)
	return r
}