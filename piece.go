package checkers

import (
	"fmt"
)

type Piece struct {
	Coord     *Coordinate
	Direction *Coordinate

	King bool
}

func NewPiece(king bool, coord *Coordinate, direction *Coordinate) *Piece {
	piece := Piece{
		Coord:     coord,
		Direction: direction,
		King:      king,
	}

	return &piece
}

func (piece *Piece) String() string {
	var side string
	var pieceType string

	if piece.Direction.Row == 1 {
		side = "red"
	} else {
		side = "black"
	}

	if piece.King {
		pieceType = "king"
	} else {
		pieceType = "pawn"
	}

	return fmt.Sprintf("%s %s %#v", side, pieceType, piece.Coord)
}

func (piece *Piece) GoString() string {
	return piece.String()
}

func (piece *Piece) SetKing(king bool) {
	piece.King = king
}

func (piece *Piece) SetDirection(direction *Coordinate) {
	piece.Direction = direction
}

func (piece *Piece) SetCoord(coord *Coordinate) {
	piece.Coord = coord
}

func (piece *Piece) SetRow(row int) {
	piece.Coord.Row = row
}

func (piece *Piece) SetColumn(column int) {
	piece.Coord.Column = column
}

func (piece *Piece) Diagonal() map[*Coordinate][]*Coordinate {
	set := make(map[*Coordinate][]*Coordinate, 4)

	for i := -1; i <= 1; i++ {
		if i == 0 {
			continue
		}

		for j := -1; j <= 1; j++ {
			if j == 0 {
				continue
			}

			set[Coordinate{Row: i, Column: j}] = nil
		}
	}

	for k, v := range set {
		if piece.King || piece.Direction.Row == k.Row {
			set[k] = piece.Coord.Diagonal(k)
		}
	}

	return set
}
