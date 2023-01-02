package app

import (
	"chat-app-golang/internal/controller"
	"github.com/labstack/echo/v4"
	"log"
)

type App struct {
}

func New() *App {
	return &App{}
}

func (a App) Run() error {
	e := echo.New()

	e.GET("/test", controller.GetTest)

	if err := e.Start(":8080"); err != nil {
		log.Fatal(err)
	}

	return nil
}
