package fsuinstance

import (
	"encoding/json"
	"fmt"
	fs "grpc/fsugrpcserver"
	rc "grpc/realdataclient"
	pb "message/littlebee"
	"net"
	"reflect"
	"runtime"
	"strconv"
	"strings"
	"time"
	log "util/log"

	"go.uber.org/zap"
	"golang.org/x/text/encoding/simplifiedchinese"
	"google.golang.org/protobuf/types/known/timestamppb"
)

const JITONG_SIGNAL_BINDING_METADATA string = "JitongSignalBindingMetadata"

type JitongConfigureBinding struct {
	productID string
	deviceID  string
	paraID    string
}

type FsuJitong struct {
	serverIp            string
	serverPort          int
	requestPeriod       int
	quit                chan bool
	conn                net.Conn
	mapConfigureBinding map[int32]*JitongConfigureBinding
	timeoutCount        int
	lastReceiveTime     int64
	BaseFsu
}

func RegistJitongFsuReflectType(m map[string]reflect.Type) {
	m["FsuJitong"] = reflect.TypeOf(&FsuJitong{})
}

func (f *FsuJitong) Init() {
	f.mapConfigureBinding = make(map[int32]*JitongConfigureBinding)
	f.lastReceiveTime = 0
}

func (f *FsuJitong) SetFsuParameter(parameter map[string]*pb.MetadataAttribute) {
	f.FsuParameter = parameter
	for _, v := range parameter {
		switch v.MetadataAttributeName {
		case "ServerIp":
			f.serverIp = v.MetadataAttributeValue
		case "ServerPort":
			f.serverPort, _ = strconv.Atoi(v.MetadataAttributeValue)
		case "RequestPeriod":
			f.requestPeriod, _ = strconv.Atoi(v.MetadataAttributeValue)
		}
	}
}

func (f *FsuJitong) SetFsuId(fsuId int32) {
	f.FsuId = fsuId
}

func (f *FsuJitong) GetFsuId() int32 {
	return f.FsuId
}

func (f *FsuJitong) SetAgentId(agentId int32) {
	f.AgentId = agentId
}

func (f *FsuJitong) GetAgentId() int32 {
	return f.AgentId
}

func (f *FsuJitong) SetFsuTypeId(fsuTypeId int32) {
	f.FsuTypeId = fsuTypeId
}

func (f *FsuJitong) GetFsuTypeId() int32 {
	return f.FsuTypeId
}

func (f *FsuJitong) SetFSUName(fsuName string) {
	f.FsuName = fsuName
}

func (f *FsuJitong) GetFsuName() string {
	return f.FsuName
}

func (f *FsuJitong) GetFsuTypeName() string {
	return "Jitong FSU Type"
}

func (f *FsuJitong) GetFsuRunState() pb.FsuRunStateEnum {
	return f.FsuRunState
}

func (f *FsuJitong) SetSendSignalRawValueFn(realdataclient *rc.RealdataClient) {
	f.RealdataClient = realdataclient
}

func (f *FsuJitong) SetFsuToDriverFn(fsuserver *fs.FsuServer) {
	f.FsuServer = fsuserver
}

func (f *FsuJitong) GetSignalBindingMetadata() string {
	return JITONG_SIGNAL_BINDING_METADATA
}

func (f *FsuJitong) SetSignalBindingConfigure(configureId int32, configure string) {
	var m map[string]*pb.MetadataAttribute
	err := json.Unmarshal([]byte(configure), &m)
	if err != nil {
		log.Logger.Error("failed to binding signal configure", zap.Int("fsu id", int(f.FsuId)),
			zap.String("configure", configure), zap.String("error", err.Error()))
		return
	}
	b := &JitongConfigureBinding{}
	for _, val := range m {
		switch val.MetadataAttributeName {
		case "ProductID":
			b.productID = val.MetadataAttributeValue
		case "DeviceID":
			b.deviceID = val.MetadataAttributeValue
		case "ParaID":
			b.paraID = val.MetadataAttributeValue
		}
	}
	f.mapConfigureBinding[configureId] = b
}

func (f *FsuJitong) SetDriverRelated(driverId int32, configureId int32, configure string) {

}

func (f *FsuJitong) Start() {
	f.quit = make(chan bool)
	f.timeoutCount = 0
}

func (f *FsuJitong) Run() {
	t := f.requestPeriod / 2
	for {
		select {
		case <-f.quit:
			close(f.quit)
			runtime.Goexit()
		case <-time.After(time.Duration(t) * time.Millisecond):
			if f.conn != nil {
				f.sendRequest()
				time.Sleep(time.Duration(t) * time.Millisecond)
				f.receiveResponse()
				f.sendJitongFSUStatus(f.lastReceiveTime, int32(pb.FsuRunStateEnum_FSU_RUN_STATE_CONNECTED), "running")
			} else {
				f.sendJitongFSUStatus(f.lastReceiveTime, int32(pb.FsuRunStateEnum_FSU_RUN_STATE_DISCONNECTED), "disconnected")
				f.connect()
			}
		}
	}
}

