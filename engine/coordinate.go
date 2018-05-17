package checkers

import (
	"fmt"
)

type Coordinate struct {
	Row    int
	Column int
}

func NewCoordinate(row int, column int) Coordinate {
	return Coordinate{
		Row:    row,
		Column: column,
	}
}

func (coord Coordinate) String() string {
	return fmt.Sprintf("(%d, %d)", coord.Row, coord.Column)
}

func (coord Coordinate) GoString() string {
	return coord.String()
}

func (a Coordinate) ApplyCoordinate(b Coordinate) Coordinate {
	return NewCoordinate(a.Row+b.Row, a.Column+b.Column)
}
