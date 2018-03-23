package checkers

import (
	"fmt"
	"strings"
)

const (
	BLANK byte = iota
	WHITE byte = iota
	BLACK byte = iota
)

// row major
type State struct {
	Rules Rule

	Board [][]Piece
	Turn  byte

	White map[Coordinate]Piece
	Black map[Coordinate]Piece
}

func NewState(rule Rule, instantiateBoard bool) *State {
	state := State{
		Rules: rule,
		Board: make([][]Piece, rule.Rows),
		Turn:  rule.First,
		White: make(map[Coordinate]Piece),
		Black: make(map[Coordinate]Piece),
	}

	for i := 0; i < rule.Rows; i++ {
		state.Board[i] = make([]Piece, rule.Columns)
	}

	if instantiateBoard {
		var topMap *map[Coordinate]Piece
		var bottomMap *map[Coordinate]Piece
		var blackSide int
		var whiteSide int
		var top byte
		var bottom byte

		if rule.First == BLACK {
			blackSide = rule.Side
			whiteSide = blackSide ^ int(BLACK)
		} else {
			whiteSide = rule.Side
			blackSide = whiteSide ^ int(BLACK)
		}

		if blackSide == TOP {
			top = BLACK
			bottom = WHITE

			topMap = &state.Black
			bottomMap = &state.White
		} else {
			top = WHITE
			bottom = BLACK

			topMap = &state.White
			bottomMap = &state.Black
		}

		var piece Piece
		var coordinate Coordinate
		for i := 0; i < rule.RowsToFill; i++ {
			for j := 0; j < rule.Columns; j++ {
				if ((rule.Rows - 1 - i) % 2) == (j % 2) {
					coordinate = NewCoordinate(rule.Rows-1-i, j)
					piece = NewPiece(false, top)
					state.Board[rule.Rows-1-i][j] = piece
					(*topMap)[coordinate] = piece
				}

				if (i % 2) == (j % 2) {
					coordinate = NewCoordinate(i, j)
					piece = NewPiece(false, bottom)
					state.Board[i][j] = piece
					(*bottomMap)[coordinate] = piece
				}
			}
		}
	}

	return &state
}

func (state *State) String() string {
	var str strings.Builder

	for i := state.Rules.Columns - 1; i >= 0; i-- {
		str.WriteString(fmt.Sprintf(" %d | ", i))
		for j := 0; j < state.Rules.Columns; j++ {
			if state.Board[i][j].Side != BLANK {
				if state.Board[i][j].Side == BLACK {
					str.WriteString(" b ")
				} else {
					str.WriteString(" w ")
				}
			} else {
				str.WriteString(" . ")
			}
		}

		str.WriteRune('\n')
	}

	str.WriteString("     ")
	str.WriteString(strings.Repeat("---", state.Rules.Columns))
	str.WriteRune('\n')
	str.WriteString("     ")
	for i := 0; i < state.Rules.Columns; i++ {
		str.WriteString(fmt.Sprintf(" %d ", i))
	}

	return str.String()
}

func (state *State) GoString() string {
	return state.String()
}

func (state *State) Copy() *State {
	newState := NewState(state.Rules, false)

	newState.Turn = state.Turn

	var row int
	var column int

	for coord, _ := range state.White {
		row = coord.Row
		column = coord.Column

		newState.Board[row][column] = state.Board[row][column]
		newState.White[coord] = state.Board[row][column]
	}

	for coord, _ := range state.Black {
		row = coord.Row
		column = coord.Column

		newState.Board[row][column] = state.Board[row][column]
		newState.Black[coord] = state.Board[row][column]
	}

	return newState
}

func (state *State) CheckBound(coord Coordinate) bool {
	okay := true

	if coord.Row < 0 || coord.Row >= state.Rules.Rows {
		okay = false
	}

	if coord.Column < 0 || coord.Column >= state.Rules.Columns {
		okay = false
	}

	return okay
}

