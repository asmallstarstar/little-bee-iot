package fsuinstance

import (
	fs "grpc/fsugrpcserver"
	rc "grpc/realdataclient"
	pb "message/littlebee"
)

type FsuInterface interface {
	Init()
	SetFsuParameter(parameter map[string]*pb.MetadataAttribute)
	SetFsuId(fsuId int32)
	GetFsuId() int32
	SetAgentId(agentId int32)
	GetAgentId() int32
	SetFsuTypeId(fsuTypeId int32)
	GetFsuTypeId() int32
	SetFSUName(fsuName string)
	GetFsuTypeName() string
	GetFsuName() string
	GetFsuRunState() pb.FsuRunStateEnum
	SetSendSignalRawValueFn(realdataclient *rc.RealdataClient)
	SetFsuToDriverFn(fsuserver *fs.FsuServer)
	GetSignalBindingMetadata() string
	SetSignalBindingConfigure(configureId int32, configure string)
	SetDriverRelated(driverId int32, configureId int32, configure string)
	Start()
	Run()
	Control(controlCommandRequest *pb.BottomControlCommandRequest)
	SendToBottom(d *pb.DriverToFsu)
	Stop()
	Destroy()
}

type BaseFsu struct {
	FsuParameter                  map[string]*pb.MetadataAttribute
	FsuId                         int32
	MapDriverToRelatedConfigureId map[int32]int32 // DriverId---->RelatedFlagConfigureId
	AgentId                       int32
	FsuTypeId                     int32
	FsuName                       string
	RealdataClient                *rc.RealdataClient
	FsuServer                     *fs.FsuServer
	FsuRunState                   pb.FsuRunStateEnum
}
