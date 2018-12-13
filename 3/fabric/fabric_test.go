package fabric

import (
	"fmt"
	"testing"
)

type markTD struct {
	l, t, w, h int
	exp        [][]int
	fail       bool
}

func TestMark(t *testing.T) {
	for i, td := range []markTD{
		{
			0, 0, 2, 2,
			[][]int{{1, 1, 0, 0},
				{1, 1, 0, 0},
				{0, 0, 0, 0},
				{0, 0, 0, 0}},
			false,
		},
		{
			1, 1, 3, 3,
			[][]int{{0, 0, 0, 0},
				{0, 1, 1, 1},
				{0, 1, 1, 1},
				{0, 1, 1, 1}},
			false,
		},
		{
			2, 2, 2, 2,
			[][]int{{0, 0, 0, 0},
				{0, 0, 0, 0},
				{0, 0, 1, 1},
				{0, 0, 1, 1}},
			false,
		},
		{
			2, 2, 2, 3,
			nil,
			true,
		},
		{
			2, 2, 3, 2,
			nil,
			true,
		},
		{
			5, 0, 1, 1,
			nil,
			true,
		},
		{
			0, 5, 1, 1,
			nil,
			true,
		},
	} {
		t.Run(fmt.Sprintf("Input%d", i), func(t *testing.T) {
			p := newPiece(td)
			f := New(4, 4)
			err := f.Mark(p)
			if td.fail {
				if err == nil {
					t.Fatal("unexpected success")
				}
			} else {
				if err != nil {
					t.Fatalf("unexpected error: %v", err)
				}
			}
			for y, r := range td.exp {
				for x, exp := range r {
					if act := f.record[y][x]; act != exp {
						t.Errorf("(%d,%d) = %d, want %d",
							x, y, act, exp)
					}
				}
			}
		})
	}
}

func newPiece(td markTD) Piece {
	return Piece{
		L: td.l,
		T: td.t,
		W: td.w,
		H: td.h,
	}
}
