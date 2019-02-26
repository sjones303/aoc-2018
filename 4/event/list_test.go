package event

import (
	"testing"
)

func TestListLen(t *testing.T) {
	for _, td := range []struct {
		name string
		in   []Event
		exp  int
	}{
		{
			name: "Nil",
			in:   nil,
			exp:  0,
		},
	} {
		t.Run(td.name, func(t *testing.T) {
			l := List(td.in)
			if act := l.Len(); act != td.exp {
				t.Errorf("Len = %d, want %d", act, td.exp)
			}
		})
	}
}
