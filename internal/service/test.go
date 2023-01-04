package service

import (
	"chat-app-golang/internal/db/repository"
	"fmt"
)

type Test struct {
	userRepo repository.User
}

func NewTest(userRepo repository.User) *Test {
	return &Test{
		userRepo: userRepo,
	}
}

func (t Test) Hello() string {
	return fmt.Sprint("Hello, from server!")
}
