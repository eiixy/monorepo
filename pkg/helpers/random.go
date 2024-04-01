package helpers

import (
	"math/rand"
	"time"
)

type rule uint8

const (
	Alpha     rule = 1 // 由字母组成
	AlphaNum  rule = 2 // 只能由字母和数字组成
	AlphaDash rule = 3 // 包含字母、数字，短破折号（-）和下划线（_）
)

// RandomNumber 生成随机数据
func RandomNumber(start, end int64) int64 {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return start + r.Int63n(end-start)
}

// RandomAlphaStr 生成随机字母
func RandomAlphaStr(length int) string {
	return randomStr(Alpha, length)
}

// RandomAlphaNumStr 生成随机字母和数字
func RandomAlphaNumStr(length int) string {
	return randomStr(AlphaNum, length)
}

// RandomAlphaDashStr 生成随机字母、数字，短破折号（-）和下划线（_）
func RandomAlphaDashStr(length int) string {
	return randomStr(AlphaDash, length)
}

func randomStr(r rule, l int) string {
	rules := map[rule]string{
		Alpha:     "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz",
		AlphaNum:  "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789012345678901234567890123456789",
		AlphaDash: "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789012345678901234567890123456789-_",
	}
	s := rules[r]
	result := ""
	ra := rand.New(rand.NewSource(time.Now().UnixNano()))

	for i := 0; i < l; i++ {
		n := ra.Intn(len(s))
		result += string(s[n])
	}
	return result
}
