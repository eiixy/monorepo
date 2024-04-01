package helpers

import (
	"github.com/stretchr/testify/assert"
	"google.golang.org/protobuf/types/known/wrapperspb"
	"testing"
)

func TestGetWrapUInt(t *testing.T) {
	// Positive test case
	v := wrapperspb.UInt32Value{Value: 10}
	expected := uint(10)
	assert.Equal(t, expected, *GetWrapUInt(&v))
	// Negative test case
	assert.Nil(t, GetWrapUInt(nil))
}
func TestGetWrapUInt8(t *testing.T) {
	// Positive test case
	v := wrapperspb.UInt32Value{Value: 8}
	expected := uint8(8)
	assert.Equal(t, expected, *GetWrapUInt8(&v))
	// Negative test case
	assert.Nil(t, GetWrapUInt8(nil))
}
func TestGetWrapString(t *testing.T) {
	// Positive test case
	v := wrapperspb.StringValue{Value: "test"}
	expected := "test"
	assert.Equal(t, expected, *GetWrapString(&v))
	// Negative test case
	assert.Nil(t, GetWrapString(nil))
}
