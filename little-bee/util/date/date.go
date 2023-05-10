package date

import (
	"time"

	"github.com/golang/protobuf/ptypes"
	"github.com/golang/protobuf/ptypes/timestamp"
)

func FormatProtobufTime(time *timestamp.Timestamp) string {
	timestamp, _ := ptypes.Timestamp(time)
	return timestamp.Format("2006-01-02 15:04:05")
}

func ConvertProtobufTime(time *timestamp.Timestamp) time.Time {
	timestamp, _ := ptypes.Timestamp(time)
	return timestamp
}
