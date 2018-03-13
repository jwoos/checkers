package checkers

import (
	"fmt"
	"strings"
)

// row major
type State struct {
	Rules Rule

	Board [][]*Piece
}

func NewState(rule Rule) *State {
	state := State{
		Rules: rule,
		Board: make([][]*Piece, rule.Rows),
	}

	for i := 0; i < rule.Rows; i++ {
		state.Board[i] = make([]*Piece, rule.Columns)
	}

	return &state
}

func (state *State) String() string {
	var str strings.Builder

	for i := 0; i < state.Rules.Rows; i++ {
		for j := 0; j < state.Rules.Columns; j++ {
			if state.Board[i][j] != nil {
				if state.Board[i][j].Direction.Row == 1 {
					str.WriteRune('x')
				} else {
					str.WriteRune('o')
				}
			} else {
				str.WriteRune('.')
			}
		}

		str.WriteRune('\n')
	}

	return str.String()
}

func (state *State) GoString() string {
	return state.String()
}

func (state *State) ValidateMove(from *Coordinate, to *Coordinate) error {
	// check bounds
	if to.Row < 0 || to.Row > state.Rules.Columns {
		return NewBoundsError(fmt.Sprintf("cannot move from %#v to %#v - x is out of range", from, to))
	}

	if to.Column < 0 || to.Column > state.Rules.Rows {
		return NewBoundsError(fmt.Sprintf("cannot move from %#v to %#v - y is out of range", from, to))
	}

	// check the space is empty
	if state.Board[to.Row][to.Column] != nil {
		return NewMovementError("the position is occupied")
	}

	// check that if there is a valid captured state, it's taken
	return nil
}

func (state *State) MovePiece(from *Coordinate, to *Coordinate) error {
	err := state.ValidateMove(from, to)
	if err != nil {
		return err
	}

	piece := state.Board[from.Row][from.Column]
	piece.SetCoord(to)

	return nil
}

//func (state *State) PossibleMoves(piece *Piece) map[*Coordinate][]*Coordinate {}
