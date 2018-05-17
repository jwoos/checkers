package checkers

import (
	"fmt"
)

var NO_JUMP = Coordinate{-1, -1}

type Move struct {
	From Coordinate
	To   Coordinate
	Jump Coordinate
}

func NewMove(from Coordinate, to Coordinate, jump Coordinate) Move {
	return Move{From: from, To: to, Jump: jump}
}

func (move Move) String() string {
	from := move.From
	to := move.To
	jump := move.Jump

	if move.Jump == NO_JUMP {
		return fmt.Sprintf("from[%v] jump[NONE] to[%v]", from, to)
	} else {
		return fmt.Sprintf("from[%v] jump[%v] to[%v]", from, jump, to)
	}
}

func (move Move) GoString() string {
	return move.String()
}
