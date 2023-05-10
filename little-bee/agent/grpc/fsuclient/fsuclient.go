package fsuclient

import (
	"context"
	"errors"
	"fmt"
	"io"
	"sync"
	"time"

	pb "message/littlebee"
	log "util/log"
	yaml "util/yaml/agent"

	"go.uber.org/zap"
	"google.golang.org/grpc"
)

type FsuClient struct {
	Conn       *grpc.ClientConn
	Client     pb.CenterAnalysePayloadClient
	Stream     pb.CenterAnalysePayload_PayloadStreamClient
	Handle     handleFsuToDriver
	StreamLock *sync.Mutex
}

type handleFsuToDriver func(d *pb.FsuToDriver)

func (r *FsuClient) Connect() bool {
	var err error
	host := fmt.Sprintf("%s:%d", yaml.Yaml.Fsu.Host, yaml.Yaml.Fsu.Port)
	log.Logger.Info("gRPC fsu server", zap.String("host", host))
	r.Conn, err = grpc.Dial(host, grpc.WithInsecure(), grpc.WithTimeout(15*time.Second), grpc.WithBlock(), grpc.WithBackoffMaxDelay(time.Second*5))
	if err != nil {
		log.Logger.Error("failed to connect fsu server", zap.String("error", err.Error()))
		return false
	}
	r.Client = pb.NewCenterAnalysePayloadClient(r.Conn)
	return true
}

func (r *FsuClient) CloseConnect() {
	r.StreamLock.Lock()
	if r.Stream != nil {
		r.Stream.CloseSend()
		r.Stream = nil
	}
	r.StreamLock.Unlock()

	r.Conn.Close()
}

func (r *FsuClient) PayloadStream() {
	for {
		s, err := r.Client.PayloadStream(context.Background())
		r.StreamLock.Lock()
		r.Stream = s
		r.StreamLock.Unlock()
		if err != nil {
			log.Logger.Error("PayloadStream():failed to connect", zap.String("error", err.Error()))
			time.Sleep(5 * time.Second)
			continue
		}

		for {
			if r.Stream == nil {
				log.Logger.Error("PayloadStream():center stream is nil")
				break
			}
			d, err := r.Stream.Recv()
			if err == io.EOF {
				log.Logger.Error("PayloadStream():connection closed by server")
				break
			}
			if err != nil {
				log.Logger.Error("PayloadStream():failed to receive message", zap.String("error", err.Error()))
				break
			}
			r.Handle(d)
			//log.Logger.Debug("PayloadStream():response", zap.Int("driver id", int(d.DriverId)),
			//	zap.Int("fsuid", int(d.FsuId)), zap.String("response", log.HexString(d.Response)))
		}
	}
}

func (s *FsuClient) SendDriverToFsu(d *pb.DriverToFsu) {
	var err = errors.New("stream is nil")
	s.StreamLock.Lock()
	if s.Stream != nil {
		err = s.Stream.Send(d)
	}
	s.StreamLock.Unlock()
	if err != nil {
		log.Logger.Info("failed to send request from driver to fsu", zap.String("error ", err.Error()), zap.Binary("request", d.Request))
	}
}
