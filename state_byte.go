package checkers

import (
	"fmt"
	"strings"
)

type StateByte struct {
	Rules Rule

	Board [][]byte
	Turn  byte
	White map[Coordinate]bool
	Black map[Coordinate]bool
}

func NewStateByte(rule Rule, instantiateBoard bool) *StateByte {
	state := StateByte{
		Rules: rule,
		Board: make([][]byte, rule.Rows),
		Turn:  rule.First,
		White: make(map[Coordinate]bool),
		Black: make(map[Coordinate]bool),
	}

	for i := 0; i < rule.Rows; i++ {
		state.Board[i] = make([]byte, rule.Columns)
	}

	if instantiateBoard {
		var topMap *map[Coordinate]bool
		var bottomMap *map[Coordinate]bool
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

		var coordinate Coordinate
		for i := 0; i < rule.RowsToFill; i++ {
			for j := 0; j < rule.Columns; j++ {
				if ((rule.Rows - 1 - i) % 2) == (j % 2) {
					coordinate = NewCoordinate(rule.Rows-1-i, j)
					state.Board[rule.Rows-1-i][j] = top
					(*topMap)[coordinate] = true
				}

				if (i % 2) == (j % 2) {
					coordinate = NewCoordinate(i, j)
					state.Board[i][j] = bottom
					(*bottomMap)[coordinate] = true
				}
			}
		}
	}

	return &state
}

func (state *StateByte) String() string {
	var str strings.Builder

	for i := state.Rules.Columns - 1; i >= 0; i-- {
		str.WriteString(fmt.Sprintf(" %d | ", i))
		for j := 0; j < state.Rules.Columns; j++ {
			if state.Board[i][j] != BLANK {
				if state.Board[i][j] == BLACK {
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

func (state *StateByte) GoString() string {
	return state.String()
}

func (state *StateByte) Copy() *StateByte {
	newState := NewStateByte(state.Rules, false)

	newState.Turn = state.Turn

	var row int
	var column int

	for coord, _ := range state.White {
		row = coord.Row
		column = coord.Column

		newState.Board[row][column] = WHITE
		newState.White[coord] = true
	}

	for coord, _ := range state.Black {
		row = coord.Row
		column = coord.Column

		newState.Board[row][column] = BLACK
		newState.Black[coord] = true
	}

	return newState
}

func (state *StateByte) CheckBound(coord Coordinate) bool {
	okay := true

	if coord.Row < 0 || coord.Row >= state.Rules.Rows {
		okay = false
	}

	if coord.Column < 0 || coord.Column >= state.Rules.Columns {
		okay = false
	}

	return okay
}

func (state *StateByte) Move(move Move) {
	state.Board[move.From.Row][move.From.Column] = BLANK
	state.Board[move.To.Row][move.To.Column] = state.Turn

	if state.Turn == WHITE {
		delete(state.White, move.From)
		state.White[move.To] = true
	} else {
		delete(state.Black, move.From)
		state.Black[move.To] = true
	}

	// there was a jump
	if move.Jump != NO_JUMP {
		state.Board[move.Jump.Row][move.Jump.Column] = BLANK

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

func (state *StateByte) Validate(from Coordinate, to Coordinate) error {
	piece := state.Board[from.Row][from.Column]

	if state.Turn != piece {
		return NewMovementError(ERROR_MOVE_TURN)
	}

	if piece == BLANK {
		return NewMovementError(ERROR_MOVE_BLANK)
	}

	if piece != state.Turn {
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
	if state.Board[to.Row][to.Column] != BLANK {
		return NewMovementError(ERROR_MOVE_OCCUPIED)
	}

	// TODO check that if there is a valid captured state, it's taken

	return nil
}

func (state *StateByte) PossibleMoves(from Coordinate) map[Coordinate]Move {
	moves := make(map[Coordinate]Move)
	var dir int
	if state.Rules.First == state.Board[from.Row][from.Column] {
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

	jumpPresent := false

	for _, direction := range directions {
		target := NewCoordinate(from.Row+direction.Row, from.Column+direction.Column)

		if !state.CheckBound(target) {
			continue
		}

		if state.Board[target.Row][target.Column] != BLANK {
			if state.Board[target.Row][target.Column] != state.Turn {
				jump := NewCoordinate(target.Row+direction.Row, target.Column+direction.Column)

				if !state.CheckBound(jump) || state.Board[jump.Row][jump.Column] != BLANK {
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

	// jump must be taken
	if jumpPresent {
		for key, move := range moves {
			if move.Jump == NO_JUMP {
				delete(moves, key)
			}
		}
	}

	return moves
}

func (state *StateByte) PossibleMovesAll(turn byte) map[Move]bool {
	moves := make(map[Move]bool)

	var dir int
	if state.Rules.First == turn {
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

	var pieces map[Coordinate]bool
	if turn == WHITE {
		pieces = state.White
	} else {
		pieces = state.Black
	}

	jumpPresent := false

	for _, direction := range directions {
		for from, _ := range pieces {
			target := NewCoordinate(from.Row+direction.Row, from.Column+direction.Column)

			if !state.CheckBound(target) {
				continue
			}

			if state.Board[target.Row][target.Column] != BLANK {
				if state.Board[target.Row][target.Column] != turn {
					jump := NewCoordinate(target.Row+direction.Row, target.Column+direction.Column)

					if !state.CheckBound(jump) || state.Board[jump.Row][jump.Column] != BLANK {
						continue
					}

					jumpPresent = true
					moves[NewMove(from, jump, target)] = true
				}

				continue
			}

			if !jumpPresent {
				moves[NewMove(from, target, NO_JUMP)] = true
			}
		}
	}

	if jumpPresent {
		for move, _ := range moves {
			if move.Jump == NO_JUMP {
				delete(moves, move)
			}
		}
	}

	return moves
}

func (state *StateByte) GameEnd() (bool, byte) {
	whiteMoves := state.PossibleMovesAll(WHITE)
	blackMoves := state.PossibleMovesAll(BLACK)

	if len(whiteMoves) == 0 && len(blackMoves) == 0 {
		whitePieces := len(state.White)
		blackPieces := len(state.Black)

		if whitePieces > blackPieces {
			return true, WHITE
		} else if blackPieces > whitePieces {
			return true, BLACK
		} else {
			return true, BLANK
		}
	}

	return false, BLANK
}

func (state *StateByte) Skip() {
	if state.Turn == BLACK {
		state.Turn = WHITE
	} else {
		state.Turn = BLACK
	}
}
