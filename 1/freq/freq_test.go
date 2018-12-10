package freq_test

import (
	"fmt"
	"testing"

	"github.com/sjones303/aoc-2018/1/freq"
)

func TestSum(t *testing.T) {
	for i, td := range []struct {
		in  []int
		exp int
	}{
		{[]int{1, -2, 3, 1}, 3},
		{[]int{1, 1, 1}, 3},
		{[]int{1, 1, -2}, 0},
		{[]int{-1, -2, -3}, -6},
	} {
		t.Run(fmt.Sprintf("Input%d", i), func(t *testing.T) {
			if act := freq.Sum(td.in); act != td.exp {
				t.Errorf("act = %d, want %d", act, td.exp)
			}
		})
	}
}

func TestMatch(t *testing.T) {
	for i, td := range []struct {
		in  []int
		exp int
	}{
		{[]int{1, -2, 3, 1}, 2},
		{[]int{1, -1}, 0},
		{[]int{3, 3, 4, -2, -4}, 10},
		{[]int{-6, 3, 8, 5, -6}, 5},
		{[]int{7, 7, -2, -7, -4}, 14},
	} {
		t.Run(fmt.Sprintf("Input%d", i), func(t *testing.T) {
			if act := freq.Match(td.in); act != td.exp {
				t.Errorf("act = %d, want %d", act, td.exp)
			}
		})
	}
}
