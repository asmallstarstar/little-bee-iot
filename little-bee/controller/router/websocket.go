package router

import (
	"controller/websocket"

	"github.com/gin-gonic/gin"
)

const URL_WEBSOCKET = "/ws"

func WebsocketRouter(r *gin.RouterGroup) *gin.RouterGroup {
	r.GET(URL_WEBSOCKET, websocket.Upgrade)
	return r
}
