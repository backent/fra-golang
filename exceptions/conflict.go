package exceptions

type ConflictError struct {
	Message string
}

func NewConflictError(message string) ConflictError {
	return ConflictError{
		Message: message,
	}
}

func (implementation *ConflictError) Error() string {
	return implementation.Message
}
