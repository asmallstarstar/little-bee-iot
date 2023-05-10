package fsu

import (
	"encoding/json"
	fi "fsuinstance"
	fs "grpc/fsugrpcserver"
	rc "grpc/realdataclient"
	pb "message/littlebee"
	"os"
	"reflect"
	s "service/config"
	"sync"
	"util/log"
	yaml "util/yaml/agent"

	"go.uber.org/zap"
)

type Fsu struct {
	RealdataClient       *rc.RealdataClient
	FsuServer            *fs.FsuServer
	MapNewFsuReflectType map[string]reflect.Type

	MapFsuInstances map[int32]fi.FsuInterface
	FsuInstanceLock *sync.Mutex

	Fsus       []*pb.Fsu
	FsuTypes   []*pb.FsuType
	Configures []*pb.Configure
}

var FsuManager *Fsu = &Fsu{}

func handleBottomControlRequest(d *pb.BottomControlCommandRequest) {
	FsuManager.FsuInstanceLock.Lock()
	fsu := FsuManager.MapFsuInstances[d.ControlParam.FsuId]
	if fsu != nil {
		fsu.Control(d)
	}
	FsuManager.FsuInstanceLock.Unlock()
}

func handleDriverToFsu(d *pb.DriverToFsu) {
	FsuManager.FsuInstanceLock.Lock()
	fsu := FsuManager.MapFsuInstances[d.FsuId]
	if fsu != nil {
		fsu.SendToBottom(d)
	}
	FsuManager.FsuInstanceLock.Unlock()
}

func (f *Fsu) InitGrpc() {
	c := &rc.RealdataClient{
		Conn:                              nil,
		Client:                            nil,
		BottomRawStream:                   nil,
		CenterChannelStream:               nil,
		BottomCommandStream:               nil,
		FSUStatusReportStream:             nil,
		CenterCommandStream:               nil,
		HandleBottomControlCommandRequest: handleBottomControlRequest,
		BottomRawStreamLock:               &sync.Mutex{},
		CenterChannelStreamLock:           &sync.Mutex{},
		FSUStatusReportStreamLock:         &sync.Mutex{},
		BottomCommandStreamLock:           &sync.Mutex{},
		CenterCommandStreamLock:           &sync.Mutex{},
	}
	if !c.Connect() {
		log.Logger.Error("failed to connect grpc realdata,exit")
		os.Exit(0)
	}
	c.BottomRawValueStream()
	c.CenterChannelValueStream()
	c.FSUStatusStream()
	go c.BottomControlCommandStream()

	s := &fs.FsuServer{
		Lis:         nil,
		Server:      nil,
		RestartChan: make(chan interface{}),
		Handle:      handleDriverToFsu,
		Stream:      nil,
		StreamLock:  &sync.Mutex{},
	}
	s.PrepareRestartServer()
	s.StartServer()

	f.FsuServer = s
	f.RealdataClient = c
	f.MapNewFsuReflectType = make(map[string]reflect.Type)
	f.MapFsuInstances = make(map[int32]fi.FsuInterface)
	f.FsuInstanceLock = &sync.Mutex{}
	f.Fsus = make([]*pb.Fsu, 10)
	f.FsuTypes = make([]*pb.FsuType, 10)
	f.Configures = make([]*pb.Configure, 100)
}

//every fsu type needs to add
func (f *Fsu) RegistFsuReflectType() {
	fi.RegistJitongFsuReflectType(f.MapNewFsuReflectType)
	fi.RegistTransparentFsuReflectType(f.MapNewFsuReflectType)
}

func (f *Fsu) LoadFsu() bool {
	l1, err1 := s.GetAllFsus()
	if err1 != nil {
		log.Logger.Error("load all fsus error", zap.String("error", err1.Error()))
		return false
	}
	f.Fsus = l1.Items

	l2, err2 := s.GetAllFsuTypes()
	if err2 != nil {
		log.Logger.Error("load all fsu type error", zap.String("error", err2.Error()))
		return false
	}
	f.FsuTypes = l2.Items

	l3, err3 := s.GetAllConfigures()
	if err3 != nil {
		log.Logger.Error("load all fsus error", zap.String("error", err3.Error()))
		return false
	}
	f.Configures = l3.Items

	return true
}

func (f *Fsu) CreateFsuByAgent() {

	for _, v := range f.Fsus {
		if v.AgentId == yaml.Yaml.Agent.AgentId {
			fsuType := f.GetFsuTypeByFsuTypeId(f.FsuTypes, v.FsuTypeId)
			if fsuType == nil {
				log.Logger.Error("not found fsu type", zap.Int("fsu id", int(v.FsuId)), zap.Int("fsu type id", int(v.FsuTypeId)))
				continue
			}
			var configure *pb.Configure = nil
			if v.ConfigureId != 0 {
				configure = f.GetConfigureByConfigureId(f.Configures, v.ConfigureId)
				if configure == nil {
					log.Logger.Error("not found configure", zap.Int("fsu id", int(v.FsuId)), zap.Int("configure id", int(v.ConfigureId)))
					continue
				}
			}
			f.CreateFsu(v, fsuType, configure)
		}
	}
}

