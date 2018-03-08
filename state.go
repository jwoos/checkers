package checkers


import (
	"fmt"
)


type State struct {
	Rules Rule

	Board [][]*Piece
}

func NewState(rule Rule) *State {
	state := State{
		Rules: rule
		Board: make([][]*Piece, rows)
	}

	for i := 0; i < rows; i++ {
		state.Board[i] = make([]Piece, columns)
	}

	return &state
}

func (state *State) String() string {
}

func (state *State) GoString() string {

}

func (state *State) ValidateMove(from Coordinate, to Coordinate) error {
	// check bounds
	if to.X < 0 || to.X > state.Rules.Columns {
		return NewBoundsError(fmt.Sprintf("cannot move from %#v to %#v - x is out of range", from, to))
	}

	if to.Y < 0 || to.Y > state.Rules.Rows {
		return NewBoundsError(fmt.Sprintf("cannot move from %#v to %#v - y is out of range", from, to))
	}

	// check the space is empty
}

func (state *State) MovePiece(from Coordinate, to Coordinate) error {
	err := state.ValidateMove(from, to)
	if err != nil {
		return err
	}

	piece := state.Board[from.x][from.y]
	piece.SetCoord(to)

	return nil
}

func (state *State) PossibleMoves(from Coordinate, to Coordinate) []*State {

}
