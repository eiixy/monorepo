package helpers

import "google.golang.org/protobuf/types/known/wrapperspb"

type Wrapper[V any] interface {
	GetValue() V
}

func GetWrapUInt(v *wrapperspb.UInt32Value) *uint {
	if v == nil {
		return nil
	}
	var x = uint(v.GetValue())
	return &x
}

func GetWrapInt(v *wrapperspb.UInt32Value) *int {
	if v == nil {
		return nil
	}
	var x = int(v.GetValue())
	return &x
}

func GetWrapInt64(v *wrapperspb.UInt32Value) *int64 {
	if v == nil {
		return nil
	}
	var x = int64(v.GetValue())
	return &x
}

func GetWrapUInt8(v *wrapperspb.UInt32Value) *uint8 {
	if v == nil {
		return nil
	}
	var x = uint8(v.GetValue())
	return &x
}
func GetWrapString(v *wrapperspb.StringValue) *string {
	if v == nil {
		return nil
	}
	var x = v.GetValue()
	return &x
}
