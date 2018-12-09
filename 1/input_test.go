package main

import (
	"fmt"
	"testing"
)

func TestParse(t *testing.T) {
	for i, td := range []struct {
		in  []string
		exp []int
	}{
		{
			[]string{"+1", "-2", "+3", "+1"},
			[]int{1, -2, 3, 1},
		},
		{
			[]string{"+1", "+1", "+1"},
			[]int{1, 1, 1},
		},
		{
			[]string{"+1", "+1", "-2"},
			[]int{1, 1, -2},
		},
		{
			[]string{"-1", "-2", "-3"},
			[]int{-1, -2, -3},
		},
	} {
		t.Run(fmt.Sprintf("Input%d", i), func(t *testing.T) {
			act, err := parse(td.in)
			if err != nil {
				t.Fatalf("unexpected error %v", err)
			}
			if len(act) != len(td.exp) {
				t.Fatalf("act = %d, want %d", act, td.exp)
			}
			for j, v := range act {
				if v != td.exp[j] {
					t.Fatalf("act = %d, want %d", act, td.exp)
				}
			}
		})
	}
}
