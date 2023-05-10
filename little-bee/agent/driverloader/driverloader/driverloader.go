package driverloader

import (
	"fmt"
	fc "grpc/fsuclient"
	rc "grpc/realdataclient"
	pb "message/littlebee"
	"os"
	s "service/config"
	"strings"
	"sync"
	"syscall"
	"time"
	"util/log"

	"go.uber.org/zap"
)

type DriverLoader struct {
	FsuClient       *fc.FsuClient
	RealdataClient  *rc.RealdataClient
	DriverTypes     []*pb.DriverType
	Drivers         []*pb.Driver
	Configures      []*pb.Configure
	MapDriverHandle map[int32]*DriverHandle
	MapDriver       map[int32]*Driver
}

var DriverManager *DriverLoader = &DriverLoader{
	MapDriverHandle: make(map[int32]*DriverHandle),
	MapDriver:       make(map[int32]*Driver),
}

func handleFsuToDriver(d *pb.FsuToDriver) {
	driver := DriverManager.MapDriver[d.DriverId]
	if driver != nil {
		driver.AppendResponseData(d.Response)
	}
}

func handleCenterControlRequest(d *pb.CenterControlCommandRequest) {
	driver := DriverManager.MapDriver[d.ControlParam.DriverId]
	if driver != nil {
		driver.Control(d)
	}
}

func (d *DriverLoader) InitGrpc() {
	fc := &fc.FsuClient{
		Conn:       nil,
		Client:     nil,
		Stream:     nil,
		Handle:     handleFsuToDriver,
		StreamLock: &sync.Mutex{},
	}
	if !fc.Connect() {
		log.Logger.Error("failed to connect grpc fsu,exit")
		os.Exit(0)
	}
	go fc.PayloadStream()

	rc := &rc.RealdataClient{
		Conn:                              nil,
		Client:                            nil,
		BottomRawStream:                   nil,
		CenterChannelStream:               nil,
		BottomCommandStream:               nil,
		CenterCommandStream:               nil,
		HandleCenterControlCommandRequest: handleCenterControlRequest,
		BottomRawStreamLock:               &sync.Mutex{},
		CenterChannelStreamLock:           &sync.Mutex{},
		FSUStatusReportStreamLock:         &sync.Mutex{},
		BottomCommandStreamLock:           &sync.Mutex{},
		CenterCommandStreamLock:           &sync.Mutex{},
	}
	if !rc.Connect() {
		log.Logger.Error("failed to connect grpc realdata,exit")
		os.Exit(0)
	}
	rc.CenterChannelValueStream()
	go rc.CenterControlCommandStream()

	d.FsuClient = fc
	d.RealdataClient = rc
}

func (d *DriverLoader) LoadDriver() bool {
	l1, err1 := s.GetAllDriverTypes()
	if err1 != nil {
		log.Logger.Error("load all driver types error", zap.String("error", err1.Error()))
		return false
	}
	d.DriverTypes = l1.Items

	l2, err2 := s.GetAllDrivers()
	if err2 != nil {
		log.Logger.Error("load all drivers error", zap.String("error", err2.Error()))
		return false
	}
	d.Drivers = l2.Items

	l3, err3 := s.GetAllConfigures()
	if err3 != nil {
		log.Logger.Error("load all configures error", zap.String("error", err3.Error()))
		return false
	}
	d.Configures = l3.Items

	return true
}

