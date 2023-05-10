package serialization

import (
	"github.com/golang/protobuf/jsonpb"
	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/ptypes/any"
	"google.golang.org/protobuf/types/known/anypb"
)

func MarshalToJson(m proto.Message) (string, error) {
	marshaler := jsonpb.Marshaler{
		EnumsAsInts:  true,
		OrigName:     false,
		EmitDefaults: true,
	}
	return marshaler.MarshalToString(m)
}

func UnmarshalFromJson(json string, m proto.Message) error {
	err := jsonpb.UnmarshalString(json, m)
	return err
}

func AnyToStruct(any *any.Any, m proto.Message) error {
	err := any.UnmarshalTo(proto.MessageV2(m))
	return err
}

func StructToAny(m proto.Message) *any.Any {
	any, _ := anypb.New(proto.MessageV2(m))
	return any
}
