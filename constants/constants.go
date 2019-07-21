package constants

import "errors"

var (
	//ErrorInvalidParams :
	ErrorInvalidParams = errors.New("Invalid Params")

	//ErrorEmailNotExist :
	ErrorEmailNotExist = errors.New("Email is not registered")

	//ErrorEmailExist :
	ErrorEmailExist = errors.New("Email already registered")

	//ErrorInvalidPassword :
	ErrorInvalidPassword = errors.New("Invalid password")
)
