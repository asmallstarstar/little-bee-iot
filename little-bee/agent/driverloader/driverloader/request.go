package driverloader

import (
	pb "message/littlebee"
	"reflect"
	"runtime"
	"sync"
	"syscall"
	"time"
	"unsafe"
	"util/log"

	"go.uber.org/zap"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type DriverNextStepAction int32

const (
	NEXT_STEP_ACTION_READ  DriverNextStepAction = 1
	NEXT_STEP_ACTION_WRITE DriverNextStepAction = 2
)

const (
	READ_WRITE_INTERVAL_MS = 100
)

type DriverHandle struct {
	DriverTypeId        int32
	DllHandle           syscall.Handle
	NewDriverHandle     uintptr
	RequestHandle       uintptr
	ControlHandle       uintptr
	ResponseHandle      uintptr
	TimeoutHandle       uintptr
	DestroyDriverHandle uintptr
}

type Driver struct {
	DriverId        int32
	DriverType      *pb.DriverType
	Driver          *pb.Driver
	DH              *DriverHandle
	DC              *DriverContext
	ResponseBufLock *sync.Mutex
	quit            chan bool
	DriverAction    DriverNextStepAction
	duration        time.Duration
	timeoutCount    int32
}

func (d *Driver) NewDriver(parameter string) bool {
	paramBytes := []byte(parameter)
	ret, _, err := syscall.SyscallN(uintptr(d.DH.NewDriverHandle), uintptr(unsafe.Pointer(&paramBytes[0])))
	if err == 0 {
		d.DC = (*DriverContext)(unsafe.Pointer(ret))
		return true
	}
	return false
}

func (d *Driver) Start() {
	d.quit = make(chan bool)
}

func (d *Driver) Run() {
	d.duration = time.Duration(d.Driver.DriverAcquisitePeriod) * time.Millisecond
	for {
		select {
		case <-d.quit:
			close(d.quit)
			runtime.Goexit()
		case <-time.After(d.duration):
			if d.DriverAction == NEXT_STEP_ACTION_WRITE {
				d.timeoutCount = 0
				_, _, err := syscall.SyscallN(uintptr(d.DH.RequestHandle), uintptr(unsafe.Pointer(d.DC)))
				if err == 0 {
					if d.DC.RequestDataLength > 0 {
						d.DriverAction = NEXT_STEP_ACTION_READ
						a := d.DC.RequestDataBuffer[0:d.DC.RequestDataLength]
						DriverManager.FsuClient.SendDriverToFsu(&pb.DriverToFsu{
							DriverId:    d.DriverId,
							FsuId:       d.Driver.FsuId,
							ConfigureId: d.Driver.RelatedFlagConfigureId,
							Request:     a,
						})
						if d.Driver.IsWriteLog {
							log.Logger.Debug("write", zap.Int32("device id", d.DriverId), zap.String("data", log.HexString(a)))
						}
					} else {
						d.DriverAction = NEXT_STEP_ACTION_WRITE
					}
				} else {
					d.DriverAction = NEXT_STEP_ACTION_WRITE
				}
				d.duration = time.Duration(d.Driver.DriverAcquisitePeriod) * time.Millisecond
			} else if d.DriverAction == NEXT_STEP_ACTION_READ {
				d.responseDataAnalysis()
			}

		}
	}
}

func (d *Driver) responseDataAnalysis() {
	d.ResponseBufLock.Lock()
	defer func() {
		if r := recover(); r != nil {
			d.ResponseBufLock.Unlock()
		} else {
			d.ResponseBufLock.Unlock()
		}
	}()
	r := &ResponseDataAnalysisResult{}
	_, _, err := syscall.SyscallN(uintptr(d.DH.ResponseHandle), uintptr(unsafe.Pointer(d.DC)), uintptr(unsafe.Pointer(r)))
	if err == 0 {
		switch r.AnalysisResult {
		case RESPONSE_DATA_ANALYSIS_SUCCESS:
			d.SetAction(NEXT_STEP_ACTION_WRITE, READ_WRITE_INTERVAL_MS)
			d.DC.ResponseDataLength = 0
			d.UploadSignalData(r.ChannelValue, int32(r.ChannelCount))
		case RESPONSE_DATA_CONTINUE_READ:
			if d.timeoutCount > d.Driver.DriverTimeoutCount {
				d.SetAction(NEXT_STEP_ACTION_WRITE, d.Driver.DriverAcquisitePeriod)
				d.DC.ResponseDataLength = 0
				d.timeoutCount = 0
				syscall.SyscallN(uintptr(d.DH.TimeoutHandle), uintptr(unsafe.Pointer(d.DC)))
			} else {
				d.timeoutCount++
				d.SetAction(NEXT_STEP_ACTION_READ, d.Driver.DriverAcquisitePeriod)
			}
		case RESPONSE_DATA_DISCARD_ALL:
			d.SetAction(NEXT_STEP_ACTION_WRITE, READ_WRITE_INTERVAL_MS)
			d.DC.ResponseDataLength = 0
		case RESPONSE_DATA_DISCARD_PART:
			d.SetAction(NEXT_STEP_ACTION_WRITE, READ_WRITE_INTERVAL_MS)
			d.MoveResposeData(int32(r.DiscardCount))
			d.DC.ResponseDataLength = d.DC.ResponseDataLength - r.DiscardCount
		case RESPONSE_DATA_VALIDATION_ERROR:
			d.SetAction(NEXT_STEP_ACTION_WRITE, READ_WRITE_INTERVAL_MS)
			d.DC.ResponseDataLength = 0
		default:
			break
		}
	}
}

func (d *Driver) UploadSignalData(channelValue [MAX_CHANNEL_COUNT]DriverChannelValue, count int32) {
	for i := 0; i < int(count); i++ {
		v := channelValue[i]

		var goBytes []byte
		if v.ValueType == uint32(pb.SignalValueTypeEnum_SIGNAL_VALUE_TYPE_STRING) {
			sliceHeader := (*reflect.SliceHeader)((unsafe.Pointer(&goBytes)))
			sliceHeader.Cap = int(v.StringLen)
			sliceHeader.Len = int(v.StringLen)
			sliceHeader.Data = uintptr(unsafe.Pointer(&v.StringValue[0]))
		} else {
			goBytes = []byte{}
		}

		DriverManager.RealdataClient.SendChannelValue(&pb.ChannelValue{
			DriverId:          d.DriverId,
			X:                 int32(v.X),
			Y:                 int32(v.Y),
			Z:                 int32(v.Z),
			ValueOccurredTime: timestamppb.Now(),
			SignalValueState:  pb.SignalValueStateEnum_SIGNAL_VALUE_STATE_VALID,
			SignalValueType:   pb.SignalValueTypeEnum(v.ValueType),
			DigitalValue:      int32(v.DigitalValue),
			AnalogValue:       float64(v.AnalogValue),
			StringValue:       string(goBytes),
		})
	}
}

func (d *Driver) MoveResposeData(discardCount int32) {
	var i int32 = 0
	for i = 0; i < int32(d.DC.ResponseDataLength)-discardCount; i++ {
		d.DC.ResponseDataBuffer[i] = d.DC.ResponseDataBuffer[discardCount+i]
	}
}

func (d *Driver) SetAction(action DriverNextStepAction, duration int32) {
	d.DriverAction = action
	d.duration = time.Duration(duration) * time.Millisecond
}

func (d *Driver) AppendResponseData(r []byte) {
	log.Logger.Debug("read", zap.Int32("device id", d.DriverId), zap.Int("length", len(r)), zap.String("data", log.HexString(r)))
	if uint32(len(r)) > (MAX_RESPONES_BUFFER_SIZE - d.DC.ResponseDataLength) {
		return
	}
	d.ResponseBufLock.Lock()
	defer d.ResponseBufLock.Unlock()
	var i uint32 = 0
	for i = 0; int(i) < len(r); i++ {
		d.DC.ResponseDataBuffer[d.DC.ResponseDataLength+i] = r[i]
	}
	d.DC.ResponseDataLength = d.DC.ResponseDataLength + uint32(len(r))
	log.Logger.Debug("response buf", zap.Int32("device id", d.DriverId), zap.Int32("length", int32(d.DC.ResponseDataLength)), zap.String("data", log.HexString(d.DC.ResponseDataBuffer[0:d.DC.ResponseDataLength])))
}

func (d *Driver) Control(c *pb.CenterControlCommandRequest) {
	v := &DriverControlValue{
		X:            uint32(c.ControlParam.X),
		Y:            uint32(c.ControlParam.Y),
		Z:            uint32(c.ControlParam.Z),
		CommandType:  uint32(c.ControlParam.ControlCommandType),
		DigitalValue: uint32(c.ControlParam.DigitalValue),
		AnalogValue:  float64(c.ControlParam.AnalogValue),
	}
	if uint32(len(c.ControlParam.StringValue)) < STRING_VALUE_BUFFER_SIZE {
		copy(v.StringValue[:], []byte(c.ControlParam.StringValue))
	}
	_, _, err := syscall.SyscallN(uintptr(d.DH.ControlHandle), uintptr(unsafe.Pointer(d.DC)), uintptr(unsafe.Pointer(v)))
	if err == 0 {
		if d.DC.ControlDataLength > 0 {
			a := d.DC.ControlDataBuffer[0 : d.DC.ControlDataLength]
			DriverManager.FsuClient.SendDriverToFsu(&pb.DriverToFsu{
				DriverId:    d.DriverId,
				FsuId:       d.Driver.FsuId,
				ConfigureId: d.Driver.RelatedFlagConfigureId,
				Request:     a,
			})
			if d.Driver.IsWriteLog {
				log.Logger.Debug("control", zap.Int32("device id", d.DriverId), zap.String("data", log.HexString(a)))
			}
		}
	}
}

func (d *Driver) Stop() {
	d.quit <- true
}

func (d *Driver) Destroy() {
	syscall.SyscallN(uintptr(d.DH.DestroyDriverHandle), uintptr(unsafe.Pointer(d.DC)))
}
