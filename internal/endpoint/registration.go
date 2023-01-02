package endpoint

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
)

type Registration struct {
	s Registrationer
}

type Registrationer interface {
	Registrate() error
}

func NewRegistration(service Registrationer) *Registration {
	return &Registration{
		s: service,
	}
}

func (r Registration) Registrate(c echo.Context) error {
	err := r.s.Registrate()
	if err != nil {
		c.JSON(http.StatusInternalServerError, fmt.Sprintf("error: %s", err.Error()))
	}

	return c.JSON(http.StatusInternalServerError, fmt.Sprintf("status: Succses"))
}
