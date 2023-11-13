package exceptions

type NotFoundError struct {
	Message string
}

func NewNotFoundError(message string) NotFoundError {
	return NotFoundError{
		Message: message,
	}
}

func (implementation *NotFoundError) Error() string {
	return implementation.Message
}
