package checkers

import (
	"reflect"
	"testing"
)

func TestNewStateByte(t *testing.T) {
	rule := NewRule(8, 8, BLACK, TOP, 3, false, false, false)
	board := make([][]byte, rule.Rows)
	for i := 0; i < rule.Rows; i++ {
		board[i] = make([]byte, rule.Columns)
	}

	t.Run(
		"Without initialization of board",
		func(innerT *testing.T) {
			state := NewStateByte(rule, false)

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

func TestCopy(t *testing.T) {
	rule := NewRule(8, 8, BLACK, TOP, 3, false, false, false)
	state := NewStateByte(rule, true)

	t.Run(
		"Copying works",
		func(innerT *testing.T) {
			newState := state.Copy()

			if !reflect.DeepEqual(state, newState) {
				innerT.Errorf("States are not equal: %v %v", state, newState)
			}
		},
	)

	t.Run(
		"Board is copied",
		func(innerT *testing.T) {
			newState := state.Copy()

			newState.Board[0][0] = 5

			if reflect.DeepEqual(state.Board, newState.Board) {
				innerT.Errorf("States are equal: %v\n%v", state, newState)
			}
		},
	)

	t.Run(
		"Pieces are copied",
		func(innerT *testing.T) {
			newState := state.Copy()

			delete(newState.White, NewCoordinate(0, 0))
			delete(newState.Black, NewCoordinate(7, 7))

			if reflect.DeepEqual(state.White, newState.White) {
				innerT.Errorf("States are equal: %v\n%v", state.White, newState.White)
			}

			if reflect.DeepEqual(state.Black, newState.Black) {
				innerT.Errorf("States are equal: %v\n%v", state.Black, newState.Black)
			}
		},
	)
}
