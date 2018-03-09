package checkers


import (
	"fmt"
)


type Piece struct {
	King bool
	Coord Coordinate
	Side bool
}


func NewPiece(king bool, coord Coordinate, side bool) *Piece {
	piece := Piece{
		King: king,
		Coord: coord,
		Side: side
	}

	return &piece
}

func (piece *Piece) String() string {
	var side string
	var pieceType string

	if piece.side {
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

func (piece *Piece) SetSide(side bool) {
	piece.Side = side
}

func (piece *Piece) SetCoord(coord Coordinate) {
	piece.Coord = coord
}

func (piece *Piece) SetRow(row uint) {
	piece.Coord.Row = row
}

func (piece *Piece) SetColumn (column uint) {
	piece.Coord.Column = column
}
