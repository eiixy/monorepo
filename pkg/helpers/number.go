package helpers

import (
	"fmt"
	"math"
)

func Fib(n int) int {
	sqrt5 := math.Sqrt(5)
	p1 := math.Pow((1+sqrt5)/2, float64(n))
	p2 := math.Pow((1-sqrt5)/2, float64(n))
	return int(math.Round((p1 - p2) / sqrt5))
}

// Round 四舍五入，保留 n 位小数点
func Round(val float64, len int) float64 {
	x := math.Pow(10, float64(len))
	return math.Round(val*x) / x
}

// Reverse 对数字进行反转
func Reverse(x int) int {
	var temp int
	for x != 0 {
		remainder := x % 10
		x /= 10
		temp = temp*10 + remainder
	}
	return temp
}

func FormatNumberRange(v int) string {
	return CustomFormatNumberRange(v, "10M+", []Range{
		{1000, "0-1K"},
		{5000, "1K-5K"},
		{10000, "5K-10K"},
		{100000, "10K-100K"},
		{500000, "100K-500K"},
		{1000000, "500K-1M"},
		{10000000, "1M-10M"},
	})
}

func CustomFormatNumberRange(v int, defaultValue string, ranges []Range) string {
	for _, item := range ranges {
		if v <= item.Max {
			return item.Format
		}
	}
	return defaultValue
}

type Range struct {
	Max    int
	Format string
}

func FormatNumber(v int) string {
	if v > 1000000 {
		return fmt.Sprintf("%.2fM", float64(v)/1000000)
	}
	return fmt.Sprintf("%.2fK", float64(v)/1000)
}
