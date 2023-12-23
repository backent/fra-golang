package exceptions

type Forbidden struct {
	Message string
}

func (e *Forbidden) Error() string {
	return e.Message
}

func NewForbidden(message string) Forbidden {
	return Forbidden{Message: message}
}
