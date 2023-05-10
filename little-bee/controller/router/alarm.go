package router

import (
	"controller/config"

	"github.com/gin-gonic/gin"
)

const URL_ROUTER_ALARM = "/config/alarm"
const URL_ROUTER_ALARM_BY_ID = "/config/alarm/:alarm-id"
const URL_ROUTER_ALARM_QUERY = "/config/alarm/query"
const URL_ROUTER_ALARM_ALL = "/config/alarm/all"
const URL_ROUTER_ACK_ALARM = "/config/alarm/ack"

func AlarmRouter(r *gin.RouterGroup) *gin.RouterGroup {
	r.POST(URL_ROUTER_ALARM, config.CreateAlarm)
	r.GET(URL_ROUTER_ALARM_BY_ID, config.RetrieveAlarm)
	r.PUT(URL_ROUTER_ALARM, config.UpdateAlarm)
	r.POST(URL_ROUTER_ALARM_QUERY, config.QueryAlarm)
	r.POST(URL_ROUTER_ACK_ALARM, config.AckAlarm)
	return r
}
