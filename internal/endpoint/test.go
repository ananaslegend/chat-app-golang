package endpoint

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
	Hello() string
}

func (t Test) GetHello(c echo.Context) error {
	result := t.service.Hello()

	return c.JSON(http.StatusOK, fmt.Sprintf("message: %s", result))
}
