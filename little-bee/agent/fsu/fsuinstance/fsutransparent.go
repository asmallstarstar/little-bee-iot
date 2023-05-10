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
	"time"
	log "util/log"

	"go.uber.org/zap"
)

type FsuTransparent struct {
	serverIp           string
	basePort           int
	portCount          int
	quit               chan bool
	mapPortToConn      map[int]net.Conn //port--> conn
	mapDriverToPort    map[int32]int32  //driverId --> port
	mapTimeoutCount    map[int32]int32  //port-->timeout count
	mapLastReceiveTime map[int32]int64  //port-->last receive time

	BaseFsu
}

func RegistTransparentFsuReflectType(m map[string]reflect.Type) {
	m["FsuTransparent"] = reflect.TypeOf(&FsuTransparent{})
}

func (f *FsuTransparent) Init() {
	f.MapDriverToRelatedConfigureId = make(map[int32]int32)
	f.mapDriverToPort = make(map[int32]int32)
	f.mapPortToConn = make(map[int]net.Conn)
	f.mapTimeoutCount = make(map[int32]int32)
	f.mapLastReceiveTime = make(map[int32]int64)
}

func (f *FsuTransparent) SetFsuParameter(parameter map[string]*pb.MetadataAttribute) {
	f.FsuParameter = parameter
	for _, v := range parameter {
		switch v.MetadataAttributeName {
		case "ServerIp":
			f.serverIp = v.MetadataAttributeValue
		case "BasePort":
			f.basePort, _ = strconv.Atoi(v.MetadataAttributeValue)
		case "PortCount":
			f.portCount, _ = strconv.Atoi(v.MetadataAttributeValue)
		}
	}
}

func (f *FsuTransparent) SetFsuId(fsuId int32) {
	f.FsuId = fsuId
}

func (f *FsuTransparent) GetFsuId() int32 {
	return f.FsuId
}

func (f *FsuTransparent) SetAgentId(agentId int32) {
	f.AgentId = agentId
}

func (f *FsuTransparent) GetAgentId() int32 {
	return f.AgentId
}

func (f *FsuTransparent) SetFsuTypeId(fsuTypeId int32) {
	f.FsuTypeId = fsuTypeId
}

func (f *FsuTransparent) GetFsuTypeId() int32 {
	return f.FsuTypeId
}

func (f *FsuTransparent) SetFSUName(fsuName string) {
	f.FsuName = fsuName
}

func (f *FsuTransparent) GetFsuName() string {
	return f.FsuName
}

func (f *FsuTransparent) GetFsuTypeName() string {
	return "Transparent FSU Type"
}

func (f *FsuTransparent) GetFsuRunState() pb.FsuRunStateEnum {
	return f.FsuRunState
}

func (f *FsuTransparent) SetSendSignalRawValueFn(realdataclient *rc.RealdataClient) {
	f.RealdataClient = realdataclient
}

func (f *FsuTransparent) SetFsuToDriverFn(fsuserver *fs.FsuServer) {
	f.FsuServer = fsuserver
}

func (f *FsuTransparent) GetSignalBindingMetadata() string {
	return ""
}

func (f *FsuTransparent) SetSignalBindingConfigure(configureId int32, configure string) {

}

func (f *FsuTransparent) SetDriverRelated(driverId int32, configureId int32, configure string) {
	f.MapDriverToRelatedConfigureId[driverId] = configureId

	var m map[string]*pb.MetadataAttribute
	err := json.Unmarshal([]byte(configure), &m)
	if err != nil {
		log.Logger.Error("failed to convert configure json", zap.Int("driver id", int(driverId)),
			zap.String("configure", configure), zap.String("error", err.Error()))
		return
	}

	for _, v := range m {
		switch v.MetadataAttributeName {
		case "Port":
			p, _ := strconv.Atoi(v.MetadataAttributeValue)
			f.mapDriverToPort[driverId] = int32(p)
			f.mapPortToConn[p] = nil
			f.mapLastReceiveTime[int32(p)] = 0
			f.mapTimeoutCount[int32(p)] = 0
		}
	}
}

func (f *FsuTransparent) Start() {
	f.quit = make(chan bool)
}

