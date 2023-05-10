package realdataclient

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

type handleBottomControlCommandRequest func(d *pb.BottomControlCommandRequest)
type handleCenterControlCommandRequest func(d *pb.CenterControlCommandRequest)

type RealdataClient struct {
	Conn                              *grpc.ClientConn
	Client                            pb.SignalRawValueServiceClient
	BottomRawStream                   pb.SignalRawValueService_BottomRawValueStreamClient
	CenterChannelStream               pb.SignalRawValueService_CenterChannelValueStreamClient
	FSUStatusReportStream             pb.SignalRawValueService_FSUStatusStreamClient
	BottomCommandStream               pb.SignalRawValueService_BottomControlCommandStreamClient
	CenterCommandStream               pb.SignalRawValueService_CenterControlCommandStreamClient
	HandleBottomControlCommandRequest handleBottomControlCommandRequest
	HandleCenterControlCommandRequest handleCenterControlCommandRequest
	BottomRawStreamLock               *sync.Mutex
	CenterChannelStreamLock           *sync.Mutex
	FSUStatusReportStreamLock         *sync.Mutex
	BottomCommandStreamLock           *sync.Mutex
	CenterCommandStreamLock           *sync.Mutex
}

func (r *RealdataClient) Connect() bool {
	var err error
	host := fmt.Sprintf("%s:%d", yaml.Yaml.Realdata.Host, yaml.Yaml.Realdata.Port)
	log.Logger.Info("gRPC realdata server", zap.String("Host", host))
	r.Conn, err = grpc.Dial(host, grpc.WithInsecure(), grpc.WithTimeout(15*time.Second), grpc.WithBlock(), grpc.WithBackoffMaxDelay(time.Second*5))
	if err != nil {
		log.Logger.Error("failed to  connect realdata server", zap.String("error", err.Error()))
		return false
	}
	r.Client = pb.NewSignalRawValueServiceClient(r.Conn)
	return true
}

func (r *RealdataClient) CloseConnect() {
	r.BottomRawStreamLock.Lock()
	if r.BottomRawStream != nil {
		r.BottomRawStream.CloseSend()
		r.BottomRawStream = nil
	}
	r.BottomRawStreamLock.Unlock()

	r.CenterChannelStreamLock.Lock()
	if r.CenterChannelStream != nil {
		r.CenterChannelStream.CloseSend()
		r.CenterChannelStream = nil
	}
	r.CenterChannelStreamLock.Unlock()

	r.FSUStatusReportStreamLock.Lock()
	if r.FSUStatusReportStream != nil {
		r.FSUStatusReportStream.CloseSend()
		r.FSUStatusReportStream = nil
	}
	r.FSUStatusReportStreamLock.Unlock()

	r.BottomCommandStreamLock.Lock()
	if r.BottomCommandStream != nil {
		r.BottomCommandStream.CloseSend()
		r.BottomCommandStream = nil
	}
	r.BottomCommandStreamLock.Unlock()

	r.CenterCommandStreamLock.Lock()
	if r.CenterCommandStream != nil {
		r.CenterCommandStream.CloseSend()
		r.CenterCommandStream = nil
	}
	r.CenterCommandStreamLock.Unlock()

	r.Conn.Close()
}

func (r *RealdataClient) BottomRawValueStream() {
	s, err := r.Client.BottomRawValueStream(context.Background())
	r.BottomRawStreamLock.Lock()
	r.BottomRawStream = s
	r.BottomRawStreamLock.Unlock()
	if err != nil {
		log.Logger.Error("BottomRawValueStream():failed to connect", zap.String("error", err.Error()))
	}
}

func (r *RealdataClient) CenterChannelValueStream() {
	s, err := r.Client.CenterChannelValueStream(context.Background())
	r.CenterChannelStreamLock.Lock()
	r.CenterChannelStream = s
	r.CenterChannelStreamLock.Unlock()
	if err != nil {
		log.Logger.Error("CenterChannelValueStream():failed to connect", zap.String("error", err.Error()))
	}
}

func (r *RealdataClient) FSUStatusStream() {
	s, err := r.Client.FSUStatusStream(context.Background())
	r.FSUStatusReportStreamLock.Lock()
	r.FSUStatusReportStream = s
	r.FSUStatusReportStreamLock.Unlock()
	if err != nil {
		log.Logger.Error("FSUStatusStream():failed to connect", zap.String("error", err.Error()))
	}
}

