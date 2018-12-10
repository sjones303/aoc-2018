package letters_test

import (
	"fmt"
	"testing"

	"github.com/sjones303/aoc-2018/2/letters"
)

func TestCount(t *testing.T) {
	for i, td := range []struct {
		in     string
		e2, e3 bool
	}{
		{"abcdef", false, false},
		{"bababc", true, true},
		{"abbcde", true, false},
		{"abcccd", false, true},
		{"aabcdd", true, false},
		{"abcdee", true, false},
		{"ababab", false, true},
	} {
		t.Run(fmt.Sprintf("Input%d", i), func(t *testing.T) {
			a2, a3 := letters.Count(td.in)
			if a2 != td.e2 {
				t.Errorf("a2 = %t, want %t", a2, td.e2)
			}
			if a3 != td.e3 {
				t.Errorf("a3 = %t, want %t", a3, td.e3)
			}
		})
	}
}

func TestHamming1(t *testing.T) {
	type info struct {
		v string
		m bool
	}
	in := []info{
		{"abcde", false},
		{"fghij", true},
		{"klmno", false},
		{"pqrst", false},
		{"fguij", true},
		{"axcye", false},
		{"wvxyz", false},
	}
	for i, v1 := range in[:len(in)-1] {
		for j, v2 := range in[i+1:] {
			t.Run(fmt.Sprintf("Input%d-%d", i, j+i+1),
				func(t *testing.T) {
					var (
						act = letters.Hamming1(v1.v, v2.v)
						exp = v1.m && v2.m
					)
					if act != exp {
						t.Errorf("act = %t, want %t",
							act, exp)
					}
				},
			)
		}
	}
}
