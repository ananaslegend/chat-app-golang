package service

import "fmt"

type Test struct {
}

func NewTest() *Test {
	return &Test{}
}

func (t Test) Hello() string {
	return fmt.Sprint("Hello, from server!")
}
