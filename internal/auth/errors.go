package auth

import "errors"

var (
	EmailAlreadyExistErr    = errors.New("user with this email is already exist")
	UsernameAlreadyExistErr = errors.New("user with this username is already exist")
	IncorrectPasswordErr    = errors.New("password is incorrect")
)
