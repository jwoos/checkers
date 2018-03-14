package checkers

const (
	ERROR_MOVE_TURN     string = "Wrong piece to move"
	ERROR_MOVE_INVALID  string = "Invalid move"
	ERROR_MOVE_BOUNDS   string = "Move out of bounds"
	ERROR_MOVE_OCCUPIED string = "Space is occupied"
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
