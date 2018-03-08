package checkers


import (
	"fmt"
)


type Coordinate struct {
	X uint
	Y uint
}

func NewCoordinate(uint x, uint y) *Coordinate {
	coord := Coordinate{
		X: x,
		Y: y
	}

	return &coord
}

func (coord *Coordinate) String() string {
	return fmt.Sprintf("(%d, %d)", coord.X, coord.Y)
}

func (coord *Coordinate) GoString() string {
	return coord.String()
}

func (coord *Coordinate) SetX(uint x) {
	coord.X = x
}

func (coord *Coordinate) SetY(uint y) {
	coord.Y = y
}
