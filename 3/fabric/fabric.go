package fabric

import "errors"

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
	return errors.New("not implemented")
}
