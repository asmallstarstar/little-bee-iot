package maininit

import (
	"controller/middleware"
	"controller/router"
	"controller/websocket"
	dao "dao/database"
	c "service/consumer"
	"util/log"
	yaml "util/yaml/service"

	"util/i18n"

	"github.com/gin-gonic/gin"
)

func GinEngine() *gin.Engine {
	r := gin.New()

	r.Use(middleware.GinLogger(), middleware.GinRecovery(true))

	pub := r.Group(yaml.Yaml.Service.UrlRoot + "/public")
	PublicRouters(pub)

	pri := r.Group(yaml.Yaml.Service.UrlRoot)
	pri.Use(middleware.Auth(), middleware.Grant())
	PrivateRouters(pri)
	return r
}

func PrivateRouters(g *gin.RouterGroup) {
	router.ActionRouter(g)
	router.AgentRouter(g)
	router.AlarmLevelRouter(g)
	router.AlarmRouter(g)
	router.AreaTypeRouter(g)
	router.ConfigureRouter(g)
	router.ControlCommandRecordRouter(g)
	router.DepartmentRouter(g)
	router.DeviceTypeRouter(g)
	router.DriverTypeRouter(g)
	router.DriverRouter(g)
	router.FsuTypeRouter(g)
	router.FsuRouter(g)
	router.LogicAreaRouter(g)
	router.LogicDeviceRouter(g)
	router.LogicObjectRouter(g)
	router.MenuRouter(g)
	router.MetadataRouter(g)
	router.RoleRouter(g)
	router.SignalAcquisitionRouter(g)
	router.SignalAiRouter(g)
	router.SignalDiRouter(g)
	router.SignalSiRouter(g)
	router.SignalThreshholdRouter(g)
	router.SignalUnificationRouter(g)
	router.SignalVirtualRouter(g)
	router.SignalRouter(g)
	router.UserPrivateRouter(g)
	router.MenuActionRouter(g)
	router.ActionGroupRouter(g)
	router.SignalVideoRouter(g)
	router.AiRouter(g)
	router.DiRouter(g)
	router.SiRouter(g)
	router.VaiRouter(g)
	router.VdiRouter(g)
	router.VsiRouter(g)
	router.OutputRouter(g)
	router.VideoRouter(g)

	router.RealdataRouter(g)
	router.FsuStatusRouter(g)
}

func PublicRouters(g *gin.RouterGroup) {
	router.UserPublicRouter(g)
	router.WebsocketRouter(g)
}

func Initialization() {
	yaml.Init()
	i18n.InitI18n()
	log.InitLog(
		yaml.Yaml.Log.FileName,
		yaml.Yaml.Log.MaxAge,
		yaml.Yaml.Log.MaxBackups,
		yaml.Yaml.Log.MaxAge,
		yaml.Yaml.Log.Level)
	dao.InitDB(yaml.Yaml.Database.Dsn, yaml.Yaml.Database.LogLevel)
}

func SetWebRunMode(mode string) {
	switch mode {
	case "Release":
		gin.SetMode(gin.ReleaseMode)
	case "Debug":
		gin.SetMode(gin.DebugMode)
	case "Test":
		gin.SetMode(gin.TestMode)
	}
}

func MQTTInit() bool {
	if !c.Connect() {
		log.Logger.Error("failed to connect mqtt")
		return false
	}
	if e := c.SubRealdata(); e != nil {
		log.Logger.Error("failed to sub realdata")
		return false
	}
	if e := c.SubFSUStatus(); e != nil {
		log.Logger.Error("failed to sub fsu status")
		return false
	}
	if e := c.SubBeginAlarm(websocket.HandleAlarmBegin); e != nil {
		log.Logger.Error("failed to sub beginning alarm message")
		return false
	}
	if e := c.SubEndAlarm(websocket.HandleAlarmEnd); e != nil {
		log.Logger.Error("failed to sub ended alarm message")
		return false
	}
	if e := c.SubControlReport(websocket.HandleControlReport); e != nil {
		log.Logger.Error("failed to sub controlled command report")
		return false
	}
	if e := c.SubAckAlarm(websocket.HandleAckAlarm); e != nil {
		log.Logger.Error("failed to sub ack alarm message")
		return false
	}
	return true
}
