package exceptions

type BadRequestError struct {
	Message string
}

func NewBadRequestError(message string) BadRequestError {
	return BadRequestError{
		Message: message,
	}
}

func (implementation *BadRequestError) Error() string {
	return implementation.Message
}
