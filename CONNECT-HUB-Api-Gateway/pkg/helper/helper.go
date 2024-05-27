package helper

import (
	"time"

	"google.golang.org/protobuf/types/known/timestamppb"
)

func FromProtoTimestamp(ts *timestamppb.Timestamp) time.Time {
	if ts == nil {
		return time.Time{}
	}
	return ts.AsTime()
}
