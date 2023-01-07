package service

import "errors"

var (
	EmailAlreadyExistErr    = errors.New("email already exist")
	UsernameAlreadyExistErr = errors.New("username already exist")
)
