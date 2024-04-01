package helpers

import (
	"reflect"
	"strings"
	"time"
)

func If[T any](condition bool, a, b T) T {
	if condition {
		return a
	}
	return b
}

func Get[T comparable](value, defaultValue T) T {
	var t T
	if value == t {
		return defaultValue
	}
	return value
}

func Empty[T comparable](v T) bool {
	var t T
	return v == t
}

func HideEmail(email string) string {
	atIndex := strings.Index(email, "@")
	if atIndex == -1 {
		return "****"
	}

	if atIndex > 5 {
		return email[:2] + "****" + email[atIndex:]
	}
	if atIndex > 4 {
		return email[:2] + "***" + email[atIndex:]
	}
	if atIndex > 3 {
		return email[:2] + "**" + email[atIndex:]
	}
	if atIndex > 2 {
		return email[:1] + "**" + email[atIndex:]
	}

	return "**" + email[atIndex:]
}

func ToPoint[T comparable](v T, defaultValue ...*T) *T {
	if Empty(v) {
		if len(defaultValue) > 0 {
			return defaultValue[0]
		}
		return nil
	}
	return &v
}

func Point[T any](v T) *T {
	return &v
}

func GetPoint[T comparable](v *T, defaultVal ...T) T {
	if v == nil {
		if len(defaultVal) > 0 {
			return defaultVal[0]
		}
		var t T
		return t
	}
	return *v
}

// Throttle 节流函数
func Throttle(ttl time.Duration) func(fn func()) {
	var lastExecAt time.Time
	return func(fn func()) {
		if time.Now().After(lastExecAt.Add(ttl)) {
			fn()
			lastExecAt = time.Now()
		}
	}
}

func IsNil(v any) bool {
	if v == nil {
		return true
	}
	return reflect.ValueOf(v).IsNil()
}
