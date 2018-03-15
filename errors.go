package checkers

const (
	ERROR_MOVE_TURN     string = "Wrong piece to move"
	ERROR_MOVE_INVALID  string = "Invalid move"
	ERROR_MOVE_BOUNDS   string = "Move out of bounds"
	ERROR_MOVE_OCCUPIED string = "Space is occupied"
	ERROR_MOVE_BLANK    string = "There is no piece in the specified coordinate"
	ERROR_MOVE_WRONG    string = "The piece at the specified coordinate is not yours"
	ERROR_MOVE_BACK string = "You cannot move backwards"
)

type MovementError struct {
	Message string
}

func NewMovementError(message string) MovementError {
	return MovementError{Message: message}
}

func (err MovementError) Error() string {
	return err.Message
}
