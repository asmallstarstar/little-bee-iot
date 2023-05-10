package router

import (
	"controller/config"

	"github.com/gin-gonic/gin"
)

const URL_ROUTER_VIDEO = "/config/video"
const URL_ROUTER_VIDEO_BY_ID = "/config/video/:video-id"
const URL_ROUTER_VIDEO_QUERY = "/config/video/query"
const URL_ROUTER_VIDEO_ALL = "/config/video/all"

func VideoRouter(r *gin.RouterGroup) *gin.RouterGroup {
	r.POST(URL_ROUTER_VIDEO, config.CreateVideo)
	r.GET(URL_ROUTER_VIDEO_BY_ID, config.RetrieveVideo)
	r.PUT(URL_ROUTER_VIDEO, config.UpdateVideo)
	r.DELETE(URL_ROUTER_VIDEO_BY_ID, config.DeleteVideo)
	r.PUT(URL_ROUTER_VIDEO_BY_ID, config.DeleteVideoWithFlag)
	return r
}
