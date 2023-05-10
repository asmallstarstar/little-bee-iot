package fsustatus

import (
	dao "dao/config"
	"message/littlebee"
	c "service/consumer"

	"google.golang.org/protobuf/types/known/timestamppb"
)

func ResponseFsuStatus(userId int32, r *littlebee.FsuStatusRequest) *littlebee.FSUStatusList {
	list := &littlebee.FSUStatusList{
		Items: make([]*littlebee.FSUStatus, 0),
	}
	for _, v := range r.FsuIds {
		fs := c.MapFsuStatus[v]
		if fs != nil {
			list.Items = append(list.Items, &littlebee.FSUStatus{
				FsuId:            fs.FsuId,
				FsuName:          fs.FsuName,
				FsuType:          fs.FsuType,
				SubChannelName:   fs.SubChannelName,
				LastReceiveTime:  fs.LastReceiveTime,
				FsuStatus:        fs.FsuStatus,
				SubChannelStatus: fs.SubChannelStatus,
				StatusDesc:       fs.StatusDesc,
			})
		} else {
			s, e := dao.RetrieveFsu(v)
			if e != nil {
				list.Items = append(list.Items, &littlebee.FSUStatus{
					FsuId: v,
				})
			} else {
				list.Items = append(list.Items, &littlebee.FSUStatus{
					FsuId:            v,
					FsuName:          s.FsuName,
					FsuType:          "",
					SubChannelName:   "",
					LastReceiveTime:  timestamppb.Now().AsTime().Unix(),
					FsuStatus:        int32(littlebee.FsuRunStateEnum_FSU_RUN_STATE_UNKNOWN),
					SubChannelStatus: int32(littlebee.FsuRunStateEnum_FSU_RUN_STATE_UNKNOWN),
					StatusDesc:       "",
				})
			}
		}

	}
	return list
}
