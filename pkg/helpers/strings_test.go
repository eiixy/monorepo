package helpers

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSubStr(t *testing.T) {
	type args struct {
		s      string
		start  int
		length int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"", args{"0123456789", 0, 5}, "01234"},
		{"", args{"0123456789", 2, 5}, "23456"},
		{"", args{"0123456789", 3, 15}, "3456789"},
		{"", args{"0123456789", 13, 5}, ""},
		{"", args{"0123456789测试字符串", 8, 5}, "89测试字"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, SubStr(tt.args.s, tt.args.start, tt.args.length), "SubStr(%v, %v, %v)", tt.args.s, tt.args.start, tt.args.length)
		})
	}
}

func TestSnake(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"", args{"AccountService"}, "account_service"},
		{"", args{"Account_Service"}, "account_service"},
		{"", args{"Account___Service"}, "account_service"},
		{"", args{" Account-Service "}, "account_service"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, Snake(tt.args.s), "Snake(%v)", tt.args.s)
		})
	}
}

func TestStudly(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"", args{"account service"}, "AccountService"},
		{"", args{"account  service"}, "AccountService"},
		{"", args{"account-service"}, "AccountService"},
		{"", args{"account--service"}, "AccountService"},
		{"", args{"account_service"}, "AccountService"},
		{"", args{"account__service"}, "AccountService"},
		{"", args{" account_service "}, "AccountService"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, Studly(tt.args.s), "Studly(%v)", tt.args.s)
		})
	}
}

func TestSeparate(t *testing.T) {
	type args struct {
		content string
		sep     string
		chunk   int
	}
	content := "0123456789abcdef"
	tests := []struct {
		name string
		args args
		want string
	}{
		{"", args{content, "-", 2}, "01-23-45-67-89-ab-cd-ef"},
		{"", args{content, "_", 3}, "012_345_678_9ab_cde_f"},
		{"", args{content, ",", 5}, "01234,56789,abcde,f"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, Separate(tt.args.content, tt.args.sep, tt.args.chunk), "Separate(%v, %v, %v)", tt.args.content, tt.args.sep, tt.args.chunk)
		})
	}
}

func TestGetBetweenStr(t *testing.T) {
	type args struct {
		str   string
		start string
		end   string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"", args{"abababababbb", "ab", "b"}, "a"},
		{"", args{
			str:   "><meta property=\"og:site_name\" content=\"YouTube\"><meta property=\"og:url\" content=\"https://www.youtube.com/channel/UCZp5GBcPLFvzcbja_J5NdPw\"><meta property=\"og:image\" content=\"https://yt3.googleusercontent.com/ytc/AOPolaQbDTbzIqSnvsVcUNyGgi3YfSf_nnSXB1AIaZQ_FQ=s900-c-k-c0x00ffffff-no-rj\">",
			start: "<meta property=\"og:url\" content=\"",
			end:   "\">",
		}, "https://www.youtube.com/channel/UCZp5GBcPLFvzcbja_J5NdPw"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, GetBetweenStr(tt.args.str, tt.args.start, tt.args.end), "GetBetweenStr(%v, %v, %v)", tt.args.str, tt.args.start, tt.args.end)
		})
	}
}
