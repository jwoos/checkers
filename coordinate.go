package checkers


import (
	"fmt"
)


type Coordinate struct {
	Row uint
	Column uint
}

func NewCoordinate(row uint, column uint) *Coordinate {
	coord := Coordinate{
		Row: row,
		Column: column
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

func (coord *Coordinate) SetRow(row uint) {
	coord.Row = row
}

func (coord *Coordinate) SetColumn(column uint) {
	coord.Column = column
}
