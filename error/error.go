package error

type NotFoundError struct {
	message string
}

func (e *NotFoundError) Error() string {
	return e.message
}

func NewNotFoundError(message string) error {
	return &NotFoundError{message: message}
}

type BadRequestError struct {
	message string
}

func (e *BadRequestError) Error() string {
	return e.message
}

func NewBadRequestError(message string) error {
	return &BadRequestError{message: message}
}

type UnprocessableEntityError struct {
	message string
}

func (e *UnprocessableEntityError) Error() string {
	return e.message
}

func NewUnprocessableEntityError(message string) error {
	return &UnprocessableEntityError{message: message}
}

type InternalServerError struct {
	message string
}

func (e *InternalServerError) Error() string {
	return e.message
}

func NewInternalServerError(message string) error {
	return &InternalServerError{message: message}
}

func Cause(err error) error {
	type Cause interface {
		Cause() error
	}

	for err != nil {
		cause, ok := err.(Cause)
		if !ok {
			break
		}
		err = cause.Cause()
	}

	return err
}
