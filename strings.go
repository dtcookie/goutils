package goutils

import "strings"

type Strings []string

func (s Strings) Contains(v string) bool {
	for _, elem := range s {
		if elem == v {
			return true
		}
	}
	return false
}

func (s Strings) ContainsPrefixFor(v string) bool {
	for _, prefix := range s {
		if strings.HasPrefix(v, prefix) {
			return true
		}
	}
	return false
}

func (s Strings) GetPrefixFor(v string) string {
	for _, prefix := range s {
		if strings.HasPrefix(v, prefix) {
			return prefix
		}
	}
	return ""
}

func ContainsAnyPrefix(s string, prefixes Strings) bool {
	for _, prefix := range prefixes {
		if strings.HasPrefix(s, prefix) {
			return true
		}
	}
	return false

}

func (s Strings) ContainsAny(v Strings) bool {
	if v == nil {
		return false
	}
	for _, elem := range v {
		if s.Contains(elem) {
			return true
		}
	}
	return false
}
