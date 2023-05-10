package router

import (
	"controller/config"

	"github.com/gin-gonic/gin"
)

const URL_ROUTER_CONTROL_COMMAND_RECORD = "/config/control-command-record"
const URL_ROUTER_CONTROL_COMMAND_RECORD_BY_ID = "/config/control-command-record/:control-command-record-id"
const URL_ROUTER_CONTROL_COMMAND_RECORD_QUERY = "/config/control-command-record/query"
const URL_ROUTER_CONTROL_COMMAND_RECORD_ALL = "/config/control-command-record/all"

func ControlCommandRecordRouter(r *gin.RouterGroup) *gin.RouterGroup {
	r.POST(URL_ROUTER_CONTROL_COMMAND_RECORD, config.CreateControlCommandRecord)
	r.GET(URL_ROUTER_CONTROL_COMMAND_RECORD_BY_ID, config.RetrieveControlCommandRecord)
	r.PUT(URL_ROUTER_CONTROL_COMMAND_RECORD, config.UpdateControlCommandRecord)
	r.POST(URL_ROUTER_CONTROL_COMMAND_RECORD_QUERY, config.QueryControlCommandRecord)
	r.GET(URL_ROUTER_CONTROL_COMMAND_RECORD_ALL, config.GetAllControlCommandRecords)
	return r
}
