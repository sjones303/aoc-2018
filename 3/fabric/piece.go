package fabric

import "errors"

type Piece struct {
	ID   int
	L, T int
	W, H int
}

func ParsePiece(in string) (Piece, error) {
	return Piece{}, errors.New("not implemented")
}
