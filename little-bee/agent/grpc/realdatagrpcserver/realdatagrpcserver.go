package realdatagrpcserver

import (
	"errors"
	"fmt"
	"io"
	"net"
	"sync"
	"time"

	pb "message/littlebee"
	log "util/log"
	yaml "util/yaml/agent"

	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
)

const (
	RESTART_GRPC_REALDATA_SERVER = iota
)

type handleSignalRawValue func(d *pb.SignalRawValue)
type handleChannelValue func(d *pb.ChannelValue)
type handleFSUStatus func(d *pb.FSUStatus)
type handleControlCommandReport func(d *pb.ControlCommandReport)

type RealdataGrpcServer struct {
	Lis                            net.Listener
	Server                         *grpc.Server
	RestartChan                    chan interface{}
	HandleRawValue                 handleSignalRawValue
	HandleChannelValue             handleChannelValue
	HandleFSUStatus                handleFSUStatus
	HandleCommandReport            handleControlCommandReport
	BottomCommandRequestStream     pb.SignalRawValueService_BottomControlCommandStreamServer
	CenterCommandStreamStream      pb.SignalRawValueService_CenterControlCommandStreamServer
	BottomCommandRequestStreamLock *sync.Mutex
	CenterCommandStreamStreamLock  *sync.Mutex
	pb.UnimplementedSignalRawValueServiceServer
}

func (s *RealdataGrpcServer) BottomRawValueStream(stream pb.SignalRawValueService_BottomRawValueStreamServer) error {
	for {
		signalRawValue, err := stream.Recv()
		if err != nil {
			if errors.Is(err, io.EOF) {
				continue
			} else if grpc.Code(err) == codes.Canceled {
				break
			} else {
				log.Logger.Error("rece from fsu error", zap.String("error", err.Error()))
				s.RestartChan <- RESTART_GRPC_REALDATA_SERVER
				break
			}
		}
		t := signalRawValue.ValueOccurredTime.AsTime()
		log.Logger.Debug("rece from fsu", zap.Int("fsu id", int(signalRawValue.FsuId)), zap.Int("configure id", int(signalRawValue.ConfigureId)),
			zap.Timep("occurred time", &t), zap.Int("value type", int(signalRawValue.SignalValueType)), zap.Int("digital value", int(signalRawValue.DigitalValue)),
			zap.Float64("analog value", signalRawValue.AnalogValue), zap.ByteString("string value", signalRawValue.StringValue))
		s.HandleRawValue(signalRawValue)
	}
	return nil
}

func (s *RealdataGrpcServer) CenterChannelValueStream(stream pb.SignalRawValueService_CenterChannelValueStreamServer) error {
	for {
		channelValue, err := stream.Recv()
		if err != nil {
			if errors.Is(err, io.EOF) {
				continue
			} else if grpc.Code(err) == codes.Canceled {
				break
			} else {
				log.Logger.Error("rece from fsu error", zap.String("error", err.Error()))
				s.RestartChan <- RESTART_GRPC_REALDATA_SERVER
				break
			}
		}
		t := channelValue.ValueOccurredTime.AsTime()
		log.Logger.Debug("rece from driver", zap.Int("driver id", int(channelValue.DriverId)),
			zap.Int("X", int(channelValue.X)), zap.Int("Y", int(channelValue.Y)), zap.Int("Z", int(channelValue.Z)),
			zap.Timep("occurred time", &t), zap.Int("value type", int(channelValue.SignalValueType)), zap.Int("digital value", int(channelValue.DigitalValue)),
			zap.Float64("analog value", channelValue.AnalogValue), zap.String("string value", channelValue.StringValue))
		s.HandleChannelValue(channelValue)
	}
	return nil
}

func (s *RealdataGrpcServer) FSUStatusStream(stream pb.SignalRawValueService_FSUStatusStreamServer) error {
	for {
		status, err := stream.Recv()
		if err != nil {
			if errors.Is(err, io.EOF) {
				continue
			} else if grpc.Code(err) == codes.Canceled {
				break
			} else {
				log.Logger.Error("rece fsu status error", zap.String("error", err.Error()))
				s.RestartChan <- RESTART_GRPC_REALDATA_SERVER
				break
			}
		}
		//log.Logger.Debug("receive fsu status", zap.Int32("fsu id", status.FsuId),zap.String("status desc",status.StatusDesc))
		s.HandleFSUStatus(status)
	}
	return nil
}

