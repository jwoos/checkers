package checkers

type Move struct {
	From *Coordinate
	To   *Coordinate
	Jump *Coordinate
}

func NewMove(from *Coordinate, to *Coordinate, jump *Coordinate) Move {
	return Move{From: from, To: to, Jump: jump}
}
