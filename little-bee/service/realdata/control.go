package realdata

import (
	dao "dao/config"
	"message/littlebee"
	"service/consumer"
	"strconv"

	"github.com/google/uuid"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func Control(userId int32, r *littlebee.CommandParam) *littlebee.ControlCommandReport {
	uuid := uuid.New().String()
	r.ControlCommandSerialNumber = uuid
	now := timestamppb.Now()
	m, _ := proto.Marshal(r)
	if err := consumer.PublishControl(m); err != nil {
		return &littlebee.ControlCommandReport{
			ControlCommandSerialNumber:       []byte(uuid),
			ControlCommandReportOccurredTime: now,
			ControlCommandResultType:         littlebee.ControlCommandResultEnum_CONTROL_COMMAND_RESULT_FAIL,
			ControlCommandExecuteDescription: []byte{},
		}
	}

	record := &littlebee.ControlCommandRecord{
		ControlCommandSerialNumber: uuid,
		OperatorUserId:             strconv.Itoa(int(userId)),
		OperatorUserName:           string(r.OperatorUserName),
		SignalId:                   r.SignalId,
		OperateTime:                now,
		ControlCommandType:         r.ControlCommandType,
		ControlCommandResult:       littlebee.ControlCommandResultEnum_CONTROL_COMMAND_RESULT_RECEIVED,
		Delay:                      r.Delay,
		DigitalValue:               r.DigitalValue,
		AnalogValue:                r.AnalogValue,
		StringValue:                string(r.StringValue),
	}
	dao.CreateControlCommandRecord(record)

	return &littlebee.ControlCommandReport{
		ControlCommandSerialNumber:       []byte(uuid),
		ControlCommandReportOccurredTime: now,
		ControlCommandResultType:         littlebee.ControlCommandResultEnum_CONTROL_COMMAND_RESULT_RECEIVED,
		ControlCommandExecuteDescription: []byte{},
	}
}