func (f *Fsu) CreateFsu(fsu *pb.Fsu, fsuType *pb.FsuType, configure *pb.Configure) {
	pType := f.MapNewFsuReflectType[fsuType.NewInstanceName]
	if pType == nil {
		log.Logger.Error("failed to create fsu", zap.Int("fsu id", int(fsu.FsuId)), zap.String("instance name", fsuType.NewInstanceName))
		return
	}
	pValue := reflect.New(pType.Elem())
	f.InitFsu(pValue, fsu, fsuType, configure)
}

func (f *Fsu) InitFsu(pValue reflect.Value, fsu *pb.Fsu, fsuType *pb.FsuType, configure *pb.Configure) {
	p := pValue.Interface().(fi.FsuInterface)
	f.FsuInstanceLock.Lock()
	f.MapFsuInstances[fsu.FsuId] = p
	f.FsuInstanceLock.Unlock()
	p.Init()
	f.setFsuParameter(p, configure, fsu)
	if fsuType.DataAnalysisMode == pb.FsuDataParsingTypeEnum_value["FSU_DATA_PARSING_TYPE_BOTTOM"] {
		f.setSignalBindingConfigure(p)
	} else if fsuType.DataAnalysisMode == pb.FsuDataParsingTypeEnum_value["FSU_DATA_PARSING_TYPE_CENTER"] {
		f.setDriverRelated(p, fsu)
	}
	p.SetFsuId(fsu.FsuId)
	p.SetAgentId(yaml.Yaml.Agent.AgentId)
	p.SetFsuTypeId(fsu.FsuTypeId)
	p.SetFSUName(fsu.FsuName)
	p.SetSendSignalRawValueFn(f.RealdataClient)
	p.SetFsuToDriverFn(f.FsuServer)
	p.Start()
	go p.Run()

}

func (f *Fsu) setFsuParameter(p fi.FsuInterface, configure *pb.Configure, fsu *pb.Fsu) {
	if configure != nil {
		var m map[string]*pb.MetadataAttribute
		err := json.Unmarshal([]byte(configure.ConfigureJson), &m)
		if err != nil {
			log.Logger.Error("failed to convert configure json", zap.Int("fsu id", int(fsu.FsuId)),
				zap.String("configure", configure.ConfigureJson), zap.String("error", err.Error()))
			return
		}
		p.SetFsuParameter(m)
	}
}

func (f *Fsu) setSignalBindingConfigure(p fi.FsuInterface) {
	m, _ := s.GetAllMetadatas()
	for _, v := range m.Items {
		if v.MetadataName == p.GetSignalBindingMetadata() {
			for _, vv := range f.Configures {
				if vv.MetadataId == v.MetadataId {
					p.SetSignalBindingConfigure(vv.ConfigureId, vv.ConfigureJson)
				}
			}
		}
	}
}

func (f *Fsu) setDriverRelated(p fi.FsuInterface, fsu *pb.Fsu) {
	d, _ := s.GetAllDrivers()
	for _, v := range d.Items {
		if v.FsuId == fsu.FsuId {
			c := f.GetConfigureByConfigureId(f.Configures, v.RelatedFlagConfigureId)
			if c != nil {
				p.SetDriverRelated(v.DriverId, v.RelatedFlagConfigureId, c.ConfigureJson)
			}
		}
	}
}

func (f *Fsu) GetFsuByFsuId(s []*pb.Fsu, fsuId int32) *pb.Fsu {
	for _, v := range s {
		if v.FsuId == fsuId {
			return v
		}
	}
	return nil
}

func (f *Fsu) GetFsuTypeByFsuTypeId(s []*pb.FsuType, fsuTypeId int32) *pb.FsuType {
	for _, v := range s {
		if v.FsuTypeId == fsuTypeId {
			return v
		}
	}
	return nil
}

func (f *Fsu) GetConfigureByConfigureId(s []*pb.Configure, configureId int32) *pb.Configure {
	for _, v := range s {
		if v.ConfigureId == configureId {
			return v
		}
	}
	return nil
}

func (f *Fsu) Stop() {
	f.FsuInstanceLock.Lock()
	for _, v := range f.MapFsuInstances {
		v.Stop()
		v.Destroy()
	}
	f.FsuInstanceLock.Unlock()
	f.FsuServer.StopServer()
	f.RealdataClient.CloseConnect()
}
