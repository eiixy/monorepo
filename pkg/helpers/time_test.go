package helpers

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestTimeSubDuration(t *testing.T) {
	type args struct {
		t time.Time
		d time.Duration
	}
	now := time.Now()
	tests := []struct {
		name string
		args args
		want time.Time
	}{
		{"", args{now, 24 * time.Hour}, time.UnixMilli(now.UnixMilli() - 24*3600*1000)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, TimeSubDuration(tt.args.t, tt.args.d), "TimeSubDuration(%v, %v)", tt.args.t, tt.args.d)
		})
	}
}

func TestNowSubDuration(t *testing.T) {
	type args struct {
		d time.Duration
	}
	now := time.Now()
	tests := []struct {
		name string
		args args
		want time.Time
	}{
		{"", args{5 * time.Hour}, time.UnixMilli(now.UnixMilli() - 5*3600*1000)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, NowSubDuration(tt.args.d), "NowSubDuration(%v)", tt.args.d)
		})
	}
}
