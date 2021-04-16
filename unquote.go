package goutils

import "strings"

func Unquote(s string) string {
	if len(s) < 2 {
		return s
	}
	if !strings.HasPrefix(s, "'") {
		return s
	}
	if !strings.HasSuffix(s, "'") {
		return s
	}
	return s[1 : len(s)-1]
}
