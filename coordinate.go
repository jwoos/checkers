package checkers

import (
	"fmt"
)

type Coordinate struct {
	Row    int
	Column int
}

func NewCoordinate(row int, column int) *Coordinate {
	coord := Coordinate{
		Row:    row,
		Column: column,
	}

	return &coord
}

func (coord *Coordinate) String() string {
	return fmt.Sprintf("(%d, %d)", coord.Row, coord.Column)
}

func (coord *Coordinate) GoString() string {
	return coord.String()
}

func (coord *Coordinate) SetCoordinate(source *Coordinate) {
	coord.Row = source.Row
	coord.Column = source.Column
}

func (coord *Coordinate) ApplyCoordinate(source *Coordinate) {
	coord.Row += source.Row
	coord.Column += source.Column
}

func (coord *Coordinate) SetRow(row int) {
	coord.Row = row
}

func (coord *Coordinate) SetColumn(column int) {
	coord.Column = column
}

func (coord *Coordinate) Diagonal(direction Coordinate) *Coordinate {
	return NewCoordinate(coord.Row + direction.Row, coord.Column + direction.Column)
}
