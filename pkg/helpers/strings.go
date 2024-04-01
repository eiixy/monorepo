package helpers

import (
	"golang.org/x/exp/slices"
	"strings"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

// Studly 将字符串转换成驼峰命名的字符串
func Studly(s string) string {
	s = strings.ReplaceAll(s, "_", " ")
	s = strings.ReplaceAll(s, "-", " ")
	s = cases.Title(language.English).String(s)
	return strings.ReplaceAll(s, " ", "")
}

// Snake 将驼峰的函数名或者字符串转换成 _ 命名的函数或者字符串，例如 snakeCase 转换成 snake_case
func Snake(s string) string {
	s = strings.Trim(s, " ")
	data := make([]byte, 0, len(s)*2)
	num := len(s)
	for i := 0; i < num; i++ {
		c := s[i]
		if i > 0 && c >= 'A' && c <= 'Z' {
			data = append(data, '_')
		}
		if !slices.Contains([]byte{'-', '_'}, c) {
			data = append(data, c)
		}
	}
	return strings.ToLower(string(data))
}

// Separate 将字符串按长度做分割
func Separate(content, sep string, chunk int) string {
	var elems []string
	for i := 0; i < len(content); i = i + chunk {
		end := i + chunk
		if end > len(content) {
			end = len(content)
		}
		elems = append(elems, content[i:end])
	}
	return strings.Join(elems, sep)
}

func GetBetweenStr(str, start, end string) string {
	s := strings.Index(str, start) + len(start)
	e := strings.Index(str[s:], end)
	return str[s : s+e]
}

func SubStr(s string, start, length int) string {
	rs := []rune(s)
	if len(rs) > start+length {
		return string(rs[start : start+length])
	}
	if len(rs) <= start {
		return ""
	}
	return s[start:]
}
