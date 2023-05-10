package fsugrpcserver

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
	RESTART_GRPC_FSU_SERVER = iota
)

type handleDriverToFsu func(d *pb.DriverToFsu)

type FsuServer struct {
	Lis         net.Listener
	Server      *grpc.Server
	RestartChan chan interface{}
	Handle      handleDriverToFsu
	Stream      pb.CenterAnalysePayload_PayloadStreamServer
	StreamLock  *sync.Mutex
	pb.UnimplementedCenterAnalysePayloadServer
}

func (s *FsuServer) PayloadStream(stream pb.CenterAnalysePayload_PayloadStreamServer) error {
	s.StreamLock.Lock()
	s.Stream = stream
	s.StreamLock.Unlock()
	for {
		driverToFsu, err := stream.Recv()
		if err != nil {
			if errors.Is(err, io.EOF) {
				continue
			} else if grpc.Code(err) == codes.Canceled {
				break
			} else {
				log.Logger.Error("rece from fSU error", zap.String("error", err.Error()))
				s.RestartChan <- RESTART_GRPC_FSU_SERVER
				break
			}
		}
		s.Handle(driverToFsu)
		log.Logger.Debug("rece from driver", zap.Int("driver id", int(driverToFsu.DriverId)), zap.String("payLoad", log.HexString(driverToFsu.Request)))
	}
	return nil
}

func (s *FsuServer) SendFsuToDriver(d *pb.FsuToDriver) {
	var err = errors.New("stream is nil")
	s.StreamLock.Lock()
	if s.Stream != nil {
		err = s.Stream.Send(d)
	}
	s.StreamLock.Unlock()
	if err != nil {
		log.Logger.Info("failed to send response rrom fsu to driver", zap.String("error ", err.Error()), zap.Binary("response", d.Response))
	}
}

func (s *FsuServer) StartServer() {
	var err error
	host := fmt.Sprintf("%s:%d", yaml.Yaml.Fsu.Host, yaml.Yaml.Fsu.Port)
	log.Logger.Info("gRPC fsu server", zap.String("host", host))
	s.Lis, err = net.Listen("tcp", host)
	if err != nil {
		log.Logger.Error("failed to start gRPC fsu server", zap.String("error msg", err.Error()))
		s.RestartChan <- RESTART_GRPC_FSU_SERVER
		return
	}
	s.Server = grpc.NewServer()
	pb.RegisterCenterAnalysePayloadServer(s.Server, s)

	log.Logger.Info("starting gRPC fsu server")
	go func() {
		if err = s.Server.Serve(s.Lis); err != nil {
			log.Logger.Error("failed to start gRPC fsu server")
			s.RestartChan <- RESTART_GRPC_FSU_SERVER
		}
	}()
}

func (s *FsuServer) StopServer() {
	log.Logger.Info("stopping gRPC fsu server")
	if s.Server != nil {
		s.Server.Stop()
	}
	if s.Lis != nil {
		s.Lis.Close()
	}
	log.Logger.Info("stopped gRPC fsu server")
}

func (s *FsuServer) PrepareRestartServer() {
	go func() {
		for {
			<-s.RestartChan
			s.StopServer()
			time.Sleep(5 * time.Second)
			go s.StartServer()
		}
	}()
}
