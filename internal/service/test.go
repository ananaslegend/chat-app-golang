package service

import (
	"chat-app-golang/internal/db/repository"
	"chat-app-golang/internal/domain"
	"chat-app-golang/pkg/pointer"
	"context"
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

func (t Test) Test() {
	res, err := t.userRepo.Get(context.TODO(), domain.User{Username: pointer.ToString("Tester")})
	if err != nil {
		return
	}
	fmt.Println(res)
}