func (state *State) Move(move Move) {
	state.Board[move.To.Row][move.To.Column] = state.Board[move.From.Row][move.From.Column]
	state.Board[move.From.Row][move.From.Column] = Piece{}

	if state.Turn == WHITE {
		state.White[move.To] = state.White[move.From]
		delete(state.White, move.From)
	} else {
		state.Black[move.To] = state.Black[move.From]
		delete(state.Black, move.From)
	}

	// there was a jump
	if move.Jump != NO_JUMP {
		state.Board[move.Jump.Row][move.Jump.Column] = Piece{}

		if state.Turn == BLACK {
			delete(state.White, move.Jump)
		} else {
			delete(state.Black, move.Jump)
		}
	}

	if state.Turn == BLACK {
		state.Turn = WHITE
	} else {
		state.Turn = BLACK
	}
}

func (state *State) Validate(from Coordinate, to Coordinate) error {
	piece := state.Board[from.Row][from.Column]

	if state.Turn != piece.Side {
		return NewMovementError(ERROR_MOVE_TURN)
	}

	if piece.Side == BLANK {
		return NewMovementError(ERROR_MOVE_BLANK)
	}

	if piece.Side != state.Turn {
		return NewMovementError(ERROR_MOVE_WRONG)
	}

	// TODO check the jump is over the opponent's pieces

	var dir int
	if state.Rules.First == state.Turn {
		if state.Rules.Side == TOP {
			dir = -1
		} else {
			dir = 1
		}
	} else {
		if state.Rules.Side == TOP {
			dir = 1
		} else {
			dir = -1
		}
	}
	if to.Row - from.Row != dir {
		return NewMovementError(ERROR_MOVE_BACK)
	}

	if abs(from.Row-to.Row) != abs(from.Column-to.Column) {
		return NewMovementError(ERROR_MOVE_INVALID)
	}

	if (abs(from.Row-to.Row) > 2) || (abs(from.Column-to.Column) > 2) {
		return NewMovementError(ERROR_MOVE_INVALID)
	}

	// check bounds
	if to.Row < 0 || to.Row > state.Rules.Columns {
		return NewMovementError(ERROR_MOVE_BOUNDS)
	}

	if to.Column < 0 || to.Column > state.Rules.Rows {
		return NewMovementError(ERROR_MOVE_BOUNDS)
	}

	// check the space is empty
	if state.Board[to.Row][to.Column].Side != BLANK {
		return NewMovementError(ERROR_MOVE_OCCUPIED)
	}

	// TODO check that if there is a valid captured state, it's taken

	return nil
}

func (state *State) PossibleMoves(from Coordinate) map[Coordinate]Move {
	moves := make(map[Coordinate]Move)
	var dir int
	if state.Rules.First == state.Board[from.Row][from.Column].Side {
		if state.Rules.Side == TOP {
			dir = -1
		} else {
			dir = 1
		}
	} else {
		if state.Rules.Side == TOP {
			dir = 1
		} else {
			dir = -1
		}
	}

	directions := []Coordinate{
		NewCoordinate(dir, 1),
		NewCoordinate(dir, -1),
	}

	piece := state.Board[from.Row][from.Column]
	if piece.King {
		directions = append(directions, NewCoordinate(-dir, 1), NewCoordinate(-dir, -1))
	}

	jumpPresent := false

	for _, direction := range directions {
		target := NewCoordinate(from.Row+direction.Row, from.Column+direction.Column)

		// check out of bounds
		if !state.CheckBound(target) {
			continue
		}

		// check the space is empty
		if state.Board[target.Row][target.Column].Side != BLANK {
			if state.Board[target.Row][target.Column].Side != state.Turn {
				jump := NewCoordinate(target.Row+direction.Row, target.Column+direction.Column)

				if !state.CheckBound(jump) || state.Board[jump.Row][jump.Column].Side != BLANK {
					continue
				}

				jumpPresent = true
				moves[jump] = NewMove(from, jump, target)
			}

			continue
		}

		if !jumpPresent {
			moves[target] = NewMove(from, target, NO_JUMP)
		}
	}

	return moves
}
