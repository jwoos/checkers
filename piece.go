package checkers


type Piece struct {
	King bool
	Coord Coordinate
}


func NewPiece(king bool, coord Coordinate) *Piece {
	piece := Piece{
		King: king,
		Coord: coord
	}

	return &piece
}

func (piece *Piece) SetKing(king bool) {
	piece.King = king
}

func (piece *Piece) SetCoord(coord Coordinate) {
	piece.Coord = coord
}

func (piece *Piece) SetX(x uint) {
	piece.Coord.X = x
}

func (piece *Piece) SetY (x uint) {
	piece.Coord.Y = y
}