func (r *RealdataClient) BottomControlCommandStream() {
	for {
		s, err := r.Client.BottomControlCommandStream(context.Background())
		r.BottomCommandStreamLock.Lock()
		r.BottomCommandStream = s
		r.BottomCommandStreamLock.Unlock()
		if err != nil {
			log.Logger.Error("BottomControlCommandStream():failed to connect", zap.String("error", err.Error()))
			time.Sleep(5 * time.Second)
			continue
		}

		for {
			if r.BottomCommandStream == nil {
				log.Logger.Error("BottomControlCommandStream():center stream is nil")
				break
			}
			d, err := r.BottomCommandStream.Recv()
			if err != nil {
				if errors.Is(err, io.EOF) {
					//log.Logger.Error("BottomControlCommandStream():stream closed by server")
					//break
				} else {
					log.Logger.Error("BottomControlCommandStream():failed to receive message", zap.String("Error", err.Error()))
					break
				}
			}
			r.HandleBottomControlCommandRequest(d)
			log.Logger.Debug("BottomControlCommandStream():control command", zap.ByteString("operator user id", d.OperatorUserId),
				zap.ByteString("operator username", d.OperatorUserName), zap.ByteString("control command serial Number", d.ControlCommandSerialNumber), zap.Int("Fsu Id", int(d.ControlParam.FsuId)),
				zap.Int("configure id", int(d.ControlParam.ConfigureId)))
		}
	}
}

func (r *RealdataClient) CenterControlCommandStream() {
	for {
		s, err := r.Client.CenterControlCommandStream(context.Background())
		r.CenterCommandStreamLock.Lock()
		r.CenterCommandStream = s
		r.CenterCommandStreamLock.Unlock()
		if err != nil {
			log.Logger.Error("CenterControlCommandStream():failed to connect", zap.String("error", err.Error()))
			time.Sleep(5 * time.Second)
			continue
		}

		for {
			if r.CenterCommandStream == nil {
				log.Logger.Error("CenterControlCommandStream():center stream is nil")
				break
			}
			d, err := r.CenterCommandStream.Recv()
			if err != nil {
				if errors.Is(err, io.EOF) {
					//log.Logger.Error("BottomControlCommandStream():stream closed by server")
					//break
				} else {
					log.Logger.Error("BottomControlCommandStream():failed to receive message", zap.String("error", err.Error()))
					break
				}
			}
			r.HandleCenterControlCommandRequest(d)
			log.Logger.Info("CenterControlCommandStream():control command", zap.ByteString("operator user id", d.OperatorUserId),
				zap.ByteString("operator username", d.OperatorUserName), zap.ByteString("control command serial number", d.ControlCommandSerialNumber),
				zap.Int("driver id", int(d.ControlParam.DriverId)), zap.Int("X", int(d.ControlParam.X)), zap.Int("Y", int(d.ControlParam.Y)), zap.Int("Z", int(d.ControlParam.Z)))
		}
	}
}

func (r *RealdataClient) BottomControlCommandReport(p *pb.ControlCommandReport) {
	var err = errors.New("bottom command stream is nil")
	r.BottomCommandStreamLock.Lock()
	if r.BottomCommandStream != nil {
		err = r.BottomCommandStream.Send(p)
	}
	r.BottomCommandStreamLock.Unlock()
	if err != nil {
		log.Logger.Info("failed To send control command report", zap.ByteString("control command serialnumber", p.ControlCommandSerialNumber),
			zap.Int("control command result type", int(p.ControlCommandResultType)), zap.ByteString("control command execute description", p.ControlCommandExecuteDescription),
			zap.String("error", err.Error()))
	}
}

func (r *RealdataClient) CenterControlCommandReport(p *pb.ControlCommandReport) {
	var err = errors.New("center command stream is nil")
	r.CenterCommandStreamLock.Lock()
	if r.CenterCommandStream != nil {
		err = r.CenterCommandStream.Send(p)
	}
	r.CenterCommandStreamLock.Unlock()
	if err != nil {
		log.Logger.Info("failed to send control command report ", zap.ByteString("control command serialnumber", p.ControlCommandSerialNumber),
			zap.Int("control command result type", int(p.ControlCommandResultType)), zap.ByteString("control command execute description", p.ControlCommandExecuteDescription))
	}
}

func (r *RealdataClient) SendSignalRawValue(p *pb.SignalRawValue) {
	var err = errors.New("bottom raw stream is nil")
	r.BottomRawStreamLock.Lock()
	if r.BottomRawStream != nil {
		err = r.BottomRawStream.Send(p)
	}
	r.BottomRawStreamLock.Unlock()
	if err != nil {
		r.BottomRawValueStream()
		log.Logger.Info("failed to send raw value", zap.String("error ", err.Error()))
	}
}

func (r *RealdataClient) SendChannelValue(p *pb.ChannelValue) {
	var err = errors.New("center stream is nil")
	r.CenterChannelStreamLock.Lock()
	if r.CenterChannelStream != nil {
		err = r.CenterChannelStream.Send(p)
	}
	r.CenterChannelStreamLock.Unlock()
	if err != nil {
		r.CenterChannelValueStream()
		log.Logger.Info("failed to send channel value", zap.String("error ", err.Error()))
	}
}

func (r *RealdataClient) SendFSUStatus(p *pb.FSUStatus) {
	var err = errors.New("fsu status stream is nil")
	r.FSUStatusReportStreamLock.Lock()
	if r.FSUStatusReportStream != nil {
		err = r.FSUStatusReportStream.Send(p)
	}
	r.FSUStatusReportStreamLock.Unlock()
	if err != nil {
		r.FSUStatusStream()
		log.Logger.Info("failed to send fsu status", zap.String("error ", err.Error()))
	}
}
