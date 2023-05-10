package websocket

import (
	"message/littlebee"
	"message/serialization"
	"net/http"
	"time"
	"util/log"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"go.uber.org/zap"
)

var clients = make(map[*websocket.Conn]string, 5)

const (
	MESSAGE_TYPE_UNKNOWN        int = 0
	MESSAGE_TYPE_REAL_DATA      int = 1
	MESSAGE_TYPE_BEGIN_ALARM    int = 2
	MESSAGE_TYPE_END_ALARM      int = 3
	MESSAGE_TYPE_CONTROL        int = 4
	MESSAGE_TYPE_CONTROL_REPORT int = 5
	MESSAGE_TYPE_FSUSTATUS      int = 6
)

func HandleAlarmBegin(d *littlebee.BeginAlarm) {
	jsonBytes, _ := serialization.MarshalToJson(&littlebee.WebsocketMessage{
		MessageType: "BEGIN_ALARM",
		Message:     serialization.StructToAny(d),
	})

	write(websocket.TextMessage, []byte(jsonBytes))
}

func HandleAlarmEnd(d *littlebee.EndAlarm) {
	jsonBytes, _ := serialization.MarshalToJson(&littlebee.WebsocketMessage{
		MessageType: "END_ALARM",
		Message:     serialization.StructToAny(d),
	})

	write(websocket.TextMessage, []byte(jsonBytes))
}

func HandleControlReport(d *littlebee.ControlCommandReport) {
	jsonBytes, _ := serialization.MarshalToJson(&littlebee.WebsocketMessage{
		MessageType: "CONTROL_REPORT",
		Message:     serialization.StructToAny(d),
	})

	write(websocket.TextMessage, []byte(jsonBytes))
}

func HandleAckAlarm(d *littlebee.Alarm) {
	jsonBytes, _ := serialization.MarshalToJson(&littlebee.WebsocketMessage{
		MessageType: "ACK_ALARM",
		Message:     serialization.StructToAny(d),
	})

	write(websocket.TextMessage, []byte(jsonBytes))
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:   1024,
	WriteBufferSize:  1024,
	HandshakeTimeout: 5 * time.Second,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func Upgrade(c *gin.Context) {
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Logger.Error("failed to upgrade http", zap.String("error", err.Error()))
		return
	}
	clients[conn] = ""
}

func write(messageType int, data []byte) {
	for clientConn := range clients {
		err := clientConn.WriteMessage(messageType, data)
		if err != nil {
			log.Logger.Error("failed to write message", zap.String("error", err.Error()))
			clientConn.Close()
			delete(clients, clientConn)
		}
	}
}

func Close() {
	for clientConn := range clients {
		clientConn.Close()
	}
}