func (f *FsuTransparent) Run() {
	for {
		select {
		case <-f.quit:
			close(f.quit)
			runtime.Goexit()
		case <-time.After(time.Duration(10) * time.Millisecond):
			for _, v := range f.mapDriverToPort {
				if f.mapPortToConn[int(v)] != nil {
					f.receiveResponse(f.mapPortToConn[int(v)], v)
					f.sendTransparentFSUStatus(strconv.Itoa(int(v)), f.mapLastReceiveTime[v], int32(pb.FsuRunStateEnum_FSU_RUN_STATE_CONNECTED), "running")
				} else {
					f.sendTransparentFSUStatus(strconv.Itoa(int(v)), f.mapLastReceiveTime[v], int32(pb.FsuRunStateEnum_FSU_RUN_STATE_DISCONNECTED), "disconnected")
					f.connect(f.serverIp, int(v))
				}
			}
		}
	}
}

func (f *FsuTransparent) Control(controlCommandRequest *pb.BottomControlCommandRequest) {

}

func (f *FsuTransparent) SendToBottom(d *pb.DriverToFsu) {
	var driverId int32 = 0
	for k, v := range f.MapDriverToRelatedConfigureId {
		if v == d.ConfigureId {
			driverId = k
		}
	}
	p := f.mapDriverToPort[driverId]
	f.sendRequest(d.Request, f.mapPortToConn[int(p)], p)
}

func (f *FsuTransparent) Stop() {
	f.quit <- true
	for _, v := range f.mapPortToConn {
		if v != nil {
			v.Close()
		}
	}
}

func (f *FsuTransparent) Destroy() {

}

func (f *FsuTransparent) connect(serverIp string, serverPort int) {
	host := fmt.Sprintf("%s:%d", serverIp, serverPort)
	var err error
	f.mapPortToConn[serverPort], err = net.DialTimeout("tcp", host, time.Second*1)
	if err != nil {
		log.Logger.Error("failed to connect server", zap.Int32("fsu id", f.FsuId), zap.String("server", host), zap.String("error", err.Error()))
	}
}

func (f *FsuTransparent) sendRequest(data []byte, c net.Conn, port int32) {
	if c == nil {
		return
	}
	_, e := c.Write(data)
	if e != nil {
		c.Close()
		c = nil
		f.mapPortToConn[int(port)] = nil
	}
}

func (f *FsuTransparent) receiveResponse(c net.Conn, port int32) {
	if c == nil {
		return
	}
	b := make([]byte, 1024)
	c.SetReadDeadline(time.Now().Add(1 * time.Second))
	len, e := c.Read(b)
	if e != nil {
		if netErr, ok := e.(net.Error); ok && netErr.Timeout() {
			f.mapTimeoutCount[port]++
			if f.mapTimeoutCount[port] > 15 {
				c.Close()
				f.mapPortToConn[int(port)] = nil
			}
			return
		} else {
			c.Close()
			f.mapPortToConn[int(port)] = nil
			return
		}
	}
	f.mapLastReceiveTime[port] = time.Now().Unix()
	f.mapTimeoutCount[port] = 0
	f.upload(b[0:len], port)
}

func (f *FsuTransparent) upload(d []byte, port int32) {
	var driverId int32 = 0
	for k, v := range f.mapDriverToPort {
		if port == v {
			driverId = k
		}
	}
	f.FsuServer.SendFsuToDriver(&pb.FsuToDriver{
		DriverId: driverId,
		FsuId:    f.FsuId,
		Response: d,
	})
}

func (f *FsuTransparent) sendTransparentFSUStatus(subChannelName string, lastReceiveTime int64, fsuStatus int32, statusDesc string) {
	f.RealdataClient.SendFSUStatus(&pb.FSUStatus{
		FsuId:            f.FsuId,
		FsuName:          f.FsuName,
		FsuType:          f.GetFsuTypeName(),
		SubChannelName:   subChannelName,
		LastReceiveTime:  lastReceiveTime,
		FsuStatus:        fsuStatus,
		SubChannelStatus: int32(pb.FsuRunStateEnum_FSU_RUN_STATE_UNKNOWN),
		StatusDesc:       statusDesc,
	})
}
