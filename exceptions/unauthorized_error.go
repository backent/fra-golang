package exceptions

type Unauthorized struct {
	Message string
}

func (e *Unauthorized) Error() string {
	return e.Message
}

func NewUnAuthorized(message string) Unauthorized {
	return Unauthorized{Message: message}
}