func (s *RealdataGrpcServer) BottomControlCommandStream(stream pb.SignalRawValueService_BottomControlCommandStreamServer) error {
	s.BottomCommandRequestStreamLock.Lock()
	s.BottomCommandRequestStream = stream
	s.BottomCommandRequestStreamLock.Unlock()
	for {
		p, err := stream.Recv()
		if err != nil {
			if errors.Is(err, io.EOF) {
				continue
			} else if grpc.Code(err) == codes.Canceled {
				break
			} else {
				log.Logger.Error("rece from fsu Error", zap.String("error", err.Error()))
				s.RestartChan <- RESTART_GRPC_REALDATA_SERVER
				break
			}
		}
		log.Logger.Info("control command report from fsu", zap.ByteString("control command serialnumber", p.ControlCommandSerialNumber),
			zap.Int("control command result type", int(p.ControlCommandResultType)), zap.ByteString("control command execute description", p.ControlCommandExecuteDescription))
		s.HandleCommandReport(p)
	}
	return nil
}

func (s *RealdataGrpcServer) CenterControlCommandStream(stream pb.SignalRawValueService_CenterControlCommandStreamServer) error {
	s.CenterCommandStreamStreamLock.Lock()
	s.CenterCommandStreamStream = stream
	s.CenterCommandStreamStreamLock.Unlock()
	for {
		p, err := stream.Recv()
		if err != nil {
			if errors.Is(err, io.EOF) {
				continue
			} else if grpc.Code(err) == codes.Canceled {
				break
			} else {
				log.Logger.Error("rece from fsu error", zap.String("error", err.Error()))
				s.RestartChan <- RESTART_GRPC_REALDATA_SERVER
				break
			}
		}
		log.Logger.Info("control command report from driver", zap.ByteString("control command serialnumber", p.ControlCommandSerialNumber),
			zap.Int("control command result type", int(p.ControlCommandResultType)), zap.ByteString("control command execute description", p.ControlCommandExecuteDescription))
		s.HandleCommandReport(p)
	}
	return nil
}

func (s *RealdataGrpcServer) SendBottomControlCommandRequest(r *pb.BottomControlCommandRequest) {
	var err = errors.New("bottom command request stream is nil")
	s.BottomCommandRequestStreamLock.Lock()
	if s.BottomCommandRequestStream != nil {
		err = s.BottomCommandRequestStream.Send(r)
	}
	s.BottomCommandRequestStreamLock.Unlock()
	if err != nil {
		log.Logger.Info("failed to send bottom control command request", zap.String("error ", err.Error()))
	}
}

func (s *RealdataGrpcServer) SendCenterControlCommandRequest(r *pb.CenterControlCommandRequest) {
	var err = errors.New("center command request stream is nil")
	s.CenterCommandStreamStreamLock.Lock()
	if s.CenterCommandStreamStream != nil {
		err = s.CenterCommandStreamStream.Send(r)
	}
	s.CenterCommandStreamStreamLock.Unlock()
	if err != nil {
		log.Logger.Info("failed to send center control command request", zap.String("error ", err.Error()))
	}
}

func (s *RealdataGrpcServer) StartServer() {
	var err error
	host := fmt.Sprintf("%s:%d", yaml.Yaml.Realdata.Host, yaml.Yaml.Realdata.Port)
	log.Logger.Info("gRPC realdata server", zap.String("host", host))
	s.Lis, err = net.Listen("tcp", host)
	if err != nil {
		log.Logger.Error("failed to start gRPC realdata server ", zap.String("error", err.Error()))
		s.RestartChan <- RESTART_GRPC_REALDATA_SERVER
		return
	}
	s.Server = grpc.NewServer()
	pb.RegisterSignalRawValueServiceServer(s.Server, s)

	log.Logger.Info("starting gRPC realdata server")
	go func() {
		if err = s.Server.Serve(s.Lis); err != nil {
			log.Logger.Error("failed to start gRPC realdata server")
			s.RestartChan <- RESTART_GRPC_REALDATA_SERVER
		}
	}()
}

func (s *RealdataGrpcServer) StopServer() {
	log.Logger.Info("stopping gRPC realdata server")
	if s.Server != nil {
		s.Server.Stop()
	}
	if s.Lis != nil {
		s.Lis.Close()
	}
	log.Logger.Info("stopped gRPC realdata server")
}

func (s *RealdataGrpcServer) PrepareRestartServer() {
	go func() {
		for {
			<-s.RestartChan
			s.StopServer()
			time.Sleep(5 * time.Second)
			go s.StartServer()
		}
	}()
}
