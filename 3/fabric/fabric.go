package fabric

import (
	"fmt"
	"strconv"
)

type Fabric struct {
	record [][][]int
	claims map[int]map[int]struct{}
	w, h   int
}

func New(w, h int) Fabric {
	r := make([][][]int, h)
	for i := range r {
		r[i] = make([][]int, w)
	}
	return Fabric{
		record: r,
		claims: map[int]map[int]struct{}{},
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

	f.claims[p.ID] = make(map[int]struct{})
	cs := map[int]struct{}{}

	for x := 0; x < p.W; x++ {
		for y := 0; y < p.H; y++ {
			v := f.record[y+p.T][x+p.L]

			for _, c := range v {
				f.claims[c][p.ID] = struct{}{}
				cs[c] = struct{}{}
			}

			v = append(v, p.ID)
			f.record[y+p.T][x+p.L] = v
		}
	}
	f.claims[p.ID] = cs
	return nil
}

func (f *Fabric) Record() [][]int {
	ret := make([][]int, f.h)
	for i, rr := range f.record {
		r := make([]int, f.w)
		for i, v := range rr {
			r[i] = len(v)
		}
		ret[i] = r
	}
	return ret
}

func (f *Fabric) ListClaims() map[int][]int {
	ret := make(map[int][]int, len(f.claims))
	for k, v := range f.claims {
		cs := make([]int, 0, len(v))
		for i, _ := range v {
			cs = append(cs, i)
		}
		ret[k] = cs
	}
	return ret
}

func (f *Fabric) Print() {
	for _, r := range f.Record() {
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
