package checkers

import (
	"reflect"
	"testing"
)

func TestNewState(t *testing.T) {
	rule := NewRule(8, 8, BLACK, TOP, 3, false, false, false)
	board := make([][]Piece, rule.Rows)
	for i := 0; i < rule.Rows; i++ {
		board[i] = make([]Piece, rule.Columns)
	}

	t.Run(
		"Without initialization of board",
		func(innerT *testing.T) {
			state := NewState(rule, false)

			if !reflect.DeepEqual(state.Board, board) {
				innerT.Errorf("Boards are not equal: %v %v", state.Board, board)
			}

			if !reflect.DeepEqual(state.Rules, rule) {
				innerT.Errorf("Rules are not equal")
			}
		},
	)

	t.Run(
		"With initialization of board",
		func(innerT *testing.T) {
			state := NewStateByte(rule, true)

			if reflect.DeepEqual(state.Board, board) {
				innerT.Errorf("Boards are equal: %v %v", state.Board, board)
			}

			if !reflect.DeepEqual(state.Rules, rule) {
				innerT.Errorf("Rules are not equal")
			}
		},
	)
}

func TestPossibleMoves(t *testing.T) {
	rule := NewRule(8, 8, BLACK, TOP, 3, false, false, false)

	t.Run(
		"State uninitialized - No valid moves",
		func(innerT *testing.T) {
			state := NewState(rule, false)
			moves := state.PossibleMoves(NewCoordinate(0, 0))

			if len(moves) != 0 {
				innerT.Errorf("There should be no possible moves: \n%v \n but got %v", state, moves)
			}
		},
	)

	t.Run(
		"State initialized - No valid moves",
		func(innerT *testing.T) {
			state := NewState(rule, true)
			moves := state.PossibleMoves(NewCoordinate(0, 0))

			if len(moves) != 0 {
				innerT.Errorf("There should be no possible moves: \n%v \n but got %v", state, moves)
			}
		},
	)

	t.Run(
		"State initialized - One move",
		func(innerT *testing.T) {
			state := NewState(rule, true)
			moves := state.PossibleMoves(NewCoordinate(2, 0))

			if len(moves) != 1 {
				innerT.Errorf("There should be only one possible move: \n%v \n but got %v", state, moves)
			}
			val, okay := moves[NewCoordinate(3,1)]
			if !okay {
				innerT.Errorf("The key should exist in moves: %v", moves)
			}

			if val.Jump != NO_JUMP {
				innerT.Errorf("It should not be a jump: %v", val)
			}
		},
	)

	t.Run(
		"State initialized - Multiple moves",
		func(innerT *testing.T) {
			state := NewState(rule, true)
			moves := state.PossibleMoves(NewCoordinate(2, 2))

			if len(moves) != 2 {
				innerT.Errorf("There should be only two possible move: \n%v \n but got %v", state, moves)
			}

			destinations := []Coordinate{
				NewCoordinate(3, 3),
				NewCoordinate(3, 1),
			}

			for _, destination := range destinations {
				val, okay := moves[destination]
				if !okay {
					innerT.Errorf("The key should exist in moves: %v", moves)
				}

				if val.Jump != NO_JUMP {
					innerT.Errorf("It should not be a jump: %v", val)
				}
			}
		},
	)
}

func TestPossibleMovesAll(t *testing.T) {
	rule := NewRule(8, 8, BLACK, TOP, 3, false, false, false)

	t.Run(
		"State uninitialized - No moves",
		func(innerT *testing.T) {
			state := NewState(rule, false)
			moves := state.PossibleMovesAll(WHITE)

			if len(moves) != 0 {
				innerT.Errorf("There should be no possible move: \n%v \n but got %v", state, moves)
			}
		},
	)

	t.Run(
		"State initialized - Moves",
		func(innerT *testing.T) {
			state := NewState(rule, true)
			moves := state.PossibleMovesAll(WHITE)

			if len(moves) != 7 {
				innerT.Errorf("There should be no possible move: \n%v \n but got %v", state, moves)
			}
		},
	)
}
