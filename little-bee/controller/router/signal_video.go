package router

import (
	"controller/config"

	"github.com/gin-gonic/gin"
)

const URL_ROUTER_SIGNAL_VIDEO = "/config/signal-video"
const URL_ROUTER_SIGNAL_VIDEO_BY_ID = "/config/signal-video/:signal-video-id"
const URL_ROUTER_SIGNAL_VIDEO_QUERY = "/config/signal-video/query"
const URL_ROUTER_SIGNAL_VIDEO_ALL = "/config/signal-video/all"

func SignalVideoRouter(r *gin.RouterGroup) *gin.RouterGroup {
	r.POST(URL_ROUTER_SIGNAL_VIDEO, config.CreateSignalVideo)
	r.GET(URL_ROUTER_SIGNAL_VIDEO_BY_ID, config.RetrieveSignalVideo)
	r.PUT(URL_ROUTER_SIGNAL_VIDEO, config.UpdateSignalVideo)
	r.DELETE(URL_ROUTER_SIGNAL_VIDEO_BY_ID, config.DeleteSignalVideo)
	r.POST(URL_ROUTER_SIGNAL_VIDEO_QUERY, config.QuerySignalVideo)
	r.GET(URL_ROUTER_SIGNAL_VIDEO_ALL, config.GetAllSignalVideos)
	return r
}