func (f *FsuJitong) Control(controlCommandRequest *pb.BottomControlCommandRequest) {

}

func (f *FsuJitong) SendToBottom(d *pb.DriverToFsu) {

}

func (f *FsuJitong) Stop() {
	f.quit <- true
	if f.conn != nil {
		f.conn.Close()
		f.conn = nil
	}
}

func (f *FsuJitong) Destroy() {

}

func (f *FsuJitong) connect() {
	host := fmt.Sprintf("%s:%d", f.serverIp, f.serverPort)
	var err error
	f.conn, err = net.DialTimeout("tcp", host, time.Second*3)
	if err != nil {
		log.Logger.Error("failed to connect jitong server", zap.Int32("fsu id", f.FsuId), zap.String("jitong server", host), zap.String("error", err.Error()))
	}
}

func (f *FsuJitong) sendRequest() {
	if f.conn == nil {
		return
	}
	utf8json := `{"Action": "GetAllCurValue"}`
	gb2312Str, _ := simplifiedchinese.GBK.NewEncoder().Bytes([]byte(utf8json))
	_, e := f.conn.Write([]byte(gb2312Str))
	if e != nil {
		f.conn.Close()
		f.conn = nil
	}
}

func (f *FsuJitong) receiveResponse() {
	if f.conn == nil {
		return
	}
	var json string = ""
	for {
		b := make([]byte, 5120)
		f.conn.SetReadDeadline(time.Now().Add(1 * time.Second))
		len, e := f.conn.Read(b)
		if len > 0 {
			f.lastReceiveTime = time.Now().Unix()
			json = json + string(b)
		} else if e != nil {
			if netErr, ok := e.(net.Error); ok && netErr.Timeout() {
				f.timeoutCount++
				if f.timeoutCount > 5 {
					f.conn.Close()
					f.conn = nil
				}
				break
			} else {
				f.conn.Close()
				f.conn = nil
				return
			}
		}
	}

	if len(json) > 0 {
		f.timeoutCount = 0
		if !f.analyzingResponse(json) {
			f.conn.Close()
			f.conn = nil
		}
	}
}

func (f *FsuJitong) analyzingResponse(d string) bool {
	reader := strings.NewReader(d)
	decoder := json.NewDecoder(reader)
	data := make(map[string]interface{})
	if err := decoder.Decode(&data); err != nil {
		log.Logger.Debug("analyzing response error", zap.String("json", string(d)), zap.String("error", err.Error()))
		return false
	}
	pid, _ := (data["ProductID"]).(string)
	rtDataArray, _ := data["RtData"].([]interface{})
	for _, v := range rtDataArray {
		rtDataMap, _ := v.(map[string]interface{})
		configureId := f.getConfigureId(pid, rtDataMap["DeviceID"].(string), rtDataMap["ParaID"].(string))
		if configureId != 0 {
			sv := &pb.SignalRawValue{
				FsuId:             f.FsuId,
				ConfigureId:       configureId,
				ValueOccurredTime: timestamppb.Now(),
				SignalValueState:  pb.SignalValueStateEnum_SIGNAL_VALUE_STATE_VALID,
				SignalValueType:   pb.SignalValueTypeEnum_SIGNAL_VALUE_TYPE_STRING,
				DigitalValue:      0,
				AnalogValue:       0.0,
				StringValue:       []byte(rtDataMap["Value"].(string)),
			}
			f.RealdataClient.SendSignalRawValue(sv)
		}
	}
	return true
}

func (f *FsuJitong) getConfigureId(productID string, deviceID string, paraID string) int32 {
	for i, v := range f.mapConfigureBinding {
		if v.productID == productID && v.deviceID == deviceID && v.paraID == paraID {
			return i
		}
	}
	return 0
}

func (f *FsuJitong) sendJitongFSUStatus(lastReceiveTime int64, fsuStatus int32, statusDesc string) {
	f.RealdataClient.SendFSUStatus(&pb.FSUStatus{
		FsuId:            f.FsuId,
		FsuName:          f.FsuName,
		FsuType:          f.GetFsuTypeName(),
		SubChannelName:   "",
		LastReceiveTime:  lastReceiveTime,
		FsuStatus:        fsuStatus,
		SubChannelStatus: int32(pb.FsuRunStateEnum_FSU_RUN_STATE_UNKNOWN),
		StatusDesc:       statusDesc,
	})
}
