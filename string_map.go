package goutils

type StringMap map[string]string

func (m StringMap) Consume(other map[string]string) {
	if other == nil {
		return
	}
	for k, v := range other {
		m[k] = v
	}
}
