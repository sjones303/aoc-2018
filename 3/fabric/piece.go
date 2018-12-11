package fabric

import "fmt"

type Piece struct {
	ID   int
	L, T int
	W, H int
}

func ParsePiece(in string) (Piece, error) {
	p := Piece{}
	n, err := fmt.Sscanf(in, "#%d @ %d,%d: %dx%d", &p.ID, &p.L, &p.T, &p.W, &p.H)
	if err != nil {
		return p, err
	}
	if n != 5 {
		return p, fmt.Errorf("only %d values found", n)
	}
	return p, nil
}
