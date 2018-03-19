package ex02

import (
	"bytes"
	"fmt"
)

type MapIntSet struct {
	set map[int]bool
}

func NewMapIntSet() *MapIntSet {
	return &MapIntSet{
		set: map[int]bool{},
	}
}

func (m *MapIntSet) Has(x int) bool {
	k, ok := m.set[x]
	return k && ok
}

func (m *MapIntSet) Add(x int) {
	m.set[x] = true
}

func (m *MapIntSet) UnionWith(t *MapIntSet) {
	for k, v := range t.set {
		if v {
			m.set[k] = v
		}
	}
}

func (m *MapIntSet) String() string {
	var buf bytes.Buffer
	buf.WriteByte('{')
	var cnt = 0
	for k := range m.set {
		fmt.Fprintf(&buf, "%d", k)
		cnt++
		if cnt < len(m.set) {
			buf.WriteByte(' ')
		}
	}
	buf.WriteByte('}')

	return buf.String()
}
