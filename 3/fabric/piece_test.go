package fabric_test

import (
	"fmt"
	"testing"

	"github.com/sjones303/aoc-2018/3/fabric"
)

func TestParsePiece(t *testing.T) {
	for i, td := range []struct {
		in  string
		exp fabric.Piece
	}{
		{"#123 @ 3,2: 5x4", fabric.Piece{123, 3, 2, 5, 4}},
		{"#1 @ 1,3: 4x4", fabric.Piece{1, 1, 3, 4, 4}},
		{"#2 @ 3,1: 4x4", fabric.Piece{2, 3, 1, 4, 4}},
		{"#3 @ 5,5: 2x2", fabric.Piece{3, 5, 5, 2, 2}},
	} {
		t.Run(fmt.Sprintf("Input%d", i), func(t *testing.T) {
			act, err := fabric.ParsePiece(td.in)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			exp := td.exp
			if act.ID != exp.ID {
				t.Errorf("ID = %d, want %d", act.ID, exp.ID)
			}
			if act.L != exp.L {
				t.Errorf("L = %d, want %d", act.L, exp.L)
			}
			if act.T != exp.T {
				t.Errorf("T = %d, want %d", act.T, exp.T)
			}
			if act.W != exp.W {
				t.Errorf("W = %d, want %d", act.W, exp.W)
			}
			if act.H != exp.H {
				t.Errorf("H = %d, want %d", act.H, exp.H)
			}
		})
	}
}
