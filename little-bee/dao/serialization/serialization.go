package serialization

import (
	"context"
	"reflect"
	"time"

	"github.com/golang/protobuf/ptypes/timestamp"
	"google.golang.org/protobuf/types/known/timestamppb"
	"gorm.io/gorm/schema"
)

type PBTimestampSerializer struct {
}

// UTC time.Time --> timestamp.Timestamp
// Scan implements serializer interface
func (PBTimestampSerializer) Scan(ctx context.Context, field *schema.Field, dst reflect.Value, dbValue interface{}) (err error) {
	if dbValue == nil {
		field.ReflectValueOf(ctx, dst).Set(reflect.ValueOf((*timestamppb.Timestamp)(nil)))
	} else {
		ts := timestamppb.New(dbValue.(time.Time))
		field.ReflectValueOf(ctx, dst).Set(reflect.ValueOf(ts))
	}
	return
}

// timestamp.Timestamp -->UTC time.Time
// Value implements serializer interface
func (PBTimestampSerializer) Value(ctx context.Context, field *schema.Field, dst reflect.Value, fieldValue interface{}) (interface{}, error) {
	ts := fieldValue.(*timestamp.Timestamp)
	err := ts.CheckValid()
	unixtime := ts.AsTime()
	s := unixtime.UTC()
	return s, err
}
