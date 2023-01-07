package controller

import (
	"chat-app-golang/internal/dto"
	"chat-app-golang/internal/service/errors"
	"github.com/labstack/echo/v4"
	"net/http"
)

type Auth struct {
	s Auther
}

type Auther interface {
	Registration(registrationUser dto.RegistrationUser) error
}

func NewAuth(service Auther) *Auth {
	return &Auth{
		s: service,
	}
}

func (r Auth) Registration(c echo.Context) error {
	var regData dto.RegistrationUser
	if err := c.Bind(&regData); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest)
	}

	err := r.s.Registration(regData)
	if err != nil {
		switch err {
		case service.EmailAlreadyExistErr:
			return echo.NewHTTPError(http.StatusConflict, "User with this email is already exists")
		case service.UsernameAlreadyExistErr:
			return echo.NewHTTPError(http.StatusConflict, "User with this username is already exists")
		default:
			return echo.NewHTTPError(http.StatusInternalServerError)
		}
	}

	return c.NoContent(http.StatusCreated)
}
