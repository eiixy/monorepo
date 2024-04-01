package helpers

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRound(t *testing.T) {
	type args struct {
		val float64
		len int
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{"", args{1.2, 3}, 1.2},
		{"", args{12.444444444, 0}, 12},
		{"", args{12.444444444, 1}, 12.4},
		{"", args{12.444444444, 2}, 12.44},
		{"", args{12.444444444, 3}, 12.444},
		{"", args{12.555555555, 0}, 13},
		{"", args{12.555555555, 1}, 12.6},
		{"", args{12.555555555, 2}, 12.56},
		{"", args{12.555555555, 3}, 12.556},
		{"", args{-12.444444444, 0}, -12},
		{"", args{-12.444444444, 1}, -12.4},
		{"", args{-12.444444444, 2}, -12.44},
		{"", args{-12.444444444, 3}, -12.444},
		{"", args{-12.555555555, 0}, -13},
		{"", args{-12.555555555, 1}, -12.6},
		{"", args{-12.555555555, 2}, -12.56},
		{"", args{-12.555555555, 3}, -12.556},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, Round(tt.args.val, tt.args.len), "Round(%v, %v)", tt.args.val, tt.args.len)
		})
	}
}

func TestReverse(t *testing.T) {
	type args struct {
		x int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"", args{0}, 0},
		{"", args{102}, 201},
		{"", args{-102}, -201},
		{"", args{100}, 1},
		{"", args{8}, 8},
		{"", args{200100}, 1002},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, Reverse(tt.args.x), "Reverse(%v)", tt.args.x)
		})
	}
}

func TestFormatNumberRange(t *testing.T) {
	testCases := []struct {
		input    int
		expected string
	}{
		{500, "<1K"},
		{10000, "1K-10K"},
		{50000, "10K-100K"},
		{200000, "100K-500K"},
		{800000, "500K-1M"},
		{3000000, "1M-5M"},
		{8000000, "5M-10M"},
		{15000000, "10M+"},
	}

	for _, tc := range testCases {
		result := FormatNumberRange(tc.input)
		if result != tc.expected {
			t.Errorf("Expected %s for input %d, but got %s", tc.expected, tc.input, result)
		}
	}
}

func TestFormatNumber(t *testing.T) {
	testCases := []struct {
		input    int
		expected string
	}{
		{500, "0.50K"},
		{10000, "10.00K"},
		{50000, "50.00K"},
		{200000, "200.00K"},
		{800000, "800.00K"},
		{3000000, "3.00M"},
		{8000000, "8.00M"},
		{15000000, "15.00M"},
	}

	for _, tc := range testCases {
		result := FormatNumber(tc.input)
		if result != tc.expected {
			t.Errorf("Expected %s for input %d, but got %s", tc.expected, tc.input, result)
		}
	}
}
