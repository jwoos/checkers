package checkers

import (
	"fmt"
)

type Piece struct {
	Coord     *Coordinate
	Direction int
	Type      int
	King      bool
}

func NewPiece(king bool, coord *Coordinate, t int, direction int) *Piece {
	piece := Piece{
		Coord:     coord,
		Direction: direction,
		Type:      t,
		King:      king,
	}

	return &piece
}

func (piece *Piece) String() string {
	var side string
	var pieceType string

	if piece.Type == RED {
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

func (piece *Piece) SetType(t int) {
	piece.Type = t
}

func (piece *Piece) SetDirection(direction int) {
	piece.Direction = direction
}

func (piece *Piece) SetCoordinate(coord *Coordinate) {
	piece.Coord = coord
}

func (piece *Piece) ApplyCoordinate(coord *Coordinate) {
	piece.Coord.ApplyCoordinate(coord)
}

func (piece *Piece) SetRow(row int) {
	piece.Coord.Row = row
}

func (piece *Piece) SetColumn(column int) {
	piece.Coord.Column = column
}
