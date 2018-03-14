package checkers

import (
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

	Board [][]*Piece
	Turn  byte

	White map[*Piece]bool
	Black map[*Piece]bool
}

func NewState(rule Rule, instantiateBoard bool) *State {
	state := State{
		Rules: rule,
		Board: make([][]*Piece, rule.Rows),
		Turn:  rule.First,
		White: make(map[*Piece]bool),
		Black: make(map[*Piece]bool),
	}

	for i := 0; i < rule.Rows; i++ {
		state.Board[i] = make([]*Piece, rule.Columns)
	}

	if instantiateBoard {
		var topMap *map[*Piece]bool
		var bottomMap *map[*Piece]bool
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

		var piece *Piece
		var coordinate Coordinate
		for i := 0; i < rule.RowsToFill; i++ {
			for j := 0; j < rule.Columns; j++ {
				if ((rule.Rows - 1 - i) % 2) == (j % 2) {
					coordinate = NewCoordinate(rule.Rows-1-i, j)
					piece = NewPiece(false, coordinate, top, -1)
					state.Board[rule.Rows-1-i][j] = piece
					(*topMap)[piece] = true
				}

				if (i % 2) == (j % 2) {
					coordinate = NewCoordinate(i, j)
					piece = NewPiece(false, coordinate, bottom, 1)
					state.Board[i][j] = piece
					(*bottomMap)[piece] = true
				}
			}
		}
	}

	return &state
}

func (state *State) ByteArray() [][]byte {
	arr := make([][]byte, state.Rules.Rows)
	rule := state.Rules

	for i := 0; i < rule.Rows; i++ {
		arr[i] = make([]byte, rule.Columns)
	}

	var row int
	var column int
	for piece, _ := range state.White {
		row = piece.Coord.Row
		column = piece.Coord.Column

		arr[row][column] = WHITE
	}

	for piece, _ := range state.Black {
		row = piece.Coord.Row
		column = piece.Coord.Column

		arr[row][column] = BLACK
	}

	return arr
}

func (state *State) String() string {
	var str strings.Builder

	for i := state.Rules.Columns - 1; i >= 0; i-- {
		for j := 0; j < state.Rules.Columns; j++ {
			if state.Board[i][j] != nil {
				if state.Board[i][j].Type == BLACK {
					str.WriteRune('b')
				} else {
					str.WriteRune('w')
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

/*
 *func Validate(piece *Piece, move Move) error {
 *
 *}
 */

func (state *State) ValidateMove(piece *Piece, dir Coordinate) error {
	coord := piece.Coord.ApplyCoordinate(dir)
	return state.ValidateMoveTo(piece, coord)
}

func (state *State) ValidateMoveTo(piece *Piece, to Coordinate) error {
	from := piece.Coord

	if state.Turn != piece.Type {
		return NewMovementError(ERROR_MOVE_TURN)
	}

	// TODO check the jump is over the opponent's pieces

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
	if state.Board[to.Row][to.Column] != nil {
		return NewMovementError(ERROR_MOVE_OCCUPIED)
	}

	// TODO check that if there is a valid captured state, it's taken

	return nil
}

func (state *State) MovePiece(piece *Piece, application Coordinate) error {
	err := state.ValidateMove(piece, application)
	if err != nil {
		return err
	}

	state.Board[piece.Coord.Row][piece.Coord.Column] = nil
	piece.ApplyCoordinate(application)
	state.Board[piece.Coord.Row][piece.Coord.Column] = piece

	state.Turn ^= BLACK

	return nil
}

func (state *State) MovePieceTo(piece *Piece, to Coordinate) error {
	err := state.ValidateMoveTo(piece, to)
	if err != nil {
		return err
	}

	state.Board[piece.Coord.Row][piece.Coord.Column] = nil
	piece.SetCoordinate(to)
	state.Board[piece.Coord.Row][piece.Coord.Column] = piece

	state.Turn ^= BLACK

	return nil
}

func (state *State) Move(piece *Piece, move Move) {
	state.Board[piece.Coord.Row][piece.Coord.Column] = nil
	piece.SetCoordinate(move.To)
	state.Board[piece.Coord.Row][piece.Coord.Column] = piece

	// there was a jump
	if move.Jump != NO_JUMP {
		if state.Turn == BLACK {
			delete(state.White, state.Board[move.Jump.Row][move.Jump.Column])
		} else {
			delete(state.Black, state.Board[move.Jump.Row][move.Jump.Column])
		}
		state.Board[move.Jump.Row][move.Jump.Column] = nil
	}

	state.Turn ^= BLACK
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

func (state *State) PossibleMoves(piece *Piece, jumpOnly bool) map[Coordinate]Move {
	moves := make(map[Coordinate]Move)
	dir := piece.Direction

	if dir == 0 {
		dir = 1
	}

	directions := []Coordinate{
		NewCoordinate(dir, 1),
		NewCoordinate(dir, -1),
	}

	if piece.King {
		directions = append(directions, NewCoordinate(-dir, 1), NewCoordinate(-dir, -1))
	}

	for _, direction := range directions {
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

				moves[jump] = NewMove(piece.Coord, jump, target)
			}

			continue
		}

		if !jumpOnly {
			moves[target] = NewMove(piece.Coord, target, NO_JUMP)
		}
	}

	return moves
}
