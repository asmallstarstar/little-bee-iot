package router

import (
	"controller/config"

	"github.com/gin-gonic/gin"
)

const URL_ROUTER_ALARM_LEVEL = "/config/alarm-level"
const URL_ROUTER_ALARM_LEVEL_BY_ID = "/config/alarm-level/:alarm-level-id"
const URL_ROUTER_ALARM_LEVEL_QUERY = "/config/alarm-level/query"
const URL_ROUTER_ALARM_LEVEL_ALL = "/config/alarm-level/all"

func AlarmLevelRouter(r *gin.RouterGroup) *gin.RouterGroup {
	r.POST(URL_ROUTER_ALARM_LEVEL, config.CreateAlarmLevel)
	r.GET(URL_ROUTER_ALARM_LEVEL_BY_ID, config.RetrieveAlarmLevel)
	r.PUT(URL_ROUTER_ALARM_LEVEL, config.UpdateAlarmLevel)
	r.DELETE(URL_ROUTER_ALARM_LEVEL_BY_ID, config.DeleteAlarmLevel)
	r.PUT(URL_ROUTER_ALARM_LEVEL_BY_ID, config.DeleteAlarmLevelWithFlag)
	r.POST(URL_ROUTER_ALARM_LEVEL_QUERY, config.QueryAlarmLevel)
	r.GET(URL_ROUTER_ALARM_LEVEL_ALL, config.GetAllAlarmLevels)
	return r
}