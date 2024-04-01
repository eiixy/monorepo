package helpers

import (
	"fmt"
	"time"
)

const Day = 24 * 3600000 // millisecond

func NowSubDuration(d time.Duration) time.Time {
	t := time.Now().UnixMilli()
	return time.UnixMilli(t - int64(d/time.Millisecond))
}

func NowAddDuration(d time.Duration) time.Time {
	t := time.Now().UnixMilli()
	return time.UnixMilli(t + int64(d/time.Millisecond))
}

func NowSubDays(d int) time.Time {
	t := time.Now().UnixMilli()
	return time.UnixMilli(t - int64(d*Day))
}

func ISOWeek(t time.Time) string {
	y, w := t.ISOWeek()
	return fmt.Sprintf("%d-%d", y, w)
}

// TimeSubDuration time sub duration
func TimeSubDuration(t time.Time, d time.Duration) time.Time {
	return time.UnixMilli(t.UnixMilli() - int64(d/time.Millisecond))
}

func Time2Point(t time.Time) *time.Time {
	if t.IsZero() {
		return nil
	}
	return &t
}
