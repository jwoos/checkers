package checkers

type BoundsError struct {
	Message string
}

func NewBoundsError(message string) BoundsError {
	return BoundsError{Message: message}
}

func (err BoundsError) Error() string {
	return err.Message
}

type MovementError struct {
	Message string
}

func NewMovementError(message string) MovementError {
	return MovementError{Message: message}
}

func (err MovementError) Error() string {
	return err.Message
}
