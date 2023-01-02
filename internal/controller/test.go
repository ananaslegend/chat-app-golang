package controller

import (
	"chat-app-golang/internal/usecase"
	"github.com/labstack/echo/v4"
	"net/http"
)

func GetTest(c echo.Context) error {
	var data = usecase.Hello()

	return c.JSON(http.StatusOK, &echo.Map{"data": data})
}
