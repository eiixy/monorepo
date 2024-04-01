package helpers

import (
	"path"
	"runtime"
	"strings"
)

func Path(s ...string) string {
	_, filename, _, _ := runtime.Caller(0)
	prefix := []string{strings.ReplaceAll(filename, "/pkg/helpers/path.go", "")}
	if len(s) > 0 && strings.HasPrefix(s[0], "/") {
		prefix = []string{}
	}
	s = append(prefix, s...)
	return path.Join(s...)
}
