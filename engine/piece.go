package checkers

import (
	"fmt"
)

type Piece struct {
	Side      byte
	King      bool
}

func NewPiece(king bool, side byte) Piece {
	return Piece{
		Side: side,
		King: king,
	}
}

func (piece Piece) String() string {
	var side string
	var pieceType string

	if piece.Side == WHITE {
		side = "white"
	} else {
		side = "black"
	}

	if piece.King {
		pieceType = "king"
	} else {
		pieceType = "pawn"
	}

	return fmt.Sprintf("%s %s", side, pieceType)
}

func (piece Piece) GoString() string {
	return piece.String()
}
