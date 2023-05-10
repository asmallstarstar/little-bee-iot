package mqtt

import (
	"time"
	"util/log"

	MQTT "github.com/eclipse/paho.mqtt.golang"
	"go.uber.org/zap"
)

const (
	TOPIC_TYPE_REAL_DATA      string = "/real-data"
	TOPIC_TYPE_BEGIN_ALARM    string = "/begin-alarm"
	TOPIC_TYPE_END_ALARM      string = "/end-alarm"
	TOPIC_TYPE_CONTROL        string = "/control"
	TOPIC_TYPE_CONTROL_REPORT string = "/control-report"
	TOPIC_TYPE_FSUSTATUS      string = "/fsu-status"
	TOPIC_TYPE_ACKALARM       string = "/ack-alarm"
)

const (
	MQTT_QOS_0 byte = 0
	MQTT_QOS_1 byte = 1
	MQTT_QOS_2 byte = 2
)

var c MQTT.Client = nil

func Connect(broker string, userName string, password string, clientId string) bool {
	opts := MQTT.NewClientOptions()
	opts.AddBroker(broker)
	opts.SetClientID(clientId)
	opts.SetUsername(userName)
	opts.SetPassword(password)
	opts.SetAutoReconnect(true)
	opts.SetConnectRetry(true)
	opts.SetConnectRetryInterval(3 * time.Second)
	opts.SetKeepAlive(2 * time.Second)
	opts.SetPingTimeout(1 * time.Second)
	opts.SetProtocolVersion(4)
	opts.SetAutoAckDisabled(false)
	opts.SetConnectTimeout(10 * time.Second)

	c = MQTT.NewClient(opts)
	log.Logger.Info("connecting mqtt broker")
	if token := c.Connect(); token.Wait() && token.Error() != nil {
		log.Logger.Error("failed to connect mqtt broker", zap.String("error", token.Error().Error()))
		return false
	}
	log.Logger.Info("successed to connect mqtt broker")
	return true
}

func Subscribe(topic string, qos byte, callback MQTT.MessageHandler) error {
	if token := c.Subscribe(topic, qos, callback); token.Wait() && token.Error() != nil {
		return token.Error()
	}
	return nil
}

func Publish(topic string, qos byte, retained bool, payload interface{}) error {
	if token := c.Publish(topic, qos, retained, payload); token.Wait() && token.Error() != nil {
		return token.Error()
	}
	return nil
}

func PublishRealdata(payload interface{}) {
	c.Publish(TOPIC_TYPE_REAL_DATA, MQTT_QOS_0, false, payload)
}

func PublishFsuStatus(payload interface{}) {
	c.Publish(TOPIC_TYPE_FSUSTATUS, MQTT_QOS_0, false, payload)
}

func Disconnect() {
	c.Disconnect(5000)
}
