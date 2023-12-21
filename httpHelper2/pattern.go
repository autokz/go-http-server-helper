package httpHelper2

import (
	"strings"
)

func transform(pattern string) string {
	pattern = strings.TrimSpace(pattern)
	if pattern == "" || pattern == "/" {
		return "/"
	}
	if pattern[0] != '/' {
		pattern = "/" + pattern
	}
	if strings.Contains(pattern, " ") {
		ErrorLog.Panicf("invalid pattern \"%s\"", pattern)
	}
	return pattern
}
