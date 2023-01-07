package controller

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
)

type Test struct {
	service Tester
}

func NewTest(srv Tester) *Test {
	return &Test{
		service: srv,
	}
}

type Tester interface {
	Test()
}

func (t Test) GetHello(c echo.Context) error {
	t.service.Test()

	return c.JSON(http.StatusOK, fmt.Sprintf("message: %s", "OK"))
}
