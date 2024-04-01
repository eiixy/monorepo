package helpers

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMD5(t *testing.T) {
	type args struct {
		str string
	}
	author := "Pablo B."
	content := "It's the best Ebike out right now, fast reliable and so convenient. Replacement parts are easy to come by and it's a easy fix"
	tests := []struct {
		name string
		args args
		want string
	}{
		{"", args{author + content}, "7e8d3e4d9c7e9178065b272594a3e258"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, MD5(tt.args.str), "MD5(%v)", tt.args.str)
		})
	}
}