func (d *DriverLoader) getDriverProc(driverTypeId int32, driverDllName string) *DriverHandle {
	sp := strings.Split(driverDllName, ".")
	if len(sp) != 2 {
		log.Logger.Error("failed to get splitted string", zap.String("driver dll name", driverDllName))
		return nil
	}
	s := fmt.Sprintf(".\\driver\\%s\\%s", sp[0], driverDllName)
	dll, err := syscall.LoadLibrary(s)
	if err != nil {
		log.Logger.Error("failed to load library", zap.String("driver dll name", s), zap.String("error", err.Error()))
		return nil
	}

	newDriverProc, err := syscall.GetProcAddress(dll, "new_driver")
	if err != nil {
		log.Logger.Error("failed to get function handle:", zap.String("proc name", "new_driver"), zap.String("driver dll name", driverDllName), zap.String("error", err.Error()))
		return nil
	}

	requestProc, err := syscall.GetProcAddress(dll, "request")
	if err != nil {
		log.Logger.Error("failed to get function handle:", zap.String("proc name", "request"), zap.String("driver dll name", driverDllName), zap.String("error", err.Error()))
		return nil
	}

	controlProc, err := syscall.GetProcAddress(dll, "control")
	if err != nil {
		log.Logger.Error("failed to get function handle:", zap.String("proc name", "control"), zap.String("driver dll name", driverDllName), zap.String("error", err.Error()))
		return nil
	}

	responseProc, err := syscall.GetProcAddress(dll, "response")
	if err != nil {
		log.Logger.Error("failed to get function handle:", zap.String("proc name", "response"), zap.String("driver dll name", driverDllName), zap.String("error", err.Error()))
		return nil
	}

	timeoutProc, err := syscall.GetProcAddress(dll, "timeout")
	if err != nil {
		log.Logger.Error("failed to get function handle:", zap.String("proc name", "timeout"), zap.String("driver dll name", driverDllName), zap.String("error", err.Error()))
		return nil
	}

	destroyDriverProc, err := syscall.GetProcAddress(dll, "destroy_driver")
	if err != nil {
		log.Logger.Error("failed to get function handle:", zap.String("proc name", "destroy_driver"), zap.String("driver dll name", driverDllName), zap.String("error", err.Error()))
		return nil
	}

	dh := &DriverHandle{
		DllHandle:           dll,
		NewDriverHandle:     newDriverProc,
		RequestHandle:       requestProc,
		ControlHandle:       controlProc,
		ResponseHandle:      responseProc,
		TimeoutHandle:       timeoutProc,
		DestroyDriverHandle: destroyDriverProc,
	}
	return dh
}

func (d *DriverLoader) CreateDriverType() {
	for _, v := range d.DriverTypes {
		h := d.getDriverProc(v.DriverTypeId, v.DriverTypeFileName)
		if h != nil {
			d.MapDriverHandle[v.DriverTypeId] = h
		}
	}
}

func (d *DriverLoader) CreateDriver() {
	for _, v := range d.Drivers {
		if d.MapDriverHandle[v.DriverTypeId] != nil {
			driver := &Driver{
				DriverId:        v.DriverId,
				DriverType:      d.getDriverType(v.DriverTypeId),
				Driver:          v,
				DH:              d.MapDriverHandle[v.DriverTypeId],
				DC:              nil,
				ResponseBufLock: &sync.Mutex{},
				quit:            nil,
				DriverAction:    NEXT_STEP_ACTION_WRITE,
			}
			d.MapDriver[v.DriverId] = driver
			c := d.getConfigureByConfigureId(v.DriverParameterConfigureId)
			if c != nil {
				driver.NewDriver(c.ConfigureJson)
			} else {
				driver.NewDriver("{}")
			}
			driver.Start()
			go driver.Run()
		}
	}
}

func (d *DriverLoader) getDriverType(driverTypeId int32) *pb.DriverType {
	for _, v := range d.DriverTypes {
		if driverTypeId == v.DriverTypeId {
			return v
		}
	}
	return nil
}

func (d *DriverLoader) getConfigureByConfigureId(configureId int32) *pb.Configure {
	for _, v := range d.Configures {
		if v.ConfigureId == configureId {
			return v
		}
	}
	return nil
}

func (d *DriverLoader) Stop() {
	for _, v := range d.MapDriver {
		v.Stop()
	}
	time.Sleep(3 * time.Second)
	for _, v := range d.MapDriver {
		v.Destroy()
	}
	for _, v := range d.MapDriverHandle {
		syscall.FreeLibrary(v.DllHandle)
	}
	d.FsuClient.CloseConnect()
	d.RealdataClient.CloseConnect()
}
