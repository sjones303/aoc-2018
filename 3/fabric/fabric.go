package fabric

import (
	"fmt"
	"strconv"
)

type Fabric struct {
	record [][]int
	w, h   int
}

func New(w, h int) Fabric {
	r := make([][]int, h)
	for i := range r {
		r[i] = make([]int, w)
	}
	return Fabric{
		record: r,
		w:      w,
		h:      h,
	}
}

func (f *Fabric) Mark(p Piece) error {
	if p.L+1 > f.w {
		return fmt.Errorf("too far right (%d into %d)", p.L, f.w)
	}
	if p.T+1 > f.h {
		return fmt.Errorf("too far down (%d into %d)", p.T, f.h)
	}
	if p.W+p.L > f.w {
		return fmt.Errorf("ran off right (%d+%d into %d)", p.L, p.W, f.w)
	}
	if p.H+p.T > f.h {
		return fmt.Errorf("ran off bottom (%d+%d into %d)", p.T, p.H, f.h)
	}
	for x := 0; x < p.W; x++ {
		for y := 0; y < p.H; y++ {
			f.record[y+p.T][x+p.L]++
		}
	}
	return nil
}

func (f *Fabric) Print() {
	for _, r := range f.record {
		for _, v := range r {
			s := "."
			if v > 0 {
				s = strconv.Itoa(v)
			}
			fmt.Print(s)
		}
		fmt.Print("\n")
	}
}
