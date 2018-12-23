package fabric

import (
	"fmt"
	"reflect"
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
					if act := f.Record[y][x]; act != exp {
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

func TestListClaims(t *testing.T) {
	f := New(8, 8)
	mark := func(p Piece) {
		t.Helper()
		err := f.Mark(p)
		if err != nil {
			t.Fatalf("unable to mark: %v", err)
		}
	}
	mark(Piece{ID: 1, L: 1, T: 3, W: 4, H: 4})
	mark(Piece{ID: 2, L: 3, T: 1, W: 4, H: 4})
	mark(Piece{ID: 3, L: 5, T: 5, W: 2, H: 2})

	cs := f.ListClaims()

	for i, exp := range [][]int{
		{2},
		{1},
		{},
	} {
		act, ok := cs[i+1]
		if !ok {
			t.Errorf("claim %d not found", i+1)
			continue
		}
		if !reflect.DeepEqual(act, exp) {
			t.Errorf("claim %d = %d, want %d", i+1, act, exp)
		}
	}
}
