package realdataserver

import (
	"fmt"
	rs "grpc/realdatagrpcserver"
	"message/littlebee"
	pb "message/littlebee"
	m "mqtt"
	service "service/config"
	"strconv"
	"sync"
	"util/date"
	"util/log"
	yaml "util/yaml/agent"

	MQTT "github.com/eclipse/paho.mqtt.golang"

	"github.com/golang/protobuf/ptypes/timestamp"
	"go.uber.org/zap"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type RealdataServer struct {
	RealdataGrpcServer     *rs.RealdataGrpcServer
	MapLogicObjs           map[int32]*pb.LogicObject
	MapSignals             map[int32]*pb.Signal
	MapSignalAi            map[int32]*pb.SignalAi
	MapSignalDi            map[int32]*pb.SignalDi
	MapSignalSi            map[int32]*pb.SignalSi
	MapThreshold           map[int32]*pb.SignalThreshold
	MapAcquisition         map[int32]*pb.SignalAcquisition
	MapSignalCurrentStatus map[int32]*pb.SignalCurrentStata
	MapActiveAlarm         map[int32]*pb.Alarm
	MapRealValue           map[int32]*pb.SignalRealValue
}

var RealdataServerManager *RealdataServer = &RealdataServer{
	MapLogicObjs:           make(map[int32]*pb.LogicObject),
	MapSignals:             make(map[int32]*pb.Signal),
	MapSignalAi:            make(map[int32]*pb.SignalAi),
	MapSignalDi:            make(map[int32]*pb.SignalDi),
	MapSignalSi:            make(map[int32]*pb.SignalSi),
	MapThreshold:           make(map[int32]*pb.SignalThreshold),
	MapAcquisition:         make(map[int32]*pb.SignalAcquisition),
	MapSignalCurrentStatus: make(map[int32]*pb.SignalCurrentStata),
	MapActiveAlarm:         make(map[int32]*pb.Alarm),
	MapRealValue:           make(map[int32]*pb.SignalRealValue),
}

type SignalDeviceValue struct {
	SignalId          int32
	ValueOccurredTime *timestamp.Timestamp
	SignalValueState  pb.SignalValueStateEnum
	SignalValueType   pb.SignalValueTypeEnum
	DigitalValue      int32
	AnalogValue       float64
	StringValue       string
}

func handleRawValue(d *pb.SignalRawValue) {
	for k, v := range RealdataServerManager.MapAcquisition {
		s := RealdataServerManager.MapSignals[v.SignalId]
		if s != nil && v.FsuId == d.FsuId && v.SignalBindConfigureId == d.ConfigureId {
			if s.SignalType == pb.SignalTypeEnum_SIGNAL_TYPE_AI {
				var av float64 = 0.0
				if d.SignalValueType == pb.SignalValueTypeEnum_SIGNAL_VALUE_TYPE_ANALOG {
					av = d.AnalogValue
				}
				if d.SignalValueType == pb.SignalValueTypeEnum_SIGNAL_VALUE_TYPE_DIGITAL {
					av = float64(d.DigitalValue)
				}
				if d.SignalValueType == pb.SignalValueTypeEnum_SIGNAL_VALUE_TYPE_STRING {
					f, _ := strconv.ParseFloat(string(d.StringValue), 64)
					av = f
				}
				v := &SignalDeviceValue{
					SignalId:          k,
					ValueOccurredTime: timestamppb.Now(),
					SignalValueState:  d.SignalValueState,
					SignalValueType:   pb.SignalValueTypeEnum_SIGNAL_VALUE_TYPE_ANALOG,
					DigitalValue:      0,
					AnalogValue:       av,
					StringValue:       "",
				}
				RealdataServerManager.calculateSignalValue(v)
			}
			if s.SignalType == pb.SignalTypeEnum_SIGNAL_TYPE_DI {
				var dv int32 = 0
				if d.SignalValueType == pb.SignalValueTypeEnum_SIGNAL_VALUE_TYPE_ANALOG {
					dv = int32(d.AnalogValue)
				}
				if d.SignalValueType == pb.SignalValueTypeEnum_SIGNAL_VALUE_TYPE_DIGITAL {
					dv = d.DigitalValue
				}
				if d.SignalValueType == pb.SignalValueTypeEnum_SIGNAL_VALUE_TYPE_STRING {
					i, _ := strconv.Atoi(string(d.StringValue))
					dv = int32(i)
				}
				v := &SignalDeviceValue{
					SignalId:          k,
					ValueOccurredTime: timestamppb.Now(),
					SignalValueState:  d.SignalValueState,
					SignalValueType:   pb.SignalValueTypeEnum_SIGNAL_VALUE_TYPE_DIGITAL,
					DigitalValue:      dv,
					AnalogValue:       0.0,
					StringValue:       "",
				}
				RealdataServerManager.calculateSignalValue(v)
			}
			if s.SignalType == pb.SignalTypeEnum_SIGNAL_TYPE_SI {
				var sv string = ""
				if d.SignalValueType == pb.SignalValueTypeEnum_SIGNAL_VALUE_TYPE_ANALOG {
					sv = strconv.FormatFloat(d.AnalogValue, 'f', 2, 64)
				}
				if d.SignalValueType == pb.SignalValueTypeEnum_SIGNAL_VALUE_TYPE_DIGITAL {
					sv = strconv.Itoa(int(d.DigitalValue))
				}
				if d.SignalValueType == pb.SignalValueTypeEnum_SIGNAL_VALUE_TYPE_STRING {
					sv = string(d.StringValue)
				}
				v := &SignalDeviceValue{
					SignalId:          k,
					ValueOccurredTime: timestamppb.Now(),
					SignalValueState:  d.SignalValueState,
					SignalValueType:   pb.SignalValueTypeEnum_SIGNAL_VALUE_TYPE_STRING,
					DigitalValue:      0,
					AnalogValue:       0.0,
					StringValue:       sv,
				}
				RealdataServerManager.calculateSignalValue(v)
			}
		}
	}
}

func handleChannelValue(d *pb.ChannelValue) {
	for k, v := range RealdataServerManager.MapAcquisition {
		if v.DriverId == d.DriverId && v.X == d.X && v.Y == d.Y && v.Z == d.Z {
			v := &SignalDeviceValue{
				SignalId:          k,
				ValueOccurredTime: timestamppb.Now(),
				SignalValueState:  d.SignalValueState,
				SignalValueType:   d.SignalValueType,
				DigitalValue:      d.DigitalValue,
				AnalogValue:       d.AnalogValue,
				StringValue:       d.StringValue,
			}
			RealdataServerManager.calculateSignalValue(v)
		}
	}
}

func handleFSUStatus(d *pb.FSUStatus) {
	message, _ := proto.Marshal(d)
	m.PublishFsuStatus(message)
}

func handleCommandReport(d *pb.ControlCommandReport) {
	message, _ := proto.Marshal(d)
	if e := m.Publish(m.TOPIC_TYPE_CONTROL_REPORT, m.MQTT_QOS_2, false, message); e != nil {
		log.Logger.Error("failed to publish command report", zap.String("control command sn", string(d.ControlCommandSerialNumber)),
			zap.Int32("result", int32(d.ControlCommandResultType)), zap.String("error", e.Error()))
	}
	if e := service.UpdateControlCommandResult(&pb.ControlCommandRecord{
		ControlCommandSerialNumber: string(d.ControlCommandSerialNumber),
		ControlCommandResult:       d.ControlCommandResultType,
	}); e != nil {
		log.Logger.Error("failed to record command report", zap.String("control command sn", string(d.ControlCommandSerialNumber)),
			zap.Int32("result", int32(d.ControlCommandResultType)), zap.String("error", e.Error()))
	}
}

func (r *RealdataServer) calculateSignalValue(v *SignalDeviceValue) {
	var realvalue *pb.SignalRealValue = nil
	switch v.SignalValueType {
	case pb.SignalValueTypeEnum_SIGNAL_VALUE_TYPE_ANALOG:
		realvalue = r.calculateAiValue(v)
	case pb.SignalValueTypeEnum_SIGNAL_VALUE_TYPE_DIGITAL:
		realvalue = r.calculateDiValue(v)
	case pb.SignalValueTypeEnum_SIGNAL_VALUE_TYPE_STRING:
		realvalue = r.calculateSiValue(v)
	}
	if realvalue == nil {
		return
	}
	//publish realdata
	message, _ := proto.Marshal(realvalue)
	m.PublishRealdata(message)

	//alarm analysis
	r.analysisAlarm(realvalue)
}

func (r *RealdataServer) analysisAlarm(realvalue *pb.SignalRealValue) {
	if realvalue.SignalValueState != pb.SignalValueStateEnum_SIGNAL_VALUE_STATE_VALID {
		r.updateSignalCurrentStatus(realvalue.SignalId, pb.SignalValueStateEnum_SIGNAL_VALUE_STATE_INVALID,
			realvalue.SignalRunState, realvalue.AlarmThreshholdNumber, realvalue.ValueOccurredTime, realvalue.AlarmLevelNumber)
		return
	}
	currentstatus := r.MapSignalCurrentStatus[realvalue.SignalId]

	if currentstatus.SignalValueState == pb.SignalValueStateEnum_SIGNAL_VALUE_STATE_VALID &&
		realvalue.SignalRunState != pb.SignalRunStateEnum_SIGNAL_RUN_STATE_UNKNOWN &&
		currentstatus.SignalRunState != pb.SignalRunStateEnum_SIGNAL_RUN_STATE_UNKNOWN {
		rawevent := r.generateEventBysignalRealValue(realvalue, currentstatus)
		if rawevent != nil {
			if rawevent.SignalEventType == pb.SignalEventTypeEnum_SIGNAL_EVENT_TYPE_BEGIN {
				//alarm begin
				ab := r.generateBeginAlarm(rawevent)
				service.CreateAlarm(r.convertBeginAlarmToAlarm(ab), 0) //record
				mab, _ := proto.Marshal(ab)
				m.Publish(m.TOPIC_TYPE_BEGIN_ALARM, m.MQTT_QOS_2, false, mab) //publish
				log.Logger.Warn("alarm begin", zap.String("alarm serial number", ab.AlarmSerialNumber))

				//set signal current status
				r.updateSignalCurrentStatus(realvalue.SignalId, pb.SignalValueStateEnum_SIGNAL_VALUE_STATE_VALID,
					pb.SignalRunStateEnum_SIGNAL_RUN_STATE_ALARM, rawevent.AlarmThreshholdNumber,
					rawevent.EventOccurredTime, rawevent.AlarmLevelNumber)

			} else if rawevent.SignalEventType == pb.SignalEventTypeEnum_SIGNAL_EVENT_TYPE_END {
				//alarm end
				ae := r.generateEndAlarm(rawevent)
				service.EndAlarmBySerialNumber(r.convertEndAlarmToAlarm(ae)) //record
				mae, _ := proto.Marshal(ae)
				m.Publish(m.TOPIC_TYPE_END_ALARM, m.MQTT_QOS_2, false, mae) //publish
				log.Logger.Warn("alarm    end", zap.String("alarm serial number", ae.AlarmSerialNumber))

				//set signal current status
				r.updateSignalCurrentStatus(realvalue.SignalId, pb.SignalValueStateEnum_SIGNAL_VALUE_STATE_VALID,
					pb.SignalRunStateEnum_SIGNAL_RUN_STATE_NORMAL, realvalue.AlarmThreshholdNumber,
					rawevent.AlarmBeginTime, realvalue.AlarmLevelNumber)
			}
		} else {
			//set signal current status
			r.updateSignalCurrentStatus(realvalue.SignalId, pb.SignalValueStateEnum_SIGNAL_VALUE_STATE_VALID,
				realvalue.SignalRunState, realvalue.AlarmThreshholdNumber,
				rawevent.AlarmBeginTime, realvalue.AlarmLevelNumber)
		}
	} else {
		//set signal current status
		r.updateSignalCurrentStatus(realvalue.SignalId, pb.SignalValueStateEnum_SIGNAL_VALUE_STATE_VALID,
			currentstatus.SignalRunState, currentstatus.AlarmThreshholdNumber, currentstatus.AlarmBeginTime, currentstatus.AlarmLevelNumber)
	}
}

func (r *RealdataServer) generateEventBysignalRealValue(realvalue *pb.SignalRealValue, currentstate *pb.SignalCurrentStata) *pb.SignalRawEvent {
	e := &pb.SignalRawEvent{
		SignalId:              realvalue.SignalId,
		SignalName:            r.MapSignals[realvalue.SignalId].SignalName,
		SignalType:            r.MapSignals[realvalue.SignalId].SignalType,
		EventOccurredTime:     realvalue.ValueOccurredTime,
		EventDesc:             "",
		SignalEventType:       pb.SignalEventTypeEnum_SIGNAL_EVENT_TYPE_UNKNOWN,
		AlarmBeginTime:        realvalue.ValueOccurredTime,
		SignalValueState:      pb.SignalValueStateEnum_SIGNAL_VALUE_STATE_VALID,
		SignalRunState:        pb.SignalRunStateEnum_SIGNAL_RUN_STATE_UNKNOWN,
		AlarmLevelNumber:      -1,
		AlarmThreshholdNumber: -1,
		SignalValueType:       realvalue.SignalValueType,
		DigitalValue:          realvalue.DigitalValue,
		AnalogValue:           realvalue.AnalogValue,
		ValueUnit:             realvalue.ValueUnit,
		StringValue:           realvalue.StringValue,
	}
	if realvalue.SignalRunState != currentstate.SignalRunState {
		if realvalue.SignalRunState == pb.SignalRunStateEnum_SIGNAL_RUN_STATE_ALARM {
			threhold := r.getThresholdByNumber(realvalue.SignalId, realvalue.AlarmThreshholdNumber)
			e.AlarmThreshholdNumber = threhold.ThresholdNumber
			e.AlarmLevelNumber = threhold.AlarmLevelNumber
			e.SignalRunState = pb.SignalRunStateEnum_SIGNAL_RUN_STATE_ALARM
			e.SignalEventType = pb.SignalEventTypeEnum_SIGNAL_EVENT_TYPE_BEGIN
			e.EventDesc = threhold.UpNotifyNote
		} else if realvalue.SignalRunState == pb.SignalRunStateEnum_SIGNAL_RUN_STATE_NORMAL {
			threhold := r.getThresholdByNumber(realvalue.SignalId, currentstate.AlarmThreshholdNumber)
			e.AlarmThreshholdNumber = threhold.ThresholdNumber
			e.AlarmLevelNumber = threhold.AlarmLevelNumber
			e.SignalRunState = pb.SignalRunStateEnum_SIGNAL_RUN_STATE_NORMAL
			e.SignalEventType = pb.SignalEventTypeEnum_SIGNAL_EVENT_TYPE_END
			e.EventDesc = threhold.DownNotifyNote
			e.AlarmBeginTime = currentstate.AlarmBeginTime
		}
	} else {
		if realvalue.SignalRunState == pb.SignalRunStateEnum_SIGNAL_RUN_STATE_ALARM &&
			realvalue.AlarmThreshholdNumber != currentstate.AlarmThreshholdNumber {

			threhold := r.getThresholdByNumber(realvalue.SignalId, realvalue.AlarmThreshholdNumber)
			e.AlarmThreshholdNumber = threhold.ThresholdNumber
			e.AlarmLevelNumber = threhold.AlarmLevelNumber
			e.SignalRunState = pb.SignalRunStateEnum_SIGNAL_RUN_STATE_NORMAL //end alarm
			e.SignalEventType = pb.SignalEventTypeEnum_SIGNAL_EVENT_TYPE_END
			e.EventDesc = threhold.DownNotifyNote
		}
	}
	return e
}

func (r *RealdataServer) generateAlarmSerialNumber(signalId int32, alarmbegin *timestamp.Timestamp) string {
	return fmt.Sprintf("%s@%d", date.FormatProtobufTime(alarmbegin), signalId)
}

func (r *RealdataServer) updateSignalCurrentStatus(id int32, valuestate pb.SignalValueStateEnum, runstate pb.SignalRunStateEnum,
	thresholdnumber int32, t *timestamp.Timestamp, levelnumber int32) {
	r.MapSignalCurrentStatus[id] = &pb.SignalCurrentStata{
		SignalId:              id,
		SignalValueState:      valuestate,
		SignalRunState:        runstate,
		AlarmThreshholdNumber: thresholdnumber,
		AlarmBeginTime:        t,
		AlarmLevelNumber:      levelnumber,
	}
}

func (r *RealdataServer) calculateAiValue(v *SignalDeviceValue) *pb.SignalRealValue {
	s := r.MapSignals[v.SignalId]
	sai := r.MapSignalAi[v.SignalId]
	currentStatus := r.MapSignalCurrentStatus[v.SignalId]
	if s == nil || sai == nil || currentStatus == nil {
		return nil
	}
	analog := sai.MultipleRate*v.AnalogValue + sai.OffsetValue
	//analog, _ := strconv.ParseFloat(fmt.Sprintf("%."+strconv.Itoa(int(sai.DecimalPrecision))+"f", sai.MultipleRate*v.AnalogValue+sai.OffsetValue), 32)
	realvalue := &pb.SignalRealValue{
		SignalId:              v.SignalId,
		SignalUnificationId:   s.SignalUnificationId,
		SignalName:            s.SignalName,
		SignalType:            s.SignalType,
		SignalValueState:      v.SignalValueState,
		SignalRunState:        pb.SignalRunStateEnum_SIGNAL_RUN_STATE_UNKNOWN,
		AlarmThreshholdNumber: -1,
		AlarmLevelNumber:      -1,
		ValueOccurredTime:     timestamppb.Now(),
		SignalValueType:       pb.SignalValueTypeEnum_SIGNAL_VALUE_TYPE_ANALOG,
		DigitalValue:          0,
		AnalogValue:           analog,
		DecimalPrecision:      sai.DecimalPrecision,
		ValueUnit:             sai.ValueUnit,
		StringValue:           "",
		DigitalValueDesc:      "",
	}
	thresholds := r.getThresholdBySignalId(v.SignalId)
	if realvalue.SignalValueState == pb.SignalValueStateEnum_SIGNAL_VALUE_STATE_VALID {
		if currentStatus.SignalValueState == pb.SignalValueStateEnum_SIGNAL_VALUE_STATE_VALID {
			if len(thresholds) > 0 {
				threshold := r.getInversionThresholds(float32(realvalue.AnalogValue), currentStatus, thresholds)
				if threshold != nil && threshold.IsNotify {
					realvalue.SignalRunState = pb.SignalRunStateEnum_SIGNAL_RUN_STATE_ALARM
					realvalue.AlarmThreshholdNumber = threshold.ThresholdNumber
					realvalue.AlarmLevelNumber = threshold.AlarmLevelNumber
				} else {
					realvalue.SignalRunState = pb.SignalRunStateEnum_SIGNAL_RUN_STATE_NORMAL
				}
			} else {
				realvalue.SignalRunState = pb.SignalRunStateEnum_SIGNAL_RUN_STATE_NORMAL
			}
		} else {
			currentStatus.SignalValueState = pb.SignalValueStateEnum_SIGNAL_VALUE_STATE_VALID
		}
	} else {
		realvalue.SignalRunState = pb.SignalRunStateEnum_SIGNAL_RUN_STATE_UNKNOWN
	}

	return realvalue
}

func (r *RealdataServer) getInversionThresholds(currentValue float32, currentStatus *pb.SignalCurrentStata, thresholds map[int32]*pb.SignalThreshold) *pb.SignalThreshold {
	var maxLevelThreshold *pb.SignalThreshold = nil
	if currentStatus.SignalValueState == pb.SignalValueStateEnum_SIGNAL_VALUE_STATE_VALID {
		tempthreshholdList := make([]*pb.SignalThreshold, 0)
		for _, v := range thresholds {
			if v.ThresholdDirection == pb.ThresholdDirectionEnum_THRESHOLD_DIRECTION_GREAT {
				if currentValue > float32(v.ThresholdValueAi) ||
					(currentStatus.AlarmThreshholdNumber == v.ThresholdNumber && currentValue > float32(v.ThresholdValueAi-v.ThresholdDeadArea)) {
					maxLevelThreshold = v
					tempthreshholdList = append(tempthreshholdList, v)
				}
			} else if v.ThresholdDirection == pb.ThresholdDirectionEnum_THRESHOLD_DIRECTION_LESS {
				if currentValue < float32(v.ThresholdValueAi) ||
					(currentStatus.AlarmThreshholdNumber == v.ThresholdNumber && currentValue < float32(v.ThresholdValueAi+v.ThresholdDeadArea)) {
					maxLevelThreshold = v
					tempthreshholdList = append(tempthreshholdList, v)
				}
			}
		}
		for _, v := range tempthreshholdList {
			if v.AlarmLevelNumber > maxLevelThreshold.AlarmLevelNumber {
				maxLevelThreshold = v
			}
		}
	}
	return maxLevelThreshold
}

func (r *RealdataServer) calculateDiValue(v *SignalDeviceValue) *pb.SignalRealValue {
	s := r.MapSignals[v.SignalId]
	sdi := r.MapSignalDi[v.SignalId]
	currentStatus := r.MapSignalCurrentStatus[v.SignalId]
	if s == nil || sdi == nil || currentStatus == nil {
		return nil
	}
	realvalue := &pb.SignalRealValue{
		SignalId:              v.SignalId,
		SignalUnificationId:   s.SignalUnificationId,
		SignalName:            s.SignalName,
		SignalType:            s.SignalType,
		SignalValueState:      v.SignalValueState,
		SignalRunState:        pb.SignalRunStateEnum_SIGNAL_RUN_STATE_UNKNOWN,
		AlarmThreshholdNumber: -1,
		AlarmLevelNumber:      -1,
		ValueOccurredTime:     timestamppb.Now(),
		SignalValueType:       pb.SignalValueTypeEnum_SIGNAL_VALUE_TYPE_DIGITAL,
		DigitalValue:          v.DigitalValue,
		AnalogValue:           0.0,
		DecimalPrecision:      0,
		ValueUnit:             "",
		StringValue:           "",
		DigitalValueDesc:      "",
	}
	if realvalue.SignalValueState == pb.SignalValueStateEnum_SIGNAL_VALUE_STATE_VALID {
		if sdi.IsFlip {
			if v.DigitalValue == int32(pb.SignalDiValueEnum_SIGNAL_DI_VALUE_0) {
				realvalue.DigitalValue = int32(pb.SignalDiValueEnum_SIGNAL_DI_VALUE_1)
			}
			if v.DigitalValue == int32(pb.SignalDiValueEnum_SIGNAL_DI_VALUE_1) {
				realvalue.DigitalValue = int32(pb.SignalDiValueEnum_SIGNAL_DI_VALUE_0)
			}
		} else {
			realvalue.DigitalValue = v.DigitalValue
		}
		if realvalue.DigitalValue == int32(pb.SignalDiValueEnum_SIGNAL_DI_VALUE_0) {
			realvalue.DigitalValueDesc = sdi.LowDesc
		}
		if realvalue.DigitalValue == int32(pb.SignalDiValueEnum_SIGNAL_DI_VALUE_1) {
			realvalue.DigitalValueDesc = sdi.HighDesc
		}

		if currentStatus.SignalValueState == pb.SignalValueStateEnum_SIGNAL_VALUE_STATE_VALID {
			thresholds := r.getThresholdBySignalId(v.SignalId)
			if len(thresholds) > 0 {
				for _, threshold := range thresholds {
					if realvalue.DigitalValue == threshold.ThresholdValueDi && threshold.IsNotify {
						realvalue.SignalRunState = pb.SignalRunStateEnum_SIGNAL_RUN_STATE_ALARM
						realvalue.AlarmThreshholdNumber = threshold.ThresholdNumber
						realvalue.AlarmLevelNumber = threshold.AlarmLevelNumber
					} else {
						realvalue.SignalRunState = pb.SignalRunStateEnum_SIGNAL_RUN_STATE_NORMAL
					}
					break
				}
			} else {
				realvalue.SignalRunState = pb.SignalRunStateEnum_SIGNAL_RUN_STATE_NORMAL
			}
		} else {
			currentStatus.SignalValueState = pb.SignalValueStateEnum_SIGNAL_VALUE_STATE_VALID
		}
	} else {
		realvalue.SignalRunState = pb.SignalRunStateEnum_SIGNAL_RUN_STATE_UNKNOWN
	}

	return realvalue
}

func (r *RealdataServer) calculateSiValue(v *SignalDeviceValue) *pb.SignalRealValue {
	s := r.MapSignals[v.SignalId]
	ssi := r.MapSignalSi[v.SignalId]
	if s == nil || ssi == nil {
		return nil
	}
	realvalue := &pb.SignalRealValue{
		SignalId:              v.SignalId,
		SignalUnificationId:   s.SignalUnificationId,
		SignalName:            s.SignalName,
		SignalType:            s.SignalType,
		SignalValueState:      v.SignalValueState,
		SignalRunState:        pb.SignalRunStateEnum_SIGNAL_RUN_STATE_NORMAL,
		AlarmThreshholdNumber: -1,
		AlarmLevelNumber:      -1,
		ValueOccurredTime:     timestamppb.Now(),
		SignalValueType:       pb.SignalValueTypeEnum_SIGNAL_VALUE_TYPE_STRING,
		DigitalValue:          0,
		AnalogValue:           0,
		ValueUnit:             "",
		StringValue:           string(v.StringValue),
		DigitalValueDesc:      "",
	}
	return realvalue
}

func (r *RealdataServer) InitMQTT() bool {
	return m.Connect(yaml.Yaml.MQTT.Broker, yaml.Yaml.MQTT.UserName, yaml.Yaml.MQTT.Password, yaml.Yaml.MQTT.RealdataClientId) && r.SubControlCommand()
}

func (r *RealdataServer) SubControlCommand() bool {
	err := m.Subscribe(m.TOPIC_TYPE_CONTROL, m.MQTT_QOS_2, func(c MQTT.Client, m MQTT.Message) {
		cp := &pb.CommandParam{}
		e := proto.Unmarshal(m.Payload(), cp)
		if e == nil {
			sa, err1 := service.RetrieveSignalAcquisition(cp.SignalId)
			if err1 != nil {
				handleCommandReport(&pb.ControlCommandReport{
					ControlCommandSerialNumber:       []byte(cp.ControlCommandSerialNumber),
					ControlCommandReportOccurredTime: timestamppb.Now(),
					ControlCommandResultType:         littlebee.ControlCommandResultEnum_CONTROL_COMMAND_RESULT_FAIL,
					ControlCommandExecuteDescription: []byte{},
				})
				return
			}
			fsu, err2 := service.RetrieveFsu(sa.FsuId)
			if err2 != nil {
				handleCommandReport(&pb.ControlCommandReport{
					ControlCommandSerialNumber:       []byte(cp.ControlCommandSerialNumber),
					ControlCommandReportOccurredTime: timestamppb.Now(),
					ControlCommandResultType:         littlebee.ControlCommandResultEnum_CONTROL_COMMAND_RESULT_FAIL,
					ControlCommandExecuteDescription: []byte{},
				})
				return
			}
			fsutype, err3 := service.RetrieveFsuType(fsu.FsuTypeId)
			if err3 != nil {
				handleCommandReport(&pb.ControlCommandReport{
					ControlCommandSerialNumber:       []byte(cp.ControlCommandSerialNumber),
					ControlCommandReportOccurredTime: timestamppb.Now(),
					ControlCommandResultType:         littlebee.ControlCommandResultEnum_CONTROL_COMMAND_RESULT_FAIL,
					ControlCommandExecuteDescription: []byte{},
				})
				return
			}
			if fsutype.DataAnalysisMode == int32(pb.FsuDataParsingTypeEnum_FSU_DATA_PARSING_TYPE_BOTTOM) {
				r.RealdataGrpcServer.SendBottomControlCommandRequest(&pb.BottomControlCommandRequest{
					OperatorUserId:             []byte(strconv.Itoa(int(cp.OperatorUserId))),
					OperatorUserName:           []byte(cp.OperatorUserName),
					ControlCommandSerialNumber: []byte(cp.ControlCommandSerialNumber),
					ControlParam: &pb.BottomCommandParam{
						FsuId:              sa.FsuId,
						ConfigureId:        fsu.ConfigureId,
						OperateTime:        timestamppb.Now(),
						ControlCommandType: cp.ControlCommandType,
						Delay:              cp.Delay,
						DigitalValue:       cp.DigitalValue,
						AnalogValue:        cp.AnalogValue,
						StringValue:        []byte(cp.StringValue),
					},
					RecoverParam: &pb.BottomCommandParam{},
				})
				handleCommandReport(&pb.ControlCommandReport{
					ControlCommandSerialNumber:       []byte(cp.ControlCommandSerialNumber),
					ControlCommandReportOccurredTime: timestamppb.Now(),
					ControlCommandResultType:         littlebee.ControlCommandResultEnum_CONTROL_COMMAND_RESULT_DELIVERING,
					ControlCommandExecuteDescription: []byte{},
				})
				return
			}
			if fsutype.DataAnalysisMode == int32(pb.FsuDataParsingTypeEnum_FSU_DATA_PARSING_TYPE_CENTER) {
				r.RealdataGrpcServer.SendCenterControlCommandRequest(&pb.CenterControlCommandRequest{
					OperatorUserId:             []byte(strconv.Itoa(int(cp.OperatorUserId))),
					OperatorUserName:           []byte(cp.OperatorUserName),
					ControlCommandSerialNumber: []byte(cp.ControlCommandSerialNumber),
					ControlParam: &pb.CenterCommandParam{
						DriverId:           sa.DriverId,
						X:                  sa.X,
						Y:                  sa.Y,
						Z:                  sa.Z,
						OperateTime:        timestamppb.Now(),
						ControlCommandType: cp.ControlCommandType,
						Delay:              cp.Delay,
						DigitalValue:       cp.DigitalValue,
						AnalogValue:        cp.AnalogValue,
						StringValue:        []byte(cp.StringValue),
					},
					RecoverParam: &pb.CenterCommandParam{},
				})
				handleCommandReport(&pb.ControlCommandReport{
					ControlCommandSerialNumber:       []byte(cp.ControlCommandSerialNumber),
					ControlCommandReportOccurredTime: timestamppb.Now(),
					ControlCommandResultType:         littlebee.ControlCommandResultEnum_CONTROL_COMMAND_RESULT_DELIVERING,
					ControlCommandExecuteDescription: []byte{},
				})
				return
			}
		}
	})
	return err == nil
}

func (r *RealdataServer) InitSignalStatus() {
	for k, v := range r.MapAcquisition {
		if v.AgentId == yaml.Yaml.Agent.AgentId {
			s := &pb.SignalCurrentStata{
				SignalId:              v.SignalId,
				SignalValueState:      pb.SignalValueStateEnum_SIGNAL_VALUE_STATE_UNKNOWN,
				SignalRunState:        pb.SignalRunStateEnum_SIGNAL_RUN_STATE_NORMAL,
				AlarmThreshholdNumber: -1,
				AlarmBeginTime:        timestamppb.Now(),
				AlarmLevelNumber:      -1,
			}
			r.MapSignalCurrentStatus[k] = s
		}
	}

	for _, v := range r.MapActiveAlarm {
		a := r.MapSignalCurrentStatus[v.SignalId]
		if a != nil {
			a.SignalRunState = pb.SignalRunStateEnum_SIGNAL_RUN_STATE_ALARM
			a.AlarmThreshholdNumber = v.AlarmThreshholdNumber
			a.AlarmLevelNumber = v.AlarmLevelNumber
			a.AlarmBeginTime = v.AlarmBeginTime
		} else {
			log.Logger.Error("not found alarm signal,end this alarm.", zap.Int32("signal id", v.SignalId), zap.String("alarm sn", v.AlarmSerialNumber))
			v.AlarmEndTime = timestamppb.Now()
			service.EndAlarm(v)
		}
	}
}

func (r *RealdataServer) LoadSignal() bool {
	logicObj, err := service.GetAllLogicObjects()
	if err != nil {
		log.Logger.Error("load all logicobject error", zap.String("error", err.Error()))
		return false
	}
	for _, v := range logicObj.GetItems() {
		r.MapLogicObjs[v.LogicObjectId] = v
	}

	signal, err1 := service.GetAllSignals()
	if err1 != nil {
		log.Logger.Error("load all signals error", zap.String("error", err1.Error()))
		return false
	}
	for _, v := range signal.Items {
		r.MapSignals[v.SignalId] = v
	}

	ai, err2 := service.GetAllSignalAis()
	if err2 != nil {
		log.Logger.Error("load all ai error", zap.String("error", err2.Error()))
		return false
	}
	for _, v := range ai.Items {
		r.MapSignalAi[v.SignalId] = v
	}

	di, err3 := service.GetAllSignalDis()
	if err3 != nil {
		log.Logger.Error("load all di error", zap.String("error", err3.Error()))
		return false
	}
	for _, v := range di.Items {
		r.MapSignalDi[v.SignalId] = v
	}

	si, err4 := service.GetAllSignalSis()
	if err4 != nil {
		log.Logger.Error("load all si error", zap.String("error", err4.Error()))
		return false
	}
	for _, v := range si.Items {
		r.MapSignalSi[v.SignalId] = v
	}

	threshold, err5 := service.GetAllSignalThreshholds()
	if err5 != nil {
		log.Logger.Error("load all threshold error", zap.String("error", err5.Error()))
		return false
	}
	for _, v := range threshold.Items {
		r.MapThreshold[v.ThresholdId] = v
	}

	acquisition, err6 := service.GetAllSignalAcquisitions()
	if err6 != nil {
		log.Logger.Error("load all acquisition error", zap.String("error", err6.Error()))
		return false
	}
	for _, v := range acquisition.Items {
		r.MapAcquisition[v.SignalId] = v
	}

	alarm, err7 := service.GetAllAlarms()
	if err7 != nil {
		log.Logger.Error("load all acquisition error", zap.String("error", err7.Error()))
		return false
	}
	for _, v := range alarm.Items {
		if v.AlarmEndTime == nil {
			r.MapActiveAlarm[v.SignalId] = v
		}
	}

	return true
}

func (r *RealdataServer) InitGrpc() {
	realdataGrpcServer := &rs.RealdataGrpcServer{
		Lis:                            nil,
		Server:                         nil,
		RestartChan:                    make(chan interface{}),
		HandleRawValue:                 handleRawValue,
		HandleChannelValue:             handleChannelValue,
		HandleFSUStatus:                handleFSUStatus,
		HandleCommandReport:            handleCommandReport,
		BottomCommandRequestStream:     nil,
		CenterCommandStreamStream:      nil,
		BottomCommandRequestStreamLock: &sync.Mutex{},
		CenterCommandStreamStreamLock:  &sync.Mutex{},
	}
	realdataGrpcServer.PrepareRestartServer()
	realdataGrpcServer.StartServer()

	r.RealdataGrpcServer = realdataGrpcServer
}

func (r *RealdataServer) StopServer() {
	r.RealdataGrpcServer.StopServer()
}

func (r *RealdataServer) getThresholdByNumber(signalId int32, number int32) *pb.SignalThreshold {
	for _, v := range r.MapThreshold {
		if v.SignalId == signalId && v.ThresholdNumber == number {
			return v
		}
	}
	return nil
}

func (r *RealdataServer) getThresholdBySignalId(signalId int32) map[int32]*pb.SignalThreshold {
	m := make(map[int32]*pb.SignalThreshold)
	for _, v := range r.MapThreshold {
		if v.SignalId == signalId {
			m[v.ThresholdNumber] = v
		}
	}
	return m
}

func (r *RealdataServer) convertEndAlarmToAlarm(b *pb.EndAlarm) *pb.Alarm {
	return &pb.Alarm{
		AlarmSerialNumber: b.AlarmSerialNumber,
		SignalId:          b.SignalId,
		AlarmBeginTime:    b.AlarmBeginTime,
		AlarmEndTime:      b.AlarmEndTime,
		AlarmDesc:         b.AlarmDesc,
	}
}

func (r *RealdataServer) convertBeginAlarmToAlarm(b *pb.BeginAlarm) *pb.Alarm {
	return &pb.Alarm{
		AlarmSerialNumber:     b.AlarmSerialNumber,
		SignalId:              b.SignalId,
		SignalName:            b.SignalName,
		SignalType:            b.SignalType,
		AlarmBeginTime:        b.AlarmBeginTime,
		DeviceLocation:        b.DeviceLocation,
		DeviceName:            b.DeviceName,
		DeviceTypeId:          b.DeviceTypeId,
		AlarmDesc:             b.AlarmDesc,
		AlarmLevelNumber:      b.AlarmLevelNumber,
		AlarmThreshholdNumber: b.AlarmThreshholdNumber,
		SignalValueType:       b.SignalValueType,
		OccurredDigitalValue:  b.OccurredDigitalValue,
		OccurredAnalogValue:   b.OccurredAnalogValue,
		ValueUnit:             b.ValueUnit,
		OccurredStringValue:   b.OccurredStringValue,
	}
}

func (r *RealdataServer) generateEndAlarm(rawevent *pb.SignalRawEvent) *pb.EndAlarm {
	return &pb.EndAlarm{
		AlarmSerialNumber: r.generateAlarmSerialNumber(rawevent.SignalId, rawevent.AlarmBeginTime),
		SignalId:          rawevent.SignalId,
		AlarmBeginTime:    rawevent.AlarmBeginTime,
		AlarmEndTime:      rawevent.EventOccurredTime,
		AlarmDesc:         rawevent.EventDesc,
	}
}

func (r *RealdataServer) generateBeginAlarm(rawevent *pb.SignalRawEvent) *pb.BeginAlarm {
	return &pb.BeginAlarm{
		AlarmSerialNumber:     r.generateAlarmSerialNumber(rawevent.SignalId, rawevent.AlarmBeginTime),
		SignalId:              rawevent.SignalId,
		SignalName:            r.getSignalFullName(rawevent.SignalId),
		SignalType:            r.MapSignals[rawevent.SignalId].SignalType,
		AlarmBeginTime:        rawevent.EventOccurredTime,
		DeviceLocation:        r.getDeviceLocationBySignalId(rawevent.SignalId),
		DeviceName:            r.getDeviceName(rawevent.SignalId),
		DeviceTypeId:          r.getDeviceType(rawevent.SignalId),
		AlarmDesc:             rawevent.EventDesc,
		AlarmLevelNumber:      rawevent.AlarmLevelNumber,
		AlarmThreshholdNumber: rawevent.AlarmThreshholdNumber,
		SignalValueType:       rawevent.SignalValueType,
		OccurredDigitalValue:  rawevent.DigitalValue,
		OccurredAnalogValue:   rawevent.AnalogValue,
		ValueUnit:             rawevent.ValueUnit,
		OccurredStringValue:   rawevent.StringValue,
	}
}

func (r *RealdataServer) getSignalFullName(signalId int32) string {
	s := r.MapSignals[signalId]
	if s == nil {
		return ""
	}
	fullname := s.SignalName
	terminateLogicObjId := r.getDeviceIdBySignalId(signalId)
	if terminateLogicObjId == -1 {
		l := r.MapLogicObjs[s.ParentLogicId]
		if l == nil {
			return fullname
		} else {
			return l.LogicObjectName + fullname
		}
	}
	var parLogicObjId int32 = s.ParentLogicId
	for {
		if parLogicObjId == terminateLogicObjId {
			return fullname
		}
		l := r.MapLogicObjs[parLogicObjId]
		if l == nil {
			return fullname
		} else {
			fullname = l.LogicObjectName + fullname
			parLogicObjId = l.ParentLogicObjectId
		}
	}
}

func (r *RealdataServer) getDeviceLocationBySignalId(signalId int32) string {
	location := ""
	var parLogicObjId int32 = 0
	deviceId := r.getDeviceIdBySignalId(signalId)
	if deviceId == -1 {
		s := r.MapSignals[signalId]
		if s == nil {
			return location
		}
		parLogicObjId = s.ParentLogicId
	} else {
		parLogicObjId = deviceId
	}

	for {
		if parLogicObjId == 0 {
			return location
		}
		l := r.MapLogicObjs[parLogicObjId]
		if l == nil {
			return location
		} else {
			location = l.LogicObjectName + "." + location
			parLogicObjId = l.ParentLogicObjectId
		}
	}
}

func (r *RealdataServer) getDeviceType(signalId int32) int32 {
	deviceId := r.getDeviceIdBySignalId(signalId)
	if deviceId == -1 {
		s := r.MapSignals[signalId]
		if s == nil {
			return -1
		}
		l := r.MapLogicObjs[s.ParentLogicId]
		if l == nil {
			return -1
		}
		dt, err := service.RetrieveLogicDevice(l.LogicObjectId)
		if err == nil {
			return dt.DeviceType
		} else {
			return -1
		}
	}
	d := r.MapLogicObjs[deviceId]
	if d == nil {
		return -1
	} else {
		dt, err := service.RetrieveLogicDevice(d.LogicObjectId)
		if err == nil {
			return dt.DeviceType
		} else {
			return -1
		}
	}
}

func (r *RealdataServer) getDeviceName(signalId int32) string {
	deviceId := r.getDeviceIdBySignalId(signalId)
	if deviceId == -1 {
		s := r.MapSignals[signalId]
		if s == nil {
			return ""
		}
		l := r.MapLogicObjs[s.ParentLogicId]
		if l == nil {
			return ""
		}
		return l.LogicObjectName
	}
	d := r.MapLogicObjs[deviceId]
	if d == nil {
		return ""
	} else {
		return d.LogicObjectName
	}
}

func (r *RealdataServer) getDeviceIdBySignalId(signalId int32) int32 {
	s := r.MapSignals[signalId]
	if s == nil {
		return -1
	}
	parLogicObjId := s.ParentLogicId
	for {
		if parLogicObjId == 0 {
			return -1
		}
		l := r.MapLogicObjs[parLogicObjId]
		if l == nil {
			return -1
		}
		if l.LogicObjectTypeId == int32(pb.SignalTreeObjectTypeEnum_SIGNAL_TREE_OBJECT_TYPE_DEVICE) {
			return l.LogicObjectId
		} else {
			parLogicObjId = l.ParentLogicObjectId
		}
	}
}
