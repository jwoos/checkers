package checkers

import (
	"strings"
)

const (
	RED   int = iota
	BLACK int = iota
)

// row major
type State struct {
	Rules Rule

	Board [][]*Piece
	Turn  int
}

func NewState(rule Rule, instantiateBoard bool) *State {
	state := State{
		Rules: rule,
		Board: make([][]*Piece, rule.Rows),
		Turn:  rule.First,
	}

	for i := 0; i < rule.Rows; i++ {
		state.Board[i] = make([]*Piece, rule.Columns)
	}

	if instantiateBoard {
		var blackSide int
		var redSide int
		var top int
		var bottom int

		if rule.First == BLACK {
			blackSide = rule.Side
			redSide = redSide ^ BLACK
		} else {
			redSide = rule.Side
			blackSide = redSide ^ BLACK
		}

		if blackSide == TOP {
			top = BLACK
			bottom = RED
		} else {
			top = RED
			bottom = BLACK
		}

		for i := 0; i < rule.RowsToFill; i++ {
			for j := 0; j < rule.Columns; j++ {
				if ((rule.Rows - 1 - i) % 2) == (j % 2) {
					coordinate := NewCoordinate(rule.Rows - 1 - i, j)
					state.Board[rule.Rows - 1 - i][j] = NewPiece(false, coordinate, top, -1)
				}

				if (i % 2) == (j % 2) {
					coordinate := NewCoordinate(i, j)
					state.Board[i][j] = NewPiece(false, coordinate, bottom, 1)
				}
			}
		}
	}

	return &state
}

func (state *State) String() string {
	var str strings.Builder

	for i := state.Rules.Columns - 1; i >= 0; i-- {
		for j := 0; j < state.Rules.Columns; j++ {
			if state.Board[i][j] != nil {
				if state.Board[i][j].Type == BLACK {
					str.WriteRune('b')
				} else {
					str.WriteRune('r')
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

func (state *State) ValideMove(piece *Piece, to *Coordinate) bool {
	//from := piece.Coord

	// check bounds
	if to.Row < 0 || to.Row > state.Rules.Columns {
		return false
	}

	if to.Column < 0 || to.Column > state.Rules.Rows {
		return false
	}

	// check the space is empty
	if state.Board[to.Row][to.Column] != nil {
		return false
	}

	// check that if there is a valid captured state, it's taken
	return true
}

func (state *State) MovePiece(piece *Piece, application *Coordinate) error {
	err := state.ValideMove(piece, application)
	if err {
		return NewMovementError("Invalid move")
	}

	state.Board[piece.Coord.Row][piece.Coord.Column] = nil
	piece.ApplyCoordinate(application)
	state.Board[piece.Coord.Row][piece.Coord.Column] = piece

	return nil
}

func (state *State) MovePieceTo(piece *Piece, to *Coordinate) error {
	err := state.ValideMove(piece, to)
	if err {
		return NewMovementError("Invalid move")
	}

	state.Board[piece.Coord.Row][piece.Coord.Column] = nil
	piece.SetCoordinate(to)
	state.Board[piece.Coord.Row][piece.Coord.Column] = piece

	return nil
}

func (state *State) CheckBound(coord *Coordinate) bool {
	okay := true

	if coord.Row < 0 || coord.Row >= state.Rules.Rows {
		okay = false
	}

	if coord.Column < 0 || coord.Column >= state.Rules.Columns {
		okay = false
	}

	return okay
}

func (state *State) PossibleMoves(piece *Piece, jumpOnly bool) map[Coordinate]*Coordinate {
	moves := make(map[Coordinate]*Coordinate)
	dir := piece.Direction

	if dir == 0 {
		dir = 1
	}

	moves[Coordinate{Row: dir, Column: 1}] = nil
	moves[Coordinate{Row: dir, Column: -1}] = nil

	if piece.King {
		moves[Coordinate{Row: -dir, Column: 1}] = nil
		moves[Coordinate{Row: -dir, Column: -1}] = nil
	}

	for direction, _ := range moves {
		target := NewCoordinate(piece.Coord.Row+direction.Row, piece.Coord.Column+direction.Column)

		// check out of bounds
		if !state.CheckBound(target) {
			continue
		}

		// check the space is empty
		if state.Board[target.Row][target.Column] != nil {
			if (state.Board[target.Row][target.Column].Type ^ piece.Type) == 1 {
				jump := NewCoordinate(target.Row+direction.Row, target.Column+direction.Column)

				if !state.CheckBound(jump) {
					continue
				}

				moves[direction] = jump
			}

			continue
		}

		if !jumpOnly {
			moves[direction] = target
		}
	}

	return moves
}
