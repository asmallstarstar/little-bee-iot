package consumer

import (
	"message/littlebee"
	m "mqtt"
	"util/log"
	yaml "util/yaml/service"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	"go.uber.org/zap"
	"google.golang.org/protobuf/proto"
)

type handleAlarmBegin func(d *littlebee.BeginAlarm)
type handleAlarmEnd func(d *littlebee.EndAlarm)
type handleControlReport func(d *littlebee.ControlCommandReport)
type handleAckAlarm func(d *littlebee.Alarm)

var MapRealdata map[int32]*littlebee.SignalRealValue = make(map[int32]*littlebee.SignalRealValue)
var MapFsuStatus map[int32]*littlebee.FSUStatus = make(map[int32]*littlebee.FSUStatus)

func Connect() bool {
	return m.Connect(yaml.Yaml.MQTT.Broker, yaml.Yaml.MQTT.UserName, yaml.Yaml.MQTT.Password, yaml.Yaml.MQTT.ClientId)
}

func PublishControl(payload interface{}) error {
	return m.Publish(m.TOPIC_TYPE_CONTROL, m.MQTT_QOS_2, false, payload)
}

func PublishAckAlarm(payload interface{}) error {
	return m.Publish(m.TOPIC_TYPE_ACKALARM, m.MQTT_QOS_2, false, payload)
}

func SubRealdata() error {
	return m.Subscribe(m.TOPIC_TYPE_REAL_DATA, m.MQTT_QOS_0, func(c mqtt.Client, m mqtt.Message) {
		v := &littlebee.SignalRealValue{}
		e := proto.Unmarshal(m.Payload(), v)
		if e == nil {
			MapRealdata[v.SignalId] = v
		} else {
			log.Logger.Error("unmarshal error", zap.String("error", e.Error()))
		}
	})
}

func SubFSUStatus() error {
	return m.Subscribe(m.TOPIC_TYPE_FSUSTATUS, m.MQTT_QOS_0, func(c mqtt.Client, m mqtt.Message) {
		s := &littlebee.FSUStatus{}
		e := proto.Unmarshal(m.Payload(), s)
		if e == nil {
			MapFsuStatus[s.FsuId] = s
		}
	})
}

func SubBeginAlarm(h handleAlarmBegin) error {
	return m.Subscribe(m.TOPIC_TYPE_BEGIN_ALARM, m.MQTT_QOS_2, func(c mqtt.Client, m mqtt.Message) {
		s := &littlebee.BeginAlarm{}
		e := proto.Unmarshal(m.Payload(), s)
		if e == nil {
			h(s)
		}
	})
}

func SubEndAlarm(h handleAlarmEnd) error {
	return m.Subscribe(m.TOPIC_TYPE_END_ALARM, m.MQTT_QOS_2, func(c mqtt.Client, m mqtt.Message) {
		s := &littlebee.EndAlarm{}
		e := proto.Unmarshal(m.Payload(), s)
		if e == nil {
			h(s)
		}
	})
}

func SubControlReport(h handleControlReport) error {
	return m.Subscribe(m.TOPIC_TYPE_CONTROL_REPORT, m.MQTT_QOS_2, func(c mqtt.Client, m mqtt.Message) {
		s := &littlebee.ControlCommandReport{}
		e := proto.Unmarshal(m.Payload(), s)
		if e == nil {
			h(s)
		}
	})
}

func SubAckAlarm(h handleAckAlarm) error {
	return m.Subscribe(m.TOPIC_TYPE_ACKALARM, m.MQTT_QOS_2, func(c mqtt.Client, m mqtt.Message) {
		s := &littlebee.Alarm{}
		e := proto.Unmarshal(m.Payload(), s)
		if e == nil {
			h(s)
		}
	})
}

func Disconnect() {
	m.Disconnect()
}
