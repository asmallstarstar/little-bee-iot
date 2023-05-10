package realdata

import (
	dao "dao/config"
	"message/littlebee"
	c "service/consumer"

	"google.golang.org/protobuf/types/known/timestamppb"
)

func ResponseRealdata(userId int32, r *littlebee.RealdataRequest) *littlebee.SignalRealValueList {
	list := &littlebee.SignalRealValueList{
		Items: make([]*littlebee.SignalRealValue, 0),
	}
	for _, v := range r.SignalIds {
		sv := c.MapRealdata[v]
		if sv != nil {
			list.Items = append(list.Items, &littlebee.SignalRealValue{
				SignalId:              v,
				SignalUnificationId:   sv.SignalUnificationId,
				SignalName:            sv.SignalName,
				SignalType:            sv.SignalType,
				SignalValueState:      sv.SignalValueState,
				SignalRunState:        sv.SignalRunState,
				AlarmThreshholdNumber: sv.AlarmThreshholdNumber,
				AlarmLevelNumber:      sv.AlarmLevelNumber,
				ValueOccurredTime:     sv.ValueOccurredTime,
				SignalValueType:       sv.SignalValueType,
				DigitalValue:          sv.DigitalValue,
				AnalogValue:           sv.AnalogValue,
				ValueUnit:             sv.ValueUnit,
				DecimalPrecision:      sv.DecimalPrecision,
				StringValue:           sv.StringValue,
				DigitalValueDesc:      sv.DigitalValueDesc,
			})
		} else {
			s, e := dao.RetrieveSignal(v)
			if e != nil {
				list.Items = append(list.Items, &littlebee.SignalRealValue{
					SignalId: v,
				})
			} else {
				list.Items = append(list.Items, &littlebee.SignalRealValue{
					SignalId:              v,
					SignalUnificationId:   s.SignalUnificationId,
					SignalName:            s.SignalName,
					SignalType:            s.SignalType,
					SignalValueState:      littlebee.SignalValueStateEnum_SIGNAL_VALUE_STATE_UNKNOWN,
					SignalRunState:        littlebee.SignalRunStateEnum_SIGNAL_RUN_STATE_UNKNOWN,
					AlarmThreshholdNumber: 0,
					AlarmLevelNumber:      0,
					ValueOccurredTime:     timestamppb.Now(),
					SignalValueType:       0,
					DigitalValue:          0,
					AnalogValue:           0.0,
					ValueUnit:             "",
					DecimalPrecision:      0,
					StringValue:           "",
					DigitalValueDesc:      "",
				})
			}
		}
	}
	return list
}
