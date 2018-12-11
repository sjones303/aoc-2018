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

func TestFirstMatch(t *testing.T) {
	for i, td := range []struct {
		in1, in2 string
		exp      int
	}{
		{"aaaaaa", "aaaaaa", 0},
		{"aaaaaa", "baaaaa", 1},
		{"aaaaaa", "bcaaaa", 2},
		{"aaaaaa", "bcdefg", -1},
		{"baaaaa", "aaaaaa", 1},
		{"bbbbba", "aaaaaa", 5},
	} {
		t.Run(fmt.Sprintf("Input%d", i), func(t *testing.T) {
			m := letters.FirstMatch(td.in1, td.in2)
			if td.exp < 0 {
				if m >= 0 {
					t.Fatalf("m = %d, want %d", m, td.exp)
				}
			}
			if m != td.exp {
				t.Fatalf("m = %d, want %d", m, td.exp)
			}
		})
	}
}
