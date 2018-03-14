package checkers

import (
	"reflect"
	"testing"
)

func TestNewState(t *testing.T) {
	rule := NewRule(8, 8, BLACK, TOP, 3, false, false, false)
	board := make([][]*Piece, rule.Rows)
	for i := 0; i < rule.Rows; i++ {
		board[i] = make([]*Piece, rule.Columns)
	}

	t.Run(
		"Without initialization of board",
		func(innerT *testing.T) {
			state := NewState(rule, false)

			if !reflect.DeepEqual(state.Board, board) {
				innerT.Errorf("Boards are not equal: %v %v", state.Board, board)
			}

			if state.Rules != rule {
				innerT.Errorf("Rules are not equal")
			}
		},
	)

	t.Run(
		"With initialization of board",
		func(innerT *testing.T) {
			state := NewState(rule, true)

			if reflect.DeepEqual(state.Board, board) {
				innerT.Errorf("Boards are equal: %v %v", state.Board, board)
			}

			if state.Rules != rule {
				innerT.Errorf("Rules are not equal")
			}
		},
	)
}

func TestPossibleMoves(t *testing.T) {
	rule := NewRule(8, 8, BLACK, TOP, 3, false, false, false)
	state := NewState(rule, true)

	t.Run(
		"Corner - no moves possible",
		func(innerT *testing.T) {
			moves := state.PossibleMoves(state.Board[0][0], false)

			for _, v := range moves {
				if v != nil {
					innerT.Errorf("Expected nil but got %v", v)
				}
			}
		},
	)
}

func TestMovePiece(t *testing.T) {
	rule := NewRule(8, 8, BLACK, BOTTOM, 3, false, false, false)
	state := NewState(rule, true)

	t.Run(
		"Moves piece",
		func(innerT *testing.T) {
			err := state.MovePiece(state.Board[2][0], NewCoordinate(1, 1))
			if err != nil {
				innerT.Error(err)
				return
			}

			if state.Board[2][0] != nil {
				innerT.Errorf("Piece still at original location %v", state)
			}

			if state.Board[3][1] == nil {
				innerT.Errorf("Piece not at new location")
			}
		},
	)
}
