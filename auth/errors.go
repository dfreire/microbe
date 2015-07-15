package Auth

import "errors"

const (
	EntityNotFound   error = errors.New("The entity was not found.")
	EmailAlreadyUsed error = errors.New("The email is already being used.")
)
